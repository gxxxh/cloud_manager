package service

import (
	"cloud_manager/src/service"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	regionId := "cn-beijing"
	accessKeyId := "LTAI5tJKWj6qWB7t4VooErRx"
	accessKeySecret := "FsCABqUiecxe2NQmjlJl1321RcfxFV"
	mcm, _ := service.NewMultiCloudManager("aliyun", regionId, accessKeyId, accessKeySecret)
	request := ecs.CreateDescribeRegionsRequest()
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
