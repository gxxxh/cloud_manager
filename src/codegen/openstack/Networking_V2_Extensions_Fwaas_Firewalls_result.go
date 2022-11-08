package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/firewalls"
)

// extract response info from pager for ListNetworkingV2ExtensionsFwaasFirewalls
func ExtractListNetworkingV2ExtensionsFwaasFirewallsResponse(response *ListNetworkingV2ExtensionsFwaasFirewallsResponse) ([]firewalls.Firewall, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return firewalls.ExtractFirewalls(page)
}

// call result's extract function
func ExtractCreateNetworkingV2ExtensionsFwaasFirewallsResponse(response *CreateNetworkingV2ExtensionsFwaasFirewallsResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// call result's extract function
func ExtractGetNetworkingV2ExtensionsFwaasFirewallsResponse(response *GetNetworkingV2ExtensionsFwaasFirewallsResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractUpdateNetworkingV2ExtensionsFwaasFirewallsResponse(response *UpdateNetworkingV2ExtensionsFwaasFirewallsResponse) (interface{}, error) {
	return response.UpdateResult.Body, response.UpdateResult.Err
}

// call result's extract function
func ExtractDeleteNetworkingV2ExtensionsFwaasFirewallsResponse(response *DeleteNetworkingV2ExtensionsFwaasFirewallsResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}