package gen

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	multicloud_service "github.com/kube-stack/multicloud_service/src/analyzer"
	"github.com/kube-stack/multicloud_service/src/codegen/openstack"
	"github.com/kube-stack/multicloud_service/src/utils"
	"log"
	"path/filepath"
	"reflect"
	"strings"
)

var TypeMap = map[string]string{
	"string": "String",
	"int":    "Integer",
	"bool":   "Boolean",
}

type JavaSDKGenerator struct {
	Config           *CloudConfig
	CloudAPIAnalyzer *multicloud_service.CloudAPIAnalyzer
	JavaResources    map[string]*JavaResource
}

func NewJavaSDKGenerator(config *CloudConfig) *JavaSDKGenerator {
	cloudAPIAnalyzer := multicloud_service.NewCloudAPIAnalyzer()
	switch config.CloudType {
	case "Aliyun":
		client := ecs.Client{}
		cloudAPIAnalyzer.ExtractCloudAPIs(client)
	case "Openstack":
		client := openstack.OpenstackClient{}
		cloudAPIAnalyzer.ExtractCloudAPIs(client)
	}
	javaSDKGenerator := &JavaSDKGenerator{
		Config:           config,
		CloudAPIAnalyzer: cloudAPIAnalyzer,
		JavaResources:    nil,
	}
	javaSDKGenerator.initJavaResources()
	return javaSDKGenerator
}

func (j *JavaSDKGenerator) getDomain(target string) *JavaClass {
	for methodName, method := range j.CloudAPIAnalyzer.MethodMap {
		actionName, cloudResourceName, _ := utils.ParseRequestName(methodName)
		methodType := method.Type
		//思路；get方法（GetComputeV2Servers） → 返回类型（GetComputeV2ServersResponse）
		// → result类型（GetResult） → extract方法（serverResult.Extract） → server类型
		if actionName == "Get" && cloudResourceName == target {
			log.Println(methodType.Name())
			returnType := methodType.Out(0)
			for returnType.Kind() == reflect.Ptr {
				returnType = returnType.Elem()
			}
			for i := 0; i < returnType.NumField(); i++ {
				structField := returnType.Field(i)
				if strings.Contains(structField.Name, "Result") {
					getResultType := structField.Type
					for k := 0; k < getResultType.NumMethod(); k++ {
						if getResultType.Method(k).Name == "Extract" && getResultType.Method(k).Type.NumOut() > 0 {
							domainType := getResultType.Method(k).Type.Out(0)
							return j.genJavaClass(domainType, cloudResourceName, 1)
						}
					}
				}
			}
		}
	}
	return nil
}

func (j *JavaSDKGenerator) initJavaResources() {
	j.JavaResources = make(map[string]*JavaResource)
	for requestName, requestType := range j.CloudAPIAnalyzer.RequestMap {
		actionName, cloudResourceName, resourceName := utils.ParseRequestName(requestName)
		//todo 删除不支持的操作
		if actionName == "List" {
			continue
		}
		if _, ok := j.JavaResources[cloudResourceName]; !ok {

			javaClassName := utils.GetJavaResourceName(j.Config.CloudType, resourceName)
			javaResource := NewJavaResource(javaClassName, cloudResourceName)
			javaResource.JavaClass = *NewJavaClass("Lifecycle", "Lifecycle", 1)
			j.JavaResources[cloudResourceName] = javaResource
		}

		//创建request type
		for requestType.Kind() == reflect.Ptr {
			requestType = requestType.Elem()
		}
		if requestType.Kind() != reflect.Struct {
			log.Panicf("Request %s Type should be a struct\n", requestType.Name())
		}
		log.Println("Analyze request ", requestName)
		requestClass := j.genJavaClass(requestType, actionName+cloudResourceName, 2)
		//ClassName: CreateComputeV2Request → Create
		requestClass.ClassName = actionName
		requestVariable := NewMemberVariable(actionName, actionName, actionName+cloudResourceName, Basic)
		//add to lifecycle class
		j.JavaResources[cloudResourceName].JavaClass.Add(requestVariable, requestClass)
	}
}

// class code → java_resource_class.tmpl
func (j *JavaSDKGenerator) GenResourceClass(target string) {
	data := make(map[string]interface{})
	javaResource, ok := j.JavaResources[target]
	if !ok {
		log.Panicf("no resource for %v in %v\n", target, j.Config.CloudType)
	}
	data["JavaResourceName"] = javaResource.Name
	templatePath := filepath.Join(j.Config.JavaCodeConfig.TemplatePath, JavaResourceClassTemplate)
	codePath := filepath.Join(j.Config.JavaCodeConfig.CodePath, "api", "models", strings.ToLower(j.Config.CloudType), javaResource.Name+".java")
	GenAndSaveCode(templatePath, codePath, data, nil)

}

// spec code → java_resource_spec.tmpl
func (j *JavaSDKGenerator) GenResourceSpec(target string) {
	data := make(map[string]interface{})
	javaResource, ok := j.JavaResources[target]
	if !ok {
		log.Panicf("no resource for %v in %v\n", target, j.Config.CloudType)
	}
	data["JavaResourceName"] = javaResource.Name
	templatePath := filepath.Join(j.Config.JavaCodeConfig.TemplatePath, JavaResourceSpecTemplate)
	codePath := filepath.Join(j.Config.JavaCodeConfig.CodePath, "api", "specs", strings.ToLower(j.Config.CloudType), strings.ToLower(javaResource.Name)+"Spec.java")
	GenAndSaveCode(templatePath, codePath, data, nil)

}

func (j *JavaSDKGenerator) GenResourceDomain(target string) {

	javaResource, ok := j.JavaResources[target]
	if !ok {
		log.Panicf("no resource for %v in %v\n", target, j.Config.CloudType)
	}
	//generate header
	headerData := utils.Struct2Map(javaResource)
	params := make(map[string]interface{})
	params["CloudName"] = j.Config.CloudType
	headerTemplatePath := filepath.Join(j.Config.JavaCodeConfig.TemplatePath, JavaResourceDomainHeaderTemplate)
	headerCode, err := GenCode(headerTemplatePath, headerData, params)
	if err != nil {
		log.Panicf("Gen Lifecycle Header err, %v\n", err)
	}
	//generate lifecycle class
	domainClass := j.getDomain(target)
	domainClass.ClassName = "Domain"

	classData := utils.Struct2Map(domainClass)
	classTemplatePath := filepath.Join(j.Config.JavaCodeConfig.TemplatePath, JavaClassTemplate)
	classCode, err := GenCode(classTemplatePath, classData, nil)
	if err != nil {
		log.Panicf("Gen Lifecycle Class err, %v\n", err)
	}
	//save lifecycle code
	codePath := filepath.Join(j.Config.JavaCodeConfig.CodePath, "api", "specs", strings.ToLower(j.Config.CloudType), strings.ToLower(javaResource.Name), "Domain.java")
	content := append(headerCode, classCode...)
	utils.Save(content, codePath)
}

func (j *JavaSDKGenerator) GenResourceLifecycle(target string) {
	javaResource, ok := j.JavaResources[target]
	if !ok {
		log.Panicf("no resource for %v in %v\n", target, j.Config.CloudType)
	}
	//generate header
	headerData := utils.Struct2Map(javaResource)
	params := make(map[string]interface{})
	params["CloudName"] = j.Config.CloudType
	headerTemplatePath := filepath.Join(j.Config.JavaCodeConfig.TemplatePath, JavaResourceDomainHeaderTemplate)
	headerCode, err := GenCode(headerTemplatePath, headerData, params)
	if err != nil {
		log.Panicf("Gen Lifecycle Header err, %v\n", err)
	}
	//generate lifecycle class
	classData := utils.Struct2Map(javaResource.JavaClass)
	classTemplatePath := filepath.Join(j.Config.JavaCodeConfig.TemplatePath, JavaClassTemplate)
	classCode, err := GenCode(classTemplatePath, classData, nil)
	if err != nil {
		log.Panicf("Gen Lifecycle Class err, %v\n", err)
	}
	//save lifecycle code
	codePath := filepath.Join(j.Config.JavaCodeConfig.CodePath, "api", "specs", strings.ToLower(j.Config.CloudType), strings.ToLower(javaResource.Name), "Lifecycle.java")
	content := append(headerCode, classCode...)
	utils.Save(content, codePath)
}

func (j *JavaSDKGenerator) GenResourceImpl() {

}

func (j *JavaSDKGenerator) genMemberVariable(name string, t reflect.Type, jsonInfo string) *MemberVariable {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		elemType := t.Elem()
		for elemType.Kind() == reflect.Ptr {
			elemType = elemType.Elem()
		}
		return NewMemberVariable(name, elemType.Name(), jsonInfo, Array)
	case reflect.Map:
		valueType := t.Elem()
		for valueType.Kind() == reflect.Ptr {
			valueType = valueType.Elem()
		}
		return NewMemberVariable(name, valueType.Name(), jsonInfo, Map)
	default:
		return NewMemberVariable(name, t.Name(), jsonInfo, Basic)
	}
	return nil
}

func (j *JavaSDKGenerator) genJavaClass(t reflect.Type, jsonInfo string, depth int) *JavaClass {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Name() == "Time" {
		return nil
	}
	switch t.Kind() {
	case reflect.Struct:
		return j.handleStruct(t, jsonInfo, depth)
	case reflect.Array, reflect.Slice:
		return j.genArrayClass(t, "", depth)
	case reflect.Map:
		return j.genMapClass(t, "", depth)
	default:
		return nil
	}
	return nil
}

// 解析struct生成javaClass
func (j *JavaSDKGenerator) handleStruct(t reflect.Type, jsonInfo string, depth int) *JavaClass {
	javaClass := NewJavaClass(t.Name(), jsonInfo, depth)
	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i)
		if structField.Name != "RpcRequest" { //aliyun,不考虑rpcRequest
			//若不是基础类型，生成memberClass
			fieldClass := j.genJavaClass(structField.Type, structField.Tag.Get("json"), depth+1)
			memberVariable := j.genMemberVariable(structField.Name, structField.Type, structField.Tag.Get("json"))
			javaClass.Add(memberVariable, fieldClass)
		}
	}
	return javaClass
}

// 解析array/slice生成javaclass
func (j *JavaSDKGenerator) genArrayClass(t reflect.Type, jsonInfo string, depth int) *JavaClass {
	elemType := t.Elem()
	return j.genJavaClass(elemType, jsonInfo, depth)
}

// 解析map的value生成javaclass
func (j *JavaSDKGenerator) genMapClass(t reflect.Type, jsonInfo string, depth int) *JavaClass {
	//todo key一般为string，略过
	valueType := t.Elem()
	return j.genJavaClass(valueType, jsonInfo, depth)
}
