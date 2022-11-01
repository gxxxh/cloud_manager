package service

import (
	"cloud_manager/src/codegen/openstack"
	"cloud_manager/src/codegen/registry"
	"cloud_manager/src/utils"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"log"
)

type MultiCloudManager struct {
	Kind            string
	Client          interface{}
	requestRegistry map[string]interface{}
}

func NewMultiCloudManager(params map[string]string) (mcm *MultiCloudManager, err error) {
	kind, ok := params["kind"]
	if !ok {
		err = fmt.Errorf("Error, the kind can't be empty")
		return
	}
	mcm = &MultiCloudManager{
		Kind: kind,
	}
	err = mcm.Init(params)
	return
}

func (m *MultiCloudManager) Init(params map[string]string) (err error) {
	switch m.Kind {
	case "aliyun":
		//regionId, accessId, accessKeySecret
		m.Client, err = ecs.NewClientWithAccessKey(params["regionId"], params["accessId"], params["accessKeySecret"])
		m.requestRegistry = registry.AliyunCreateRequestRegistry
	case "openstack":
		//IdentityEndPoint, Username, Password
		m.Client, err = openstack.NewOpenstackClient(params)
		m.requestRegistry = registry.OpenstackCreateRequestRegistry
	default:
		err = fmt.Errorf("unsupport cloud type")
	}
	if err != nil {
		log.Println("Init MultiCloudManager error: ", err)
	}
	return
}

/*
using reflect to construct the parameters and call
*/
func (m *MultiCloudManager) CallCloudAPI(cloudAPIName string, requestParameters []byte) (string, error) {
	requestName := cloudAPIName + "Request"
	request, err := utils.CallFunction(requestName, m.requestRegistry)
	if len(request) != 1 {
		err := fmt.Errorf("error, CreateRequestFunction return more than one value!, cloudAPIName is:%v", cloudAPIName)
		log.Println("CallCloudAPI error: ", err)
		return "", err
	}
	if err != nil {
		return "", err
	}
	err = utils.ConstructStructOfPtr(request[0], requestParameters)
	if err != nil {
		return "", err
	}
	//fmt.Printf("%v", request)
	//createRequest only has one return value
	return m.doRequest(cloudAPIName, request[0])
}

func (m *MultiCloudManager) doRequest(actionName string, request interface{}) (string, error) {
	//find the client's method
	ret, err := utils.CallMethod(m.Client, actionName, request)
	if err != nil {
		return "", err
	}
	if len(ret) != 2 {
		err = fmt.Errorf("the action %s should only return two result\n", actionName)
		log.Println("doRequest Error: ", err)
		return "", err
	}
	//ret[1] should be a error
	if ret[1] != nil {
		err = ret[1].(error)
		log.Println("sdk do request error: ", err)
	}
	//retValue := reflect.ValueOf(ret[0]).Elem()
	//fmt.Println(retValue.NumField())
	//tmp1 := retValue.Field(0).Interface()
	////fmt.Println(tmp1)
	//tmp2 := reflect.ValueOf(tmp1).Elem()
	//fmt.Println(tmp2.Kind())
	//fmt.Println(tmp2.NumField())
	//fmt.Println(tmp2.Interface())
	str := fmt.Sprintf("%v", ret[0])
	return str, err
}
