package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/resetnetwork"
)
//request struct for the ResetNetworkComputeV2ExtensionsResetnetwork
type ResetNetworkComputeV2ExtensionsResetnetworkRequest struct{
    Id string
}

func NewResetNetworkComputeV2ExtensionsResetnetworkRequest()*ResetNetworkComputeV2ExtensionsResetnetworkRequest{
    return &ResetNetworkComputeV2ExtensionsResetnetworkRequest{}
}

//response struct for the ResetNetworkComputeV2ExtensionsResetnetwork
type ResetNetworkComputeV2ExtensionsResetnetworkResponse struct{
    ResetResult resetnetwork.ResetResult
}

func NewResetNetworkComputeV2ExtensionsResetnetworkResponse(resetResult resetnetwork.ResetResult,)*ResetNetworkComputeV2ExtensionsResetnetworkResponse {
    return &ResetNetworkComputeV2ExtensionsResetnetworkResponse{
            ResetResult:resetResult,
    }
}

// action function
func (oc *OpenstackClient) ResetNetworkComputeV2ExtensionsResetnetwork(req *ResetNetworkComputeV2ExtensionsResetnetworkRequest)(*ResetNetworkComputeV2ExtensionsResetnetworkResponse){
    return NewResetNetworkComputeV2ExtensionsResetnetworkResponse(resetnetwork.ResetNetwork(oc.client,req.Id, ))

}