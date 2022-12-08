package service

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	"github.com/kube-stack/multicloud_service/src/codegen/openstack"
	"testing"
)

func TestOpenstackCreateRouter(t *testing.T) {
	service := InitByOpenstackType("network")
	request := openstack.CreateNetworkingV2ExtensionsLayer3RoutersRequest{}

	asu := false
	enableSNAT := false
	efi := []routers.ExternalFixedIP{
		{
			SubnetID: "21ed2f62-39bd-430c-900a-62728cb4fa8d",
		},
	}
	gwi := routers.GatewayInfo{
		NetworkID:        "a3176333-3df6-480f-af2b-dc5e02ea1aa0",
		EnableSNAT:       &enableSNAT,
		ExternalFixedIPs: efi,
	}

	request.Opts = routers.CreateOpts{
		Name:         "foo_router",
		AdminStateUp: &asu,
		GatewayInfo:  &gwi,
		//AvailabilityZoneHints: []string{"zone1", "zone2"},
	}
	requestByte, err := json.Marshal(request)
	//{"Opts":{"name":"foo_router","admin_state_up":false,"external_gateway_info":{"network_id":"a3176333-3df6-480f-af2b-dc5e02ea1aa0","enable_snat":false,"external_fixed_ips":[{"subnet_id":"21ed2f62-39bd-430c-900a-62728cb4fa8d"}]}}}
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}

	//{"router":{"admin_state_up":false,"created_at":"2022-12-08T12:38:52Z","descrip
	resp, err := service.CallCloudAPI("CreateNetworkingV2ExtensionsLayer3Routers", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackGetRouter(t *testing.T) {
	service := InitByOpenstackType("network")
	request := openstack.GetNetworkingV2ExtensionsLayer3RoutersRequest{}
	request.Id = "928f9b5e-0bce-4a99-a653-c341b866f92a"
	requestByte, err := json.Marshal(request)
	//{"Id":"928f9b5e-0bce-4a99-a653-c341b866f92a"}
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}

	//{"router":{"admin_state_up":true,"created_at":"2022-10-08T10:21:40Z","descri
	resp, err := service.CallCloudAPI("GetNetworkingV2ExtensionsLayer3Routers", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackUpdateRouter(t *testing.T) {
	service := InitByOpenstackType("network")
	request := openstack.UpdateNetworkingV2ExtensionsLayer3RoutersRequest{}
	request.Id = "928f9b5e-0bce-4a99-a653-c341b866f92a"
	request.Opts.Name = "update-router-name"
	requestByte, err := json.Marshal(request)
	//{"Id":"928f9b5e-0bce-4a99-a653-c341b866f92a","Opts":{"name":"update-router-name"}}
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}

	//{"router":{"admin_state_up":true,"created_at":"2022-10-08T10:21:40Z","descri
	resp, err := service.CallCloudAPI("UpdateNetworkingV2ExtensionsLayer3Routers", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}

func TestOpenstackDeleteRouter(t *testing.T) {
	service := InitByOpenstackType("network")
	request := openstack.DeleteNetworkingV2ExtensionsLayer3RoutersRequest{}
	request.Id = "928f9b5e-0bce-4a99-a653-c341b866f92a"
	requestByte, err := json.Marshal(request)
	//{"Id":"928f9b5e-0bce-4a99-a653-c341b866f92a"}
	fmt.Println(string(requestByte))
	if err != nil {
		t.Error(err)
	}

	resp, err := service.CallCloudAPI("DeleteNetworkingV2ExtensionsLayer3Routers", requestByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}
