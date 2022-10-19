package service

import (
	"cloud_manager/src/service"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	mcm := service.NewMultiCloudManager("aliyun")
	request := ecs.CreateDescribeRegionsRequest()
	request.ResourceOwnerId = requests.Integer("1")
	request.InstanceChargeType = "instance charge type"
	request.ResourceOwnerAccount = "resource owner account"
	request.OwnerAccount = "owner account"
	request.OwnerId = requests.Integer("1")
	request.ResourceType = "resource type"
	request.AcceptLanguage = "accept language"
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(jsonBytes))
	ret, err := mcm.HandleRequest("DescribeRegions", jsonBytes)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(ret)
}
