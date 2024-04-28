package service

import (
	"encoding/json"
	"fmt"
	"github.com/kube-stack/multicloud_service/src/codegen/openstack"
	"testing"
)

func TestOpenstackCreateVolume(t *testing.T) {
	service := InitByOpenstackType("volumev3", "3.0")
	request := openstack.CreateBlockstorageV3VolumesRequest{}

	request.Opts.Name = "volume-test"
	request.Opts.Size = 1
	request.Opts.VolumeType = "iscsi"
	request.Opts.AvailabilityZone = "nova"
	request.Opts.Description = "test volume"
	requestByte, err := json.Marshal(request)
	//{"Opts":{"volume_id":"3fd158d7-4800-4c64-9058-05f8d71a0e27","name":"test-create"}}
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}

	//{"volume:":{}}
	resp, err := service.CallCloudAPI("CreateBlockstorageV3Volumes", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackGetVolume(t *testing.T) {
	service := InitByOpenstackType("volumev3", "3.0")
	request := openstack.GetBlockstorageV3VolumesRequest{}
	request.Id = "404d836b-8420-45f2-9308-d5091efa0091"
	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}
	//{"snapshot:":{}}
	resp, err := service.CallCloudAPI("GetBlockstorageV3Volumes", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackUpdateVolume(t *testing.T) {
	service := InitByOpenstackType("volumev3", "3.0")
	request := openstack.UpdateBlockstorageV3VolumesRequest{}
	request.Id = "404d836b-8420-45f2-9308-d5091efa0091"
	newName := "volueme updated"
	request.Opts.Description = &newName
	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}
	//{"Id":"e5e6eb1d-ee04-4c4a-8cf8-0b1bef12ada3","Opts":{"name":"updated_snapshot"}}
	resp, err := service.CallCloudAPI("UpdateBlockstorageV3Volumes", requestByte)
	if err != nil {
		t.Error(err)
	}
	//{"snapshot":{"created_at":"2022-12-12T14:48:54.000000","description":null,"id":"e5e6eb1d-ee04-4c4a-8cf8-0b1bef12ada3","metadata":{},"na
	fmt.Println(string(resp))
}

func TestOpenstackDeleteVolume(t *testing.T) {
	service := InitByOpenstackType("volumev3", "3.0")
	request := openstack.DeleteBlockstorageV3VolumesRequest{}
	request.Id = "404d836b-8420-45f2-9308-d5091efa0091"
	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}
	//{"snapshot:":{}}
	resp, err := service.CallCloudAPI("DeleteBlockstorageV3Volumes", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}
