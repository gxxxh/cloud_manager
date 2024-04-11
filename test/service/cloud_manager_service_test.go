package service

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kube-stack/multicloud_service/src/codegen/openstack"
	"github.com/kube-stack/multicloud_service/src/service"
	"github.com/tidwall/gjson"
	"testing"
)

func TestCallAliyunAPI(t *testing.T) {
	regionId := "cn-beijing"
	accessKeyId := ""
	accessKeySecret := ""
	params := make(map[string]string)
	params["regionId"] = regionId
	params["accessId"] = accessKeyId
	params["accessKeySecret"] = accessKeySecret
	params["cloudType"] = "aliyun"
	mcm, _ := service.NewMultiCloudService(params)
	request := ecs.CreateDescribeInstancesRequest()
	request.InstanceIds = "[\"i-2zegiq87g0txkt1bvrb5\"]"
	//request.RegionId = "cn-beijing"
	//request := ecs.CreateDescribeRegionsRequest()
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(jsonBytes))
	//ret, err := mcm.CallCloudAPI("DescribeRegions", jsonBytes)
	ret, err := mcm.CallCloudAPI("DescribeInstances", jsonBytes)
	if err != nil {
		t.Error(err)
	}
	instanceId := gjson.GetBytes(ret, "Instances.Instance.0.InstanceId")
	fmt.Println(instanceId)
	fmt.Println(string(ret))
}

func TestCallOpenstackReturnPager(t *testing.T) {
	authInfo := map[string]string{
		"projectName":         "admin",
		"domainName":          "Default",
		"identityEndpoint":    "http://133.133.135.136:5000/v3",
		"username":            "admin",
		"password":            "ef1aa1ad78c442e1",
		"Region":              "RegionOne",
		"openstackClientType": "compute",
		"cloudType":           "openstack",
	}
	mcm, err := service.NewMultiCloudService(authInfo)
	request := openstack.NewListDetailComputeV2ImagesRequest()
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(jsonBytes))
	ret, err := mcm.CallCloudAPI("ListDetailComputeV2Images", jsonBytes)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ret)
}

func TestCallOpenstackReturnResult(t *testing.T) {
	authInfo := map[string]string{
		"projectName":         "admin",
		"domainName":          "Default",
		"identityEndpoint":    "http://133.133.135.136:5000/v3",
		"username":            "admin",
		"password":            "ef1aa1ad78c442e1",
		"Region":              "RegionOne",
		"openstackClientType": "identityv3",
		"cloudType":           "openstack",
	}
	mcm, err := service.NewMultiCloudService(authInfo)
	request := openstack.NewGetIdentityV3TokensRequest()
	request.Token = "gAAAAABjafSUjIpPSF8EzOierSPpMsFZP8tcOOKvrzG_0f_VsThXWULt0pt9aFavxYcCdyLrCECg7EBauwMMQT84TkfHNF1W3WC-COnoyDSjPDoP3X2QzqqFfMRRDqRlIYrC8eHmnZoYOpnEnkeofSMk3oNqHNO-C5X2WopJ00PlZLvG-BSwX9w"
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(jsonBytes))
	ret, err := mcm.CallCloudAPI("GetIdentityV3Tokens", jsonBytes)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ret)
}
func TestCallOpenstacCreateServer(t *testing.T) {
	authInfo := map[string]string{
		"projectName":         "admin",
		"domainName":          "Default",
		"identityEndpoint":    "http://133.133.135.136:5000/v3",
		"username":            "admin",
		"password":            "ef1aa1ad78c442e1",
		"Region":              "RegionOne",
		"openstackClientType": "compute",
		"cloudType":           "openstack",
	}
	mcm, err := service.NewMultiCloudService(authInfo)
	request := openstack.NewCreateComputeV2ServersRequest()
	request.Opts.Name = "test"
	request.Opts.ImageRef = "952b386b-6f30-46f6-b019-f522b157aa3a"
	request.Opts.FlavorRef = "2"
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(jsonBytes))
	ret, err := mcm.CallCloudAPI("CreateComputeV2Servers", jsonBytes)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ret)
}
func TestCallOpenstacGetServer(t *testing.T) {
	authInfo := map[string]string{
		"projectName":         "admin",
		"domainName":          "Default",
		"identityEndpoint":    "http://133.133.135.136:5000/v3",
		"username":            "admin",
		"password":            "ef1aa1ad78c442e1",
		"Region":              "RegionOne",
		"openstackClientType": "compute",
		"cloudType":           "openstack",
	}
	mcm, err := service.NewMultiCloudService(authInfo)
	request := openstack.NewGetComputeV2ServersRequest()
	request.Id = "ddf4da9b-84fc-436d-86b4-5378c8b7a80f"
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(jsonBytes))
	ret, err := mcm.CallCloudAPI("GetComputeV2Servers", jsonBytes)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(ret))
}
