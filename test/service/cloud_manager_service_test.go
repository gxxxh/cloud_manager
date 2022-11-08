package service

import (
	"cloud_manager/src/codegen/openstack"
	"cloud_manager/src/service"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"testing"
)

func TestCallAliyunAPI(t *testing.T) {
	regionId := "cn-beijing"
	accessKeyId := "LTAI5tJKWj6qWB7t4VooErRx"
	accessKeySecret := "FsCABqUiecxe2NQmjlJl1321RcfxFV"
	params := make(map[string]string)
	params["regionId"] = regionId
	params["accessId"] = accessKeyId
	params["accessKeySecret"] = accessKeySecret
	params["cloudType"] = "aliyun"
	mcm, _ := service.NewMultiCloudManager(params)
	request := ecs.CreateDescribeRegionsRequest()
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(jsonBytes))
	ret, err := mcm.CallCloudAPI("DescribeRegions", jsonBytes)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ret)
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
	mcm, err := service.NewMultiCloudManager(authInfo)
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
	mcm, err := service.NewMultiCloudManager(authInfo)
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
