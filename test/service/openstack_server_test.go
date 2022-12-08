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
