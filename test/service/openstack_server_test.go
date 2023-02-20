package service

import (
	"encoding/json"
	"fmt"
	"github.com/kube-stack/multicloud_service/src/codegen/openstack"
	"testing"
)

func TestOpenstackCreateServer(t *testing.T) {
	service := InitByOpenstackType("compute")
	request := openstack.CreateComputeV2ServersRequest{}
	request.Opts.Name = "test-service-create"
	request.Opts.ImageRef = "952b386b-6f30-46f6-b019-f522b157aa3a"
	request.Opts.FlavorRef = "3"
	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}
	resp, err := service.CallCloudAPI("CreateComputeV2Servers", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackGetServer(t *testing.T) {
	service := InitByOpenstackType("compute")
	request := openstack.GetComputeV2ServersRequest{}
	request.Id = "65992d97-a29c-4b6c-b2b4-c12e5bfd475f"

	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}
	resp, err := service.CallCloudAPI("GetComputeV2Servers", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackServerResize(t *testing.T) {
	service := InitByOpenstackType("compute")
	request := openstack.ResizeComputeV2ServersRequest{}
	request.Id = "8cec0165-f1a9-4224-8e45-aca635c84562"
	request.Opts.FlavorRef = "4"
	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	//{"Id":"8cec0165-f1a9-4224-8e45-aca635c84562","Opts":{"flavorRef":"4"}}
	if err != nil {
		t.Error(err)
	}
	resp, err := service.CallCloudAPI("ResizeComputeV2Servers", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestGetFlavor(t *testing.T) {
	service := InitByOpenstackType("compute")
	request := openstack.GetComputeV2FlavorsRequest{}

	request.Id = "60ac07db-8435-4526-bbda-7bb91365f908"
	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	//{"Id":"8cec0165-f1a9-4224-8e45-aca635c84562","Opts":{"flavorRef":"4"}}
	if err != nil {
		t.Error(err)
	}
	resp, err := service.CallCloudAPI("GetComputeV2Flavors", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}
