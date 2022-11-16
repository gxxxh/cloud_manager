package service

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kube-stack/multicloud_service/src/codegen/openstack"
	"github.com/kube-stack/multicloud_service/src/codegen/registry"
	"github.com/kube-stack/multicloud_service/src/utils"
	"log"
)

type MultiCloudService struct {
	CloudType        string
	Client           interface{}
	requestRegistry  map[string]interface{}
	responseRegistry map[string]interface{}
}

func NewMultiCloudService(params map[string]string) (mcm *MultiCloudService, err error) {
	cloudType, ok := params["cloudType"]
	if !ok {
		err = fmt.Errorf("Error, the cloudType can't be empty")
		return
	}
	mcm = &MultiCloudService{
		CloudType: cloudType,
	}
	err = mcm.Init(params)
	return
}

func (m *MultiCloudService) Init(params map[string]string) (err error) {
	switch m.CloudType {
	case "aliyun":
		//regionId, accessId, accessKeySecret
		m.Client, err = ecs.NewClientWithAccessKey(params["regionId"], params["accessId"], params["accessKeySecret"])
		m.requestRegistry = registry.AliyunCreateRequestRegistry
		m.responseRegistry = nil
	case "openstack":
		//IdentityEndPoint, Username, Password
		m.Client, err = openstack.NewOpenstackClient(params)
		m.requestRegistry = registry.OpenstackCreateRequestRegistry
		m.responseRegistry = registry.OpenstackExtractResponseRegistry
	default:
		err = fmt.Errorf("unsupport cloud type")
	}
	if err != nil {
		log.Println("Init MultiCloudService error: ", err)
	}
	return
}

/*
using reflect to construct the parameters and call
*/
func (m *MultiCloudService) CallCloudAPI(cloudAPIName string, requestParameters []byte) ([]byte, error) {
	requestName := cloudAPIName
	request, err := utils.CallFunction(requestName, m.requestRegistry)
	if len(request) != 1 {
		err := fmt.Errorf("error, CreateRequestFunction return more than one value!, cloudAPIName is:%v", cloudAPIName)
		log.Println("CallCloudAPI error: ", err)
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	err = utils.ConstructStructOfPtr(request[0], requestParameters)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("%v", request)
	//createRequest only has one return value
	return m.doRequest(cloudAPIName, request[0])
}

func (m *MultiCloudService) doRequest(actionName string, request interface{}) ([]byte, error) {
	//find the client's method
	ret, err := utils.CallMethod(m.Client, actionName, request)
	if err != nil {
		return nil, err
	}
	switch m.CloudType {
	case "aliyun":

	case "openstack":
		if len(ret) != 1 {
			err = fmt.Errorf("the action %s in aliyun should only return two result\n", actionName)
			log.Println("doRequest Error: ", err)
			return nil, err
		}
		// extract response
		ret, err = utils.CallFunction(actionName, m.responseRegistry, ret[0])
	}
	if len(ret) != 2 {
		err = fmt.Errorf("the action %s in aliyun should only return two result\n", actionName)
		log.Println("doRequest Error: ", err)
		return nil, err
	}
	//ret[1] should be a error
	if ret[1] != nil {
		err = ret[1].(error)
		log.Println("sdk do request error: ", err)
	}
	jsonBytes, err := json.Marshal(ret[0])
	if err != nil {
		log.Println("marshal response error: ", err)
		return nil, err
	}
	return jsonBytes, err
	//retValue := reflect.ValueOf(ret[0]).Elem()
	//fmt.Println(retValue.NumField())
	//tmp1 := retValue.Field(0).Interface()
	////fmt.Println(tmp1)
	//tmp2 := reflect.ValueOf(tmp1).Elem()
	//fmt.Println(tmp2.CloudType())
	//fmt.Println(tmp2.NumField())
	//fmt.Println(tmp2.Interface())

}
