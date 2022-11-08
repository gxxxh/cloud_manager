package service

import (
	openstack2 "cloud_manager/src/codegen/openstack"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	"github.com/gophercloud/gophercloud/pagination"
	"log"
	"reflect"
	"testing"
)

func InitOpenstackClient(openstackClientType string) (*openstack2.OpenstackClient, error) {
	authInfo := map[string]string{
		"projectName":         "admin",
		"domainName":          "Default",
		"identityEndpoint":    "http://133.133.135.136:5000/v3",
		"username":            "admin",
		"password":            "ef1aa1ad78c442e1",
		"Region":              "RegionOne",
		"openstackClientType": openstackClientType,
	}
	oc, err := openstack2.NewOpenstackClient(authInfo)
	if err != nil {
		return nil, err
	}
	oc.InitClient(openstackClientType, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	return oc, nil
}
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
	tmp := resp.BaseResponse.GetHttpContentString()
	fmt.Println(tmp)
	//fmt.Printf("%v", resp)
}
func TestConstructErr(t *testing.T) {
	err := fmt.Errorf("this a err")
	errInterface := reflect.ValueOf(err).Interface()
	nerr := errInterface.(error)
	fmt.Println(nerr)
}

func TestOpenstackListFunc(t *testing.T) {
	scope := gophercloud.AuthScope{
		ProjectName: "admin",
		DomainName:  "Default",
	}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "http://133.133.135.136:5000/v3",
		Username:         "admin",
		Password:         "ef1aa1ad78c442e1",
		DomainName:       "Default",
		Scope:            &scope,
	}
	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		t.Error(err)
		return
	}
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		//Region: os.Getenv("OS_REGION_NAME"),
		Region: "RegionOne",
	})
	if err != nil {
		t.Error(err)
		return
	}
	pager := images.ListDetail(client, nil)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		imageList, err := images.ExtractImages(page)
		if err != nil {
			fmt.Errorf("Fatal error Extract Images:  %s \n", err)
		}
		for _, i := range imageList {
			// "i" will be an images.Image
			fmt.Printf("images is %v \n", i)
		}
		return false, err
	})
	if err != nil {
		t.Error(err)
	}
}

func TestOpenstackCodeGen(t *testing.T) {
	oc, err := InitOpenstackClient("compute")
	if err != nil {
		t.Error(err)
	}
	request := openstack2.NewListDetailComputeV2ImagesRequest()
	res := oc.ListDetailComputeV2Images(request)
	fmt.Println("direct: ", res)
	info, err := openstack2.ExtractListDetailComputeV2ImagesResponse(res)
	if err != nil {
		t.Error(err)
	}
	log.Println(info)
	////todo handle page type
	//err = res.Pager.EachPage(func(page pagination.Page) (bool, error) {
	//	imageList, err := images.ExtractImages(page)
	//	if err != nil {
	//		fmt.Errorf("Fatal error Extract Images:  %s \n", err)
	//	}
	//	for _, i := range imageList {
	//		// "i" will be a images.Image
	//		fmt.Printf("images is %v \n", i)
	//	}
	//	return false, err
	//})
}

func TestOpenstackReturnResultFunc(t *testing.T) {
	oc, err := InitOpenstackClient("identityv3")
	if err != nil {
		t.Error(err)
	}
	request := openstack2.NewGetIdentityV3TokensRequest()
	request.Token = "gAAAAABjafSUjIpPSF8EzOierSPpMsFZP8tcOOKvrzG_0f_VsThXWULt0pt9aFavxYcCdyLrCECg7EBauwMMQT84TkfHNF1W3WC-COnoyDSjPDoP3X2QzqqFfMRRDqRlIYrC8eHmnZoYOpnEnkeofSMk3oNqHNO-C5X2WopJ00PlZLvG-BSwX9w"
	res := oc.GetIdentityV3Tokens(request)
	content, err := json.Marshal(res.GetResult.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(content))
}
