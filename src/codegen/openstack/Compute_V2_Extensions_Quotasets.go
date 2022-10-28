package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/quotasets"
)
//request struct for the GetComputeV2ExtensionsQuotasets
type GetComputeV2ExtensionsQuotasetsRequest struct{
    TenantID string
}

func NewGetComputeV2ExtensionsQuotasetsRequest()*GetComputeV2ExtensionsQuotasetsRequest{
    return &GetComputeV2ExtensionsQuotasetsRequest{}
}

//response struct for the GetComputeV2ExtensionsQuotasets
type GetComputeV2ExtensionsQuotasetsResponse struct{
    GetResult quotasets.GetResult
}

func NewGetComputeV2ExtensionsQuotasetsResponse(getResult quotasets.GetResult,)*GetComputeV2ExtensionsQuotasetsResponse {
    return &GetComputeV2ExtensionsQuotasetsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetComputeV2ExtensionsQuotasets(req *GetComputeV2ExtensionsQuotasetsRequest)(*GetComputeV2ExtensionsQuotasetsResponse){
    return NewGetComputeV2ExtensionsQuotasetsResponse(quotasets.Get(oc.client,req.TenantID, ))

}
//request struct for the GetDetailComputeV2ExtensionsQuotasets
type GetDetailComputeV2ExtensionsQuotasetsRequest struct{
    TenantID string
}

func NewGetDetailComputeV2ExtensionsQuotasetsRequest()*GetDetailComputeV2ExtensionsQuotasetsRequest{
    return &GetDetailComputeV2ExtensionsQuotasetsRequest{}
}

//response struct for the GetDetailComputeV2ExtensionsQuotasets
type GetDetailComputeV2ExtensionsQuotasetsResponse struct{
    GetDetailResult quotasets.GetDetailResult
}

func NewGetDetailComputeV2ExtensionsQuotasetsResponse(getDetailResult quotasets.GetDetailResult,)*GetDetailComputeV2ExtensionsQuotasetsResponse {
    return &GetDetailComputeV2ExtensionsQuotasetsResponse{
            GetDetailResult:getDetailResult,
    }
}

// action function
func (oc *OpenstackClient) GetDetailComputeV2ExtensionsQuotasets(req *GetDetailComputeV2ExtensionsQuotasetsRequest)(*GetDetailComputeV2ExtensionsQuotasetsResponse){
    return NewGetDetailComputeV2ExtensionsQuotasetsResponse(quotasets.GetDetail(oc.client,req.TenantID, ))

}
//request struct for the UpdateComputeV2ExtensionsQuotasets
type UpdateComputeV2ExtensionsQuotasetsRequest struct{
    TenantID string
    Opts quotasets.UpdateOptsBuilder
}

func NewUpdateComputeV2ExtensionsQuotasetsRequest()*UpdateComputeV2ExtensionsQuotasetsRequest{
    return &UpdateComputeV2ExtensionsQuotasetsRequest{}
}

//response struct for the UpdateComputeV2ExtensionsQuotasets
type UpdateComputeV2ExtensionsQuotasetsResponse struct{
    UpdateResult quotasets.UpdateResult
}

func NewUpdateComputeV2ExtensionsQuotasetsResponse(updateResult quotasets.UpdateResult,)*UpdateComputeV2ExtensionsQuotasetsResponse {
    return &UpdateComputeV2ExtensionsQuotasetsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateComputeV2ExtensionsQuotasets(req *UpdateComputeV2ExtensionsQuotasetsRequest)(*UpdateComputeV2ExtensionsQuotasetsResponse){
    return NewUpdateComputeV2ExtensionsQuotasetsResponse(quotasets.Update(oc.client,req.TenantID,req.Opts, ))

}
//request struct for the DeleteComputeV2ExtensionsQuotasets
type DeleteComputeV2ExtensionsQuotasetsRequest struct{
    TenantID string
}

func NewDeleteComputeV2ExtensionsQuotasetsRequest()*DeleteComputeV2ExtensionsQuotasetsRequest{
    return &DeleteComputeV2ExtensionsQuotasetsRequest{}
}

//response struct for the DeleteComputeV2ExtensionsQuotasets
type DeleteComputeV2ExtensionsQuotasetsResponse struct{
    DeleteResult quotasets.DeleteResult
}

func NewDeleteComputeV2ExtensionsQuotasetsResponse(deleteResult quotasets.DeleteResult,)*DeleteComputeV2ExtensionsQuotasetsResponse {
    return &DeleteComputeV2ExtensionsQuotasetsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteComputeV2ExtensionsQuotasets(req *DeleteComputeV2ExtensionsQuotasetsRequest)(*DeleteComputeV2ExtensionsQuotasetsResponse){
    return NewDeleteComputeV2ExtensionsQuotasetsResponse(quotasets.Delete(oc.client,req.TenantID, ))

}