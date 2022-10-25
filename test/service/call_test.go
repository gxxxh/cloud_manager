package service

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"os"
	"reflect"
	"testing"
)

func TestAliyunSDK(t *testing.T) {
	regionId := ""        //need to be added
	accessKeyId := ""     //need to be added
	accessKeySecret := "" //need to be added
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
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "",
		Username:         "{}",
		Password:         "{}",
	}
	provider, _ := openstack.AuthenticatedClient(opts)
	client, _ := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	server, _ := servers.Create(client, servers.CreateOpts{
		Name:      "My new server!",
		FlavorRef: "flavor_id",
		ImageRef:  "image_id",
	}).Extract()
	fmt.Println(server)
}
