package service

import (
	openstack2 "cloud_manager/src/codegen/openstack"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	"github.com/gophercloud/gophercloud/pagination"
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
	//fmt.Printf("%v", resp)
	fmt.Println(resp)
}
func TestConstructErr(t *testing.T) {
	err := fmt.Errorf("this a err")
	errInterface := reflect.ValueOf(err).Interface()
	nerr := errInterface.(error)
	fmt.Println(nerr)
}

func TestOpenstackSDK(t *testing.T) {
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
			// "i" will be a images.Image
			fmt.Printf("images is %v \n", i)
		}
		return false, err
	})
	if err != nil {
		t.Error(err)
	}
}

func TestOpenstackCodeGen(t *testing.T) {
	authInfo := map[string]string{
		"projectName":      "admin",
		"domainName":       "Default",
		"identityEndpoint": "http://133.133.135.136:5000/v3",
		"username":         "admin",
		"password":         "ef1aa1ad78c442e1",
	}
	params := map[string]interface{}{
		"kind":     "Openstack",
		"authInfo": authInfo,
	}
	params["kind"] = "openstack"
	oc, err := openstack2.NewOpenstackClient(authInfo)
	if err != nil {
		t.Error(err)
	}
	oc.InitClient("compute", gophercloud.EndpointOpts{
		//Region: os.Getenv("OS_REGION_NAME"),
		Region: "RegionOne",
	})
	request := openstack2.NewListDetailComputeV2ImagesRequest()
	res := oc.ListDetailComputeV2Images(request)
	fmt.Println("direct: ", res)
	//todo handle page type
	err = res.Pager.EachPage(func(page pagination.Page) (bool, error) {
		imageList, err := images.ExtractImages(page)
		if err != nil {
			fmt.Errorf("Fatal error Extract Images:  %s \n", err)
		}
		for _, i := range imageList {
			// "i" will be a images.Image
			fmt.Printf("images is %v \n", i)
		}
		return false, err
	})
}
