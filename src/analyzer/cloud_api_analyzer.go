package analyzer

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
	"strings"
)

/**
基于反射的方式，动态的分析client包含的方法，并能解析方法对应的参数。
*/

// todo, 是针对每个ecs生成一个代码，编译后只能针对该云操作，
// using to analyze cloudapi by reflect(dynamically)
type CloudAPIAnalyzer struct {
	MethodMap  map[string]reflect.Method //控制云资源的方法
	RequestMap map[string]reflect.Type   //方法相关的request类型
	//RequestInfos []RequestInfo
}

func NewCloudAPIAnalyzer() *CloudAPIAnalyzer {
	ca := &CloudAPIAnalyzer{}
	ca.Init()
	return ca
}

func (c *CloudAPIAnalyzer) Init() {
	c.MethodMap = make(map[string]reflect.Method)
	c.RequestMap = make(map[string]reflect.Type)
}

func (c *CloudAPIAnalyzer) ExtractCloudAPIs(client interface{}) {
	//log.Println("CloudAPIAnalyzer: analyzer client's method")
	// Notice the *Type can access all the Type Methods. But Type cannot access to *Type Methods.
	clientType := reflect.TypeOf(client)
	//use reflect.PtrTo(Type) to convert from Type to *Type
	if clientType.Kind() != reflect.Ptr && clientType.Kind() != reflect.Pointer {
		clientType = reflect.PtrTo(clientType)
	}

	for i := 0; i < clientType.NumMethod(); i++ {
		method := clientType.Method(i)
		//filter by method format
		if method.Type.NumIn() == 2 && !strings.HasSuffix(method.Name, "WithChan") {
			c.ExtractMethodParameters(method)
		}
	}
	//log.Printf("extract %v methods\n", len(c.MethodMap))
	//log.Println("CloudAPIAnalyzer: analyzer client's method done")
}

func (c *CloudAPIAnalyzer) ExtractMethodParameters(method reflect.Method) {
	methodType := method.Type
	paraType := methodType.In(1)         // the first element is client itself
	for paraType.Kind() == reflect.Ptr { //pointer type's name is non
		paraType = paraType.Elem()
	}
	// filter by parameter name
	if _, ok := c.RequestMap[paraType.Name()]; !ok && strings.HasSuffix(paraType.Name(), "Request") && paraType.Name() != "CommonRequest" {
		//log.Printf("extract parameter type for method:%v, num parameters:%v, package path: %v\n", method.Name, method.Type.NumIn(), method.Type.String())
		c.MethodMap[method.Name] = method
		c.RequestMap[paraType.Name()] = paraType
	}
}

/*
this function is used to extract a type eecursively to basic golang type
*/
func (c *CloudAPIAnalyzer) ExtractType(dataType reflect.Type) interface{} {
	for dataType.Kind() == reflect.Ptr {
		dataType = dataType.Elem()
	}
	switch dataType.Kind() {
	case reflect.Struct:
		typeInfo := make(map[string]interface{})
		for i := 0; i < dataType.NumField(); i++ {
			structFiled := dataType.Field(i)
			if structFiled.Name != "RpcRequest" {
				typeInfo[structFiled.Name] = c.ExtractType(structFiled.Type)
			}
		}
		return typeInfo
	case reflect.Array, reflect.Slice:
		typeInfo := make(map[string]interface{})
		typeInfo[dataType.Kind().String()] = c.ExtractType(dataType.Elem())
		return typeInfo
	case reflect.Map:
		typeInfo := make(map[string]interface{})
		typeInfo["map_key_type"] = c.ExtractType(dataType.Key())
		typeInfo["map_value_type"] = c.ExtractType(dataType.Elem())
		return typeInfo
	default:
		return dataType.Kind().String()
	}
}

// 需要在调用了extractCloudAPIs之后再使用，
func (c *CloudAPIAnalyzer) ExtractRequestInfos(createFuncPre string) *RequestRegistryInfo {
	requestRegistryInfo := NewRequestRegistryInfo()
	for requestTypeName, requestType := range c.RequestMap {
		requestRegistryInfo.RequestInfos = append(requestRegistryInfo.RequestInfos, NewRequestInfo(requestTypeName, requestType, createFuncPre))
	}
	return requestRegistryInfo
}

func (c *CloudAPIAnalyzer) SaveToJson(path string) {
	paraInfoMap := make(map[string]interface{})
	for paraName, paraType := range c.RequestMap {
		typeInfo := c.ExtractType(paraType)
		paraInfoMap[paraName] = typeInfo
	}
	paraInfoJson, err := json.Marshal(paraInfoMap)
	if err != nil {
		log.Panicf("SaveToJson Error: %v \n", err)
	}
	filePtr, err := os.Create(path)
	if err != nil {
		log.Panicf("SaveToJson Error: %v \n", err)
	}
	defer filePtr.Close()
	filePtr.Write(paraInfoJson)
}
