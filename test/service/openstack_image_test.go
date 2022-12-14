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
	//request.Opts.ID = "e7db3b45-8db7-47ad-8109-3fb55c2c24fe"
	request.Opts.Properties = map[string]string{
		"architecture": "x86_64",
	}
	request.Opts.Tags = []string{"ubuntu", "quantal"}
	requestByte, err := json.Marshal(request)
	if err != nil {
		t.Error(err)
	}
	// {"Opts":{"name":"test-create","tags":["ubuntu","quantal"]}}
	fmt.Println(string(requestByte))
	resp, err := service.CallCloudAPI("CreateImageserviceV2Images", requestByte)
	if err != nil {
		t.Error(err)
	}
	//{"checksum":null,"container_format":null,"created_at":"2022-12-14T08:40:59Z","disk_format":null,"file":"/v2/images/87911a0d-c9af-406e-9fd2-6e5b80402fee/file","id":"87911a0d-c9af-406e-9fd2-6e5b80402fee","min_disk":0,"min_ram":0,"name":"test-create","os_hash_algo":null,"os_hash_value":null,"os_hidden":false,"owner":"aac94320146c464ab84146e35aa61c77","protected":false,"schema":"/v2/schemas/image","self":"/v2/images/87911a0d-c9af-406e-9fd2-6e5b80402fee","size":null,"status":"queued","tags":["quantal","ubuntu"],"updated_at":"2022-12-14T08:40:59Z","virtual_size":null,"visibility":"shared"}
	fmt.Println(string(resp))
}
func TestOpenstackGetImage(t *testing.T) {
	service := InitByOpenstackType("image")
	request := openstack.GetImageserviceV2ImagesRequest{}
	request.Id = "87911a0d-c9af-406e-9fd2-6e5b80402fee"
	requestByte, err := json.Marshal(request)
	if err != nil {
		t.Error(err)
	}
	//{"Id":"e7db3b45-8db7-47ad-8109-3fb55c2c24fe"}
	fmt.Println(string(requestByte))
	resp, err := service.CallCloudAPI("GetImageserviceV2Images", requestByte)
	if err != nil {
		t.Error(err)
	}
	//{"checksum":null,"container_format":null,"created_at":"2022-12-12T14:24:25Z","disk_format"
	fmt.Println(string(resp))
}

func TestOpenstackUpdateImage(t *testing.T) {
	service := InitByOpenstackType("image")
	request := openstack.UpdateImageserviceV2ImagesRequest{}
	request.Id = "87911a0d-c9af-406e-9fd2-6e5b80402fee"
	updateVisibility := openstack.UpdateImageserviceV2ImagesProperty{
		Op:    "replace",
		Name:  "visibility",
		Value: images.ImageVisibilityPublic,
	}
	//newHidden := true
	//updateImageHidden := images.UpdateImageProperty{
	//	Op:    "replace",
	//	Name:  "os_hidden",
	//	Value: fmt.Sprintf("%v", newHidden),
	//}
	updateImageName := openstack.UpdateImageserviceV2ImagesProperty{
		Op:    "replace",
		Name:  "name",
		Value: "updated-name",
	}
	//updateImageChecksum := images.UpdateImageProperty{
	//	Op:    "replace",
	//	Name:  "checksum",
	//	Value: "",
	//}

	updateImageTags := openstack.UpdateImageserviceV2ImagesProperty{
		Op:    "replace",
		Name:  "/tags",
		Value: []string{"windows"},
	}
	//updateImageMinDisk := images.UpdateImageProperty{
	//	Op:    "replace",
	//	Name:  "min_disk",
	//	Value: "1",
	//}

	//updateImageMinRam := images.UpdateImageProperty{
	//	Op:    "replace",
	//	Name:  "min_ram",
	//	Value: "1",
	//}
	//updateImageTags := images.UpdateImageProperty{
	//	Op:    "replace",
	//	Name:  "protected",
	//	Value: "true",
	//}
	request.Opts = []openstack.UpdateImageserviceV2ImagesProperty{updateVisibility, updateImageName, updateImageTags}

	requestByte, err := json.Marshal(request)
	//{"Id":"87911a0d-c9af-406e-9fd2-6e5b80402fee","Opts":[{"Op":"replace","Name":"visibility","Value":"public"},{"Op":"replace","Name":"name","Value":"updated-name"},{"Op":"replace","Name":"/tags","Value":["windows"]}]}
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}
	resp, err := service.CallCloudAPI("UpdateImageserviceV2Images", requestByte)
	if err != nil {
		t.Error(err)
	}
	//{"checksum":null,"container_format":null,"created_at":"2022-12-14T08:40:59Z","disk_format":null,"file":"/v2/images/87911a0d-c9af-406e-9fd2-6e5b80402fee/file","id":"87911a0d-c9af-406e-9fd2-6e5b80402fee","min_disk":0,"min_ram":0,"name":"updated-name","os_hash_algo":null,"os_hash_value":null,"os_hidden":false,"owner":"aac94320146c464ab84146e35aa61c77","protected":false,"schema":"/v2/schemas/image","self":"/v2/images/87911a0d-c9af-406e-9fd2-6e5b80402fee","size":null,"status":"queued","tags":["windows"],"updated_at":"2022-12-14T08:50:02Z","virtual_size":null,"visibility":"public"}
	fmt.Println(string(resp))
}

func TestOpenstackDeleteImage(t *testing.T) {
	service := InitByOpenstackType("image")
	request := openstack.DeleteImageserviceV2ImagesRequest{}
	request.Id = "87911a0d-c9af-406e-9fd2-6e5b80402fee"
	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}
	//{"Id":"e7db3b45-8db7-47ad-8109-3fb55c2c24fe"}
	resp, err := service.CallCloudAPI("DeleteImageserviceV2Images", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}
