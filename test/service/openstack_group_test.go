package service

import (
	"encoding/json"
	"fmt"
	"github.com/kube-stack/multicloud_service/src/codegen/openstack"
	"testing"
)

func TestOpenstackCreateServerGroup(t *testing.T) {
	service := InitByOpenstackType("compute")
	request := openstack.CreateComputeV2ExtensionsServergroupsRequest{}
	request.Opts.Name = "auto scale group"
	request.Opts.Policy = "anti-affinity"
	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	//{"Opts":{"name":"auto scale group","policy":"anti-affinity"}}
	if err != nil {
		t.Error(err)
	}
	//{"Opts":{"name":"auto scale group","policy":"anti-affinity"}}
	resp, err := service.CallCloudAPI("CreateComputeV2ExtensionsServergroups", requestByte)
	//{"server_group":{"id":"5e79ae3f-7601-4be0-a6f6-5e9bd448f89c","members":[],"name":"auto scale group","policy":"anti-affinity","project_id":"aac94320146c464ab84146e35aa61c77","rules":{},"user_id":"f8db2401acfb4c3b98400dac8fa22207"}}
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackGetServerGroup(t *testing.T) {
	service := InitByOpenstackType("compute")
	request := openstack.GetComputeV2ExtensionsServergroupsRequest{}
	request.Id = "5e79ae3f-7601-4be0-a6f6-5e9bd448f89c"

	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}
	resp, err := service.CallCloudAPI("GetComputeV2ExtensionsServergroups", requestByte)
	//{"server_group":{"id":"5e79ae3f-7601-4be0-a6f6-5e9bd448f89c","members":[],"metadata":{},"name":"auto scale group","policies":["anti-affinity"]}}
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackDeleteServerGroup(t *testing.T) {
	service := InitByOpenstackType("compute")
	request := openstack.DeleteComputeV2ExtensionsServergroupsRequest{}
	request.Id = "5e79ae3f-7601-4be0-a6f6-5e9bd448f89c"
	//{"Id":"5e79ae3f-7601-4be0-a6f6-5e9bd448f89c"}
	requestByte, err := json.Marshal(request)
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}
	resp, err := service.CallCloudAPI("DeleteComputeV2ExtensionsServergroups", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}
