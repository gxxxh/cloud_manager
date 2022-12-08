package service

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	"github.com/kube-stack/multicloud_service/src/codegen/openstack"
	"github.com/kube-stack/multicloud_service/src/service"
	"testing"
)

func InitByOpenstackType(openstackClientType string) *service.MultiCloudService {
	authInfo := map[string]string{
		"projectName":         "admin",
		"domainName":          "Default",
		"identityEndpoint":    "http://133.133.135.136:5000/v3",
		"username":            "admin",
		"password":            "ef1aa1ad78c442e1",
		"Region":              "RegionOne",
		"openstackClientType": openstackClientType,
		"cloudType":           "openstack",
	}
	mcm, err := service.NewMultiCloudService(authInfo)
	if err != nil {
		panic(err)
	}
	return mcm
}

func TestOpenstackCreateImage(t *testing.T) {
	service := InitByOpenstackType("image")
	request := openstack.CreateImageserviceV2ImagesRequest{}
	request.Opts.Name = "test-create"
	request.Opts.ID = "e7db3b45-8db7-47ad-8109-3fb55c2c24fd"
	request.Opts.Properties = map[string]string{
		"architecture": "x86_64",
	}
	request.Opts.Tags = []string{"ubuntu", "quantal"}
	requestByte, err := json.Marshal(request)
	if err != nil {
		t.Error(err)
	}
	resp, err := service.CallCloudAPI("CreateImageserviceV2Images", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}
func TestOpenstackGetImage(t *testing.T) {
	service := InitByOpenstackType("image")
	request := openstack.GetImageserviceV2ImagesRequest{}
	request.Id = "e7db3b45-8db7-47ad-8109-3fb55c2c24fd"
	requestByte, err := json.Marshal(request)
	if err != nil {
		t.Error(err)
	}
	resp, err := service.CallCloudAPI("GetImageserviceV2Images", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackUpdateImage(t *testing.T) {
	service := InitByOpenstackType("image")
	request := openstack.UpdateImageserviceV2ImagesRequest{}
	request.Id = "e7db3b45-8db7-47ad-8109-3fb55c2c24fd"
	request.Opts = []images.Patch{images.ReplaceImageName{"image-create-update"}}
	requestByte, err := json.Marshal(request)
	//{"Id":"e7db3b45-8db7-47ad-8109-3fb55c2c24fd","Opts":[{"NewName":"image-create-update"}]}
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}
	//todo using updateProperties to write the info into it
	resp, err := service.CallCloudAPI("UpdateImageserviceV2Images", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackDeleteImage(t *testing.T) {
	service := InitByOpenstackType("image")
	request := openstack.DeleteImageserviceV2ImagesRequest{}
	request.Id = "e7db3b45-8db7-47ad-8109-3fb55c2c24fd"
	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}
	//todo using updateProperties to write the info into it
	resp, err := service.CallCloudAPI("DeleteImageserviceV2Images", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}
