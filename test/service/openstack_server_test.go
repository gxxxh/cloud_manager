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
	//{"Opts":{"name":"test-service-create","imageRef":"952b386b-6f30-46f6-b019-f522b157aa3a","flavorRef":"3"}}
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}
	resp, err := service.CallCloudAPI("CreateComputeV2Servers", requestByte)
	//{"server":{"OS-DCF:diskConfig":"MANUAL","adminPass":"ZTodV26JvRt5","id":"4b7fd536-147c-4d1e-893f-58a1ed337e98","links":[{"href":"http://133.133.135.136:8774/v2.1/aac94320146c464ab84146e35aa61c77/servers/4b7fd536-147c-4d1e-893f-58a1ed337e98","rel":"self"},{"href":"http://133.133.135.136:8774/aac94320146c464ab84146e35aa61c77/servers/4b7fd536-147c-4d1e-893f-58a1ed337e98","rel":"bookmark"}],"security_groups":[{"name":"default"}]}}
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackGetServer(t *testing.T) {
	service := InitByOpenstackType("compute")
	request := openstack.GetComputeV2ServersRequest{}
	request.Id = "4b7fd536-147c-4d1e-893f-58a1ed337e98"
	//{"Id":"4b7fd536-147c-4d1e-893f-58a1ed337e98"}
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
	request.Id = "4b7fd536-147c-4d1e-893f-58a1ed337e98"
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

	request.Id = "4b7fd536-147c-4d1e-893f-58a1ed337e98"
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
