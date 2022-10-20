package service

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"reflect"
	"testing"
)

func TestAliyunSDK(t *testing.T) {
	regionId := "cn-beijing"
	accessKeyId := "LTAI5tJKWj6qWB7t4VooErRx"
	accessKeySecret := "FsCABqUiecxe2NQmjlJl1321RcfxFV"
	client, err := ecs.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		t.Error(err)
	}
	describeRegionRequest := ecs.CreateDescribeRegionsRequest()
	resp, err := client.DescribeRegions(describeRegionRequest)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%v", resp.Regions)
	descrebeInstanceReqeust := ecs.CreateDescribeInstancesRequest()
	res, err := client.DescribeInstances(descrebeInstanceReqeust)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res.Instances)
	//fmt.Printf("%v", res.Instances)
}
func TestConstructErr(t *testing.T) {
	err := fmt.Errorf("this a err")
	errInterface := reflect.ValueOf(err).Interface()
	nerr := errInterface.(error)
	fmt.Println(nerr)
}
