package service

import (
	"encoding/json"
	"fmt"
	"github.com/kube-stack/multicloud_service/src/codegen/openstack"
	"testing"
)

func TestOpenstackCreateNetwork(t *testing.T) {
	service := InitByOpenstackType("network")
	request := openstack.CreateNetworkingV2NetworksRequest{}
	request.Opts.Name = "test-create-private"
	itrue := true
	request.Opts.AdminStateUp = &itrue
	requestByte, err := json.Marshal(request)
	//{"Opts":{"admin_state_up":true,"name":"test-create-private"}}
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}

	//{"network":{"admin_state_up":true,"availability_zone_hints":[],"availab
	resp, err := service.CallCloudAPI("CreateNetworkingV2Networks", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackGetNetwork(t *testing.T) {
	service := InitByOpenstackType("network")
	request := openstack.GetNetworkingV2NetworksRequest{}
	request.Id = "c8af2a45-8226-4c61-8a7f-d5b556f9215f"
	requestByte, err := json.Marshal(request)
	//{"Id":"c8af2a45-8226-4c61-8a7f-d5b556f9215f"}
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}

	//{"network":{"admin_state_up":true,"availability_zone_hints":[],"availab
	resp, err := service.CallCloudAPI("GetNetworkingV2Networks", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackUpdateNetwork(t *testing.T) {
	service := InitByOpenstackType("network")
	request := openstack.UpdateNetworkingV2NetworksRequest{}
	request.NetworkID = "c8af2a45-8226-4c61-8a7f-d5b556f9215f"
	ifalse := false
	request.Opts.AdminStateUp = &ifalse
	//{"NetworkID":"c8af2a45-8226-4c61-8a7f-d5b556f9215f","Opts":{"admin_state_up":false}}
	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}

	//{"network":{"admin_state_up":true,"availability_zone_hints":[],"availab
	resp, err := service.CallCloudAPI("UpdateNetworkingV2Networks", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackDeleteNetwork(t *testing.T) {
	service := InitByOpenstackType("network")
	request := openstack.DeleteNetworkingV2NetworksRequest{}
	request.NetworkID = "c8af2a45-8226-4c61-8a7f-d5b556f9215f"
	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}

	//{"network":{"admin_state_up":true,"availability_zone_hints":[],"availab
	resp, err := service.CallCloudAPI("DeleteNetworkingV2Networks", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}
