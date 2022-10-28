package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/blockstorage/extensions/quotasets"
)
//request struct for the GetBlockstorageExtensionsQuotasets
type GetBlockstorageExtensionsQuotasetsRequest struct{
    ProjectID string
}

func NewGetBlockstorageExtensionsQuotasetsRequest()*GetBlockstorageExtensionsQuotasetsRequest{
    return &GetBlockstorageExtensionsQuotasetsRequest{}
}

//response struct for the GetBlockstorageExtensionsQuotasets
type GetBlockstorageExtensionsQuotasetsResponse struct{
    GetResult quotasets.GetResult
}

func NewGetBlockstorageExtensionsQuotasetsResponse(getResult quotasets.GetResult,)*GetBlockstorageExtensionsQuotasetsResponse {
    return &GetBlockstorageExtensionsQuotasetsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetBlockstorageExtensionsQuotasets(req *GetBlockstorageExtensionsQuotasetsRequest)(*GetBlockstorageExtensionsQuotasetsResponse){
    return NewGetBlockstorageExtensionsQuotasetsResponse(quotasets.Get(oc.client,req.ProjectID, ))

}
//request struct for the GetDefaultsBlockstorageExtensionsQuotasets
type GetDefaultsBlockstorageExtensionsQuotasetsRequest struct{
    ProjectID string
}

func NewGetDefaultsBlockstorageExtensionsQuotasetsRequest()*GetDefaultsBlockstorageExtensionsQuotasetsRequest{
    return &GetDefaultsBlockstorageExtensionsQuotasetsRequest{}
}

//response struct for the GetDefaultsBlockstorageExtensionsQuotasets
type GetDefaultsBlockstorageExtensionsQuotasetsResponse struct{
    GetResult quotasets.GetResult
}

func NewGetDefaultsBlockstorageExtensionsQuotasetsResponse(getResult quotasets.GetResult,)*GetDefaultsBlockstorageExtensionsQuotasetsResponse {
    return &GetDefaultsBlockstorageExtensionsQuotasetsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetDefaultsBlockstorageExtensionsQuotasets(req *GetDefaultsBlockstorageExtensionsQuotasetsRequest)(*GetDefaultsBlockstorageExtensionsQuotasetsResponse){
    return NewGetDefaultsBlockstorageExtensionsQuotasetsResponse(quotasets.GetDefaults(oc.client,req.ProjectID, ))

}
//request struct for the GetUsageBlockstorageExtensionsQuotasets
type GetUsageBlockstorageExtensionsQuotasetsRequest struct{
    ProjectID string
}

func NewGetUsageBlockstorageExtensionsQuotasetsRequest()*GetUsageBlockstorageExtensionsQuotasetsRequest{
    return &GetUsageBlockstorageExtensionsQuotasetsRequest{}
}

//response struct for the GetUsageBlockstorageExtensionsQuotasets
type GetUsageBlockstorageExtensionsQuotasetsResponse struct{
    GetUsageResult quotasets.GetUsageResult
}

func NewGetUsageBlockstorageExtensionsQuotasetsResponse(getUsageResult quotasets.GetUsageResult,)*GetUsageBlockstorageExtensionsQuotasetsResponse {
    return &GetUsageBlockstorageExtensionsQuotasetsResponse{
            GetUsageResult:getUsageResult,
    }
}

// action function
func (oc *OpenstackClient) GetUsageBlockstorageExtensionsQuotasets(req *GetUsageBlockstorageExtensionsQuotasetsRequest)(*GetUsageBlockstorageExtensionsQuotasetsResponse){
    return NewGetUsageBlockstorageExtensionsQuotasetsResponse(quotasets.GetUsage(oc.client,req.ProjectID, ))

}
//request struct for the UpdateBlockstorageExtensionsQuotasets
type UpdateBlockstorageExtensionsQuotasetsRequest struct{
    ProjectID string
    Opts quotasets.UpdateOptsBuilder
}

func NewUpdateBlockstorageExtensionsQuotasetsRequest()*UpdateBlockstorageExtensionsQuotasetsRequest{
    return &UpdateBlockstorageExtensionsQuotasetsRequest{}
}

//response struct for the UpdateBlockstorageExtensionsQuotasets
type UpdateBlockstorageExtensionsQuotasetsResponse struct{
    UpdateResult quotasets.UpdateResult
}

func NewUpdateBlockstorageExtensionsQuotasetsResponse(updateResult quotasets.UpdateResult,)*UpdateBlockstorageExtensionsQuotasetsResponse {
    return &UpdateBlockstorageExtensionsQuotasetsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateBlockstorageExtensionsQuotasets(req *UpdateBlockstorageExtensionsQuotasetsRequest)(*UpdateBlockstorageExtensionsQuotasetsResponse){
    return NewUpdateBlockstorageExtensionsQuotasetsResponse(quotasets.Update(oc.client,req.ProjectID,req.Opts, ))

}
//request struct for the DeleteBlockstorageExtensionsQuotasets
type DeleteBlockstorageExtensionsQuotasetsRequest struct{
    ProjectID string
}

func NewDeleteBlockstorageExtensionsQuotasetsRequest()*DeleteBlockstorageExtensionsQuotasetsRequest{
    return &DeleteBlockstorageExtensionsQuotasetsRequest{}
}

//response struct for the DeleteBlockstorageExtensionsQuotasets
type DeleteBlockstorageExtensionsQuotasetsResponse struct{
    DeleteResult quotasets.DeleteResult
}

func NewDeleteBlockstorageExtensionsQuotasetsResponse(deleteResult quotasets.DeleteResult,)*DeleteBlockstorageExtensionsQuotasetsResponse {
    return &DeleteBlockstorageExtensionsQuotasetsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteBlockstorageExtensionsQuotasets(req *DeleteBlockstorageExtensionsQuotasetsRequest)(*DeleteBlockstorageExtensionsQuotasetsResponse){
    return NewDeleteBlockstorageExtensionsQuotasetsResponse(quotasets.Delete(oc.client,req.ProjectID, ))

}