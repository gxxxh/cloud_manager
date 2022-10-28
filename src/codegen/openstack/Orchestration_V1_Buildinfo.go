package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/orchestration/v1/buildinfo"
)
//request struct for the GetOrchestrationV1Buildinfo
type GetOrchestrationV1BuildinfoRequest struct{
}

func NewGetOrchestrationV1BuildinfoRequest()*GetOrchestrationV1BuildinfoRequest{
    return &GetOrchestrationV1BuildinfoRequest{}
}

//response struct for the GetOrchestrationV1Buildinfo
type GetOrchestrationV1BuildinfoResponse struct{
    GetResult buildinfo.GetResult
}

func NewGetOrchestrationV1BuildinfoResponse(getResult buildinfo.GetResult,)*GetOrchestrationV1BuildinfoResponse {
    return &GetOrchestrationV1BuildinfoResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetOrchestrationV1Buildinfo(req *GetOrchestrationV1BuildinfoRequest)(*GetOrchestrationV1BuildinfoResponse){
    return NewGetOrchestrationV1BuildinfoResponse(buildinfo.Get(oc.client, ))

}