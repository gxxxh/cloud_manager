package service

import (
	"encoding/json"
	"fmt"
	"github.com/kube-stack/multicloud_service/src/codegen/openstack"
	"testing"
)

func TestOpenstackCreateSnapshot(t *testing.T) {
	service := InitByOpenstackType("volumev3")
	request := openstack.CreateBlockstorageV3SnapshotsRequest{}
	request.Opts.Name = "test-create"
	request.Opts.VolumeID = "3fd158d7-4800-4c64-9058-05f8d71a0e27"
	requestByte, err := json.Marshal(request)
	//{"Opts":{"volume_id":"3fd158d7-4800-4c64-9058-05f8d71a0e27","name":"test-create"}}
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}

	//{"snapshot:":{}}
	resp, err := service.CallCloudAPI("CreateBlockstorageV3Snapshots", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackGetSnapshot(t *testing.T) {
	service := InitByOpenstackType("volumev3")
	request := openstack.GetBlockstorageV3SnapshotsRequest{}
	request.Id = "2b868ade-1c76-42c3-92c7-8aa27fcd096a"
	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}
	//{"snapshot:":{}}
	resp, err := service.CallCloudAPI("GetBlockstorageV3Snapshots", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackUpdateSnapshot(t *testing.T) {
	service := InitByOpenstackType("volumev3")
	request := openstack.UpdateBlockstorageV3SnapshotsRequest{}
	request.Id = "0a509ed0-36e5-49c5-bdfa-ca4ef53d72c6"
	newName := "updated_snapshot"
	request.Opts.Name = &newName
	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}
	//{"snapshot:":{}}
	resp, err := service.CallCloudAPI("DeleteBlockstorageV3Snapshots", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackDeleteSnapshot(t *testing.T) {
	service := InitByOpenstackType("volumev3")
	request := openstack.DeleteBlockstorageV3SnapshotsRequest{}
	request.Id = "0a509ed0-36e5-49c5-bdfa-ca4ef53d72c6"
	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}
	//{"snapshot:":{}}
	resp, err := service.CallCloudAPI("DeleteBlockstorageV3Snapshots", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}
