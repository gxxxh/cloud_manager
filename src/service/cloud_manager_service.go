package service

import (
	"cloud_manager/src/codegen/aliyun"
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

func NewMultiCloudManager(kind string, params ...string) (*MultiCloudManager, error) {
	mcm := &MultiCloudManager{
		Kind: kind,
	}
	err := mcm.Init(params...)
	if err != nil {
		return nil, err
	}
	return mcm, err
}
func (m *MultiCloudManager) Init(params ...string) error {
	switch m.Kind {
	case "aliyun":
		client, err := ecs.NewClientWithAccessKey(params[0], params[1], params[2])
		if err != nil {
			log.Println("Init MultiCloudManager error: ", err)
			return err
		}
		m.Client = client
		m.requestRegistry = aliyun.CreateRequestRegistry
	default:
		err := fmt.Errorf("unsupport cloud type")
		log.Println("Init MultiCloudManager error: ", err)
		return err
	}
	return nil
}

/*
using reflect to construct the parameters and call
*/
func (m *MultiCloudManager) HandleRequest(actionName string, requestParameters []byte) (string, error) {
	requestName := actionName + "Request"
	request, err := utils.CallFunction(requestName, m.requestRegistry)
	if len(request) != 1 {
		err := fmt.Errorf("error, CreateRequestFunction return more than one value!, actionName is:%v", actionName)
		log.Println("HandleRequest error: ", err)
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
	return m.doRequest(actionName, request[0])
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
	str := fmt.Sprintf("%v", ret[0])
	return str, err
}
