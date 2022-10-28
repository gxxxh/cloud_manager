package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/diagnostics"
)
//request struct for the GetComputeV2ExtensionsDiagnostics
type GetComputeV2ExtensionsDiagnosticsRequest struct{
    ServerId string
}

func NewGetComputeV2ExtensionsDiagnosticsRequest()*GetComputeV2ExtensionsDiagnosticsRequest{
    return &GetComputeV2ExtensionsDiagnosticsRequest{}
}

//response struct for the GetComputeV2ExtensionsDiagnostics
type GetComputeV2ExtensionsDiagnosticsResponse struct{
    ServerDiagnosticsResult diagnostics.serverDiagnosticsResult
}

func NewGetComputeV2ExtensionsDiagnosticsResponse(serverDiagnosticsResult diagnostics.serverDiagnosticsResult,)*GetComputeV2ExtensionsDiagnosticsResponse {
    return &GetComputeV2ExtensionsDiagnosticsResponse{
            ServerDiagnosticsResult:serverDiagnosticsResult,
    }
}

// action function
func (oc *OpenstackClient) GetComputeV2ExtensionsDiagnostics(req *GetComputeV2ExtensionsDiagnosticsRequest)(*GetComputeV2ExtensionsDiagnosticsResponse){
    return NewGetComputeV2ExtensionsDiagnosticsResponse(diagnostics.Get(oc.client,req.ServerId, ))

}