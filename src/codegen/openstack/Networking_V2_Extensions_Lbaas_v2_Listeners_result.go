package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/listeners"
)

// extract response info from pager for ListNetworkingV2ExtensionsLbaas_v2Listeners
func ExtractListNetworkingV2ExtensionsLbaas_v2ListenersResponse(response *ListNetworkingV2ExtensionsLbaas_v2ListenersResponse) ([]listeners.Listener, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return listeners.ExtractListeners(page)
}

// call result's extract function
func ExtractCreateNetworkingV2ExtensionsLbaas_v2ListenersResponse(response *CreateNetworkingV2ExtensionsLbaas_v2ListenersResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// call result's extract function
func ExtractGetNetworkingV2ExtensionsLbaas_v2ListenersResponse(response *GetNetworkingV2ExtensionsLbaas_v2ListenersResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractUpdateNetworkingV2ExtensionsLbaas_v2ListenersResponse(response *UpdateNetworkingV2ExtensionsLbaas_v2ListenersResponse) (interface{}, error) {
	return response.UpdateResult.Body, response.UpdateResult.Err
}

// call result's extract function
func ExtractDeleteNetworkingV2ExtensionsLbaas_v2ListenersResponse(response *DeleteNetworkingV2ExtensionsLbaas_v2ListenersResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}