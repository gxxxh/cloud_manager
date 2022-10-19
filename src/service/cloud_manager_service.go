package service

import (
	"cloud_manager/src/codegen/aliyun"
	"cloud_manager/src/utils"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"reflect"
)

type MultiCloudManager struct {
	Kind            string
	Client          interface{}
	requestRegistry map[string]interface{}
}

func NewMultiCloudManager(kind string) *MultiCloudManager {
	mcm := &MultiCloudManager{
		Kind: kind,
	}
	mcm.Init()
	return mcm
}
func (m *MultiCloudManager) Init() error {
	switch m.Kind {
	case "aliyun":
		client, err := ecs.NewClientWithAccessKey("cn-beijing", "11111111111111", "222222222222222222222")
		if err != nil {
			fmt.Println(err)
			return err
		}
		m.Client = client
		m.requestRegistry = aliyun.CreateRequestRegistry
	default:
		err := fmt.Errorf("unsupport cloud type")
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
		return "", err
	}
	if err != nil {
		return "", err
	}
	err = utils.ConstructStructByPtr(request, requestParameters)
	if err != nil {
		return "", err
	}
	fmt.Printf("%v", request)
	//createRequest only has one return value
	return m.doRequest(actionName, request[0])
}

func (m *MultiCloudManager) doRequest(actionName string, request interface{}) (string, error) {
	//find the client's method
	ret, err := utils.CallMethod(m.Client, actionName, request)
	if err != nil {
		return "", err
	}
	if len(ret) != 1 {
		return "", fmt.Errorf("the action %s should only return one result\n", actionName)
	}
	return fmt.Sprintf("%v", reflect.ValueOf(ret[0])), err
}
