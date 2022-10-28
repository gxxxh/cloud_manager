package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/limits"
)
//request struct for the GetComputeV2ExtensionsLimits
type GetComputeV2ExtensionsLimitsRequest struct{
    Opts limits.GetOptsBuilder
}

func NewGetComputeV2ExtensionsLimitsRequest()*GetComputeV2ExtensionsLimitsRequest{
    return &GetComputeV2ExtensionsLimitsRequest{}
}

//response struct for the GetComputeV2ExtensionsLimits
type GetComputeV2ExtensionsLimitsResponse struct{
    GetResult limits.GetResult
}

func NewGetComputeV2ExtensionsLimitsResponse(getResult limits.GetResult,)*GetComputeV2ExtensionsLimitsResponse {
    return &GetComputeV2ExtensionsLimitsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetComputeV2ExtensionsLimits(req *GetComputeV2ExtensionsLimitsRequest)(*GetComputeV2ExtensionsLimitsResponse){
    return NewGetComputeV2ExtensionsLimitsResponse(limits.Get(oc.client,req.Opts, ))

}