package analyzer

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"os"
	"reflect"
	"strings"
)

type CloudAPIAnalyzer struct {
	Kind         string
	client       interface{}
	MethodMap    map[string]reflect.Method
	RequestMap   map[string]reflect.Type
	RequestInfos []RequestInfo
}

func (c *CloudAPIAnalyzer) Init() {
	switch c.Kind {
	case "aliyun":
		c.client, _ = ecs.NewClientWithAccessKey("cn-beijing", "11111111111111", "222222222222222222222")
	default:
		fmt.Println("error!, you need to provide a cloud type")
		panic("cloud type is nil")
	}
	c.MethodMap = make(map[string]reflect.Method)
	c.RequestMap = make(map[string]reflect.Type)
}

func (c *CloudAPIAnalyzer) ExtractCloudAPIs() {
	// Notice the *Type can access all the Type Methods. But Type cannot access to *Type Methods.
	clientType := reflect.TypeOf(c.client)
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
	//fmt.Println("num of methods: ", len(MethodMap))

}

func (c *CloudAPIAnalyzer) ExtractMethodParameters(method reflect.Method) {
	methodType := method.Type
	paraType := methodType.In(1)         // the first element is client itself
	for paraType.Kind() == reflect.Ptr { //pointer type's name is non
		paraType = paraType.Elem()
	}
	// filter by parameter name
	if _, ok := c.RequestMap[paraType.Name()]; !ok && strings.HasSuffix(paraType.Name(), "Request") {
		fmt.Printf("extract parameter type for method:%v, num parameters:%v, package path: %v\n", method.Name, method.Type.NumIn(), method.Type.String())
		c.MethodMap[method.Name] = method
		c.RequestMap[paraType.Name()] = paraType
	}
}

func (c *CloudAPIAnalyzer) ExtractType(dataType reflect.Type) interface{} {
	for dataType.Kind() == reflect.Ptr {
		dataType = dataType.Elem()
	}
	switch dataType.Kind() {
	case reflect.Struct:
		typeInfo := make(map[string]interface{})
		for i := 0; i < dataType.NumField(); i++ {
			structFiled := dataType.Field(i)
			typeInfo[structFiled.Name] = c.ExtractType(structFiled.Type)
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
func (c *CloudAPIAnalyzer) ExtractRequestInfos() []RequestInfo {
	for requestTypeName, requestType := range c.RequestMap {
		c.RequestInfos = append(c.RequestInfos, *NewRequestInfo(requestTypeName, requestType))
	}
	return c.RequestInfos
}

func (c *CloudAPIAnalyzer) SaveToJson() {
	paraInfoMap := make(map[string]interface{})
	for paraName, paraType := range c.RequestMap {
		typeInfo := c.ExtractType(paraType)
		paraInfoMap[paraName] = typeInfo
	}
	paraInfoJson, err := json.Marshal(paraInfoMap)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	filePtr, err := os.Create("create_image.json")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer filePtr.Close()
	filePtr.Write(paraInfoJson)
}
