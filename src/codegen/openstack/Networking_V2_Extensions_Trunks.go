package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/trunks"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateNetworkingV2ExtensionsTrunks
type CreateNetworkingV2ExtensionsTrunksRequest struct{
    Opts trunks.CreateOptsBuilder
}

func NewCreateNetworkingV2ExtensionsTrunksRequest()*CreateNetworkingV2ExtensionsTrunksRequest{
    return &CreateNetworkingV2ExtensionsTrunksRequest{}
}

//response struct for the CreateNetworkingV2ExtensionsTrunks
type CreateNetworkingV2ExtensionsTrunksResponse struct{
    CreateResult trunks.CreateResult
}

func NewCreateNetworkingV2ExtensionsTrunksResponse(createResult trunks.CreateResult,)*CreateNetworkingV2ExtensionsTrunksResponse {
    return &CreateNetworkingV2ExtensionsTrunksResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsTrunks(req *CreateNetworkingV2ExtensionsTrunksRequest)(*CreateNetworkingV2ExtensionsTrunksResponse){
    return NewCreateNetworkingV2ExtensionsTrunksResponse(trunks.Create(oc.client,req.Opts, ))

}
//request struct for the DeleteNetworkingV2ExtensionsTrunks
type DeleteNetworkingV2ExtensionsTrunksRequest struct{
    Id string
}

func NewDeleteNetworkingV2ExtensionsTrunksRequest()*DeleteNetworkingV2ExtensionsTrunksRequest{
    return &DeleteNetworkingV2ExtensionsTrunksRequest{}
}

//response struct for the DeleteNetworkingV2ExtensionsTrunks
type DeleteNetworkingV2ExtensionsTrunksResponse struct{
    DeleteResult trunks.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsTrunksResponse(deleteResult trunks.DeleteResult,)*DeleteNetworkingV2ExtensionsTrunksResponse {
    return &DeleteNetworkingV2ExtensionsTrunksResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsTrunks(req *DeleteNetworkingV2ExtensionsTrunksRequest)(*DeleteNetworkingV2ExtensionsTrunksResponse){
    return NewDeleteNetworkingV2ExtensionsTrunksResponse(trunks.Delete(oc.client,req.Id, ))

}
//request struct for the ListNetworkingV2ExtensionsTrunks
type ListNetworkingV2ExtensionsTrunksRequest struct{
    Opts trunks.ListOptsBuilder
}

func NewListNetworkingV2ExtensionsTrunksRequest()*ListNetworkingV2ExtensionsTrunksRequest{
    return &ListNetworkingV2ExtensionsTrunksRequest{}
}

//response struct for the ListNetworkingV2ExtensionsTrunks
type ListNetworkingV2ExtensionsTrunksResponse struct{
    Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsTrunksResponse(pager pagination.Pager,)*ListNetworkingV2ExtensionsTrunksResponse {
    return &ListNetworkingV2ExtensionsTrunksResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsTrunks(req *ListNetworkingV2ExtensionsTrunksRequest)(*ListNetworkingV2ExtensionsTrunksResponse){
    return NewListNetworkingV2ExtensionsTrunksResponse(trunks.List(oc.client,req.Opts, ))

}
//request struct for the GetNetworkingV2ExtensionsTrunks
type GetNetworkingV2ExtensionsTrunksRequest struct{
    Id string
}

func NewGetNetworkingV2ExtensionsTrunksRequest()*GetNetworkingV2ExtensionsTrunksRequest{
    return &GetNetworkingV2ExtensionsTrunksRequest{}
}

//response struct for the GetNetworkingV2ExtensionsTrunks
type GetNetworkingV2ExtensionsTrunksResponse struct{
    GetResult trunks.GetResult
}

func NewGetNetworkingV2ExtensionsTrunksResponse(getResult trunks.GetResult,)*GetNetworkingV2ExtensionsTrunksResponse {
    return &GetNetworkingV2ExtensionsTrunksResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsTrunks(req *GetNetworkingV2ExtensionsTrunksRequest)(*GetNetworkingV2ExtensionsTrunksResponse){
    return NewGetNetworkingV2ExtensionsTrunksResponse(trunks.Get(oc.client,req.Id, ))

}
//request struct for the UpdateNetworkingV2ExtensionsTrunks
type UpdateNetworkingV2ExtensionsTrunksRequest struct{
    Id string
    Opts trunks.UpdateOptsBuilder
}

func NewUpdateNetworkingV2ExtensionsTrunksRequest()*UpdateNetworkingV2ExtensionsTrunksRequest{
    return &UpdateNetworkingV2ExtensionsTrunksRequest{}
}

//response struct for the UpdateNetworkingV2ExtensionsTrunks
type UpdateNetworkingV2ExtensionsTrunksResponse struct{
    UpdateResult trunks.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsTrunksResponse(updateResult trunks.UpdateResult,)*UpdateNetworkingV2ExtensionsTrunksResponse {
    return &UpdateNetworkingV2ExtensionsTrunksResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsTrunks(req *UpdateNetworkingV2ExtensionsTrunksRequest)(*UpdateNetworkingV2ExtensionsTrunksResponse){
    return NewUpdateNetworkingV2ExtensionsTrunksResponse(trunks.Update(oc.client,req.Id,req.Opts, ))

}
//request struct for the GetSubportsNetworkingV2ExtensionsTrunks
type GetSubportsNetworkingV2ExtensionsTrunksRequest struct{
    Id string
}

func NewGetSubportsNetworkingV2ExtensionsTrunksRequest()*GetSubportsNetworkingV2ExtensionsTrunksRequest{
    return &GetSubportsNetworkingV2ExtensionsTrunksRequest{}
}

//response struct for the GetSubportsNetworkingV2ExtensionsTrunks
type GetSubportsNetworkingV2ExtensionsTrunksResponse struct{
    GetSubportsResult trunks.GetSubportsResult
}

func NewGetSubportsNetworkingV2ExtensionsTrunksResponse(getSubportsResult trunks.GetSubportsResult,)*GetSubportsNetworkingV2ExtensionsTrunksResponse {
    return &GetSubportsNetworkingV2ExtensionsTrunksResponse{
            GetSubportsResult:getSubportsResult,
    }
}

// action function
func (oc *OpenstackClient) GetSubportsNetworkingV2ExtensionsTrunks(req *GetSubportsNetworkingV2ExtensionsTrunksRequest)(*GetSubportsNetworkingV2ExtensionsTrunksResponse){
    return NewGetSubportsNetworkingV2ExtensionsTrunksResponse(trunks.GetSubports(oc.client,req.Id, ))

}
//request struct for the AddSubportsNetworkingV2ExtensionsTrunks
type AddSubportsNetworkingV2ExtensionsTrunksRequest struct{
    Id string
    Opts trunks.AddSubportsOptsBuilder
}

func NewAddSubportsNetworkingV2ExtensionsTrunksRequest()*AddSubportsNetworkingV2ExtensionsTrunksRequest{
    return &AddSubportsNetworkingV2ExtensionsTrunksRequest{}
}

//response struct for the AddSubportsNetworkingV2ExtensionsTrunks
type AddSubportsNetworkingV2ExtensionsTrunksResponse struct{
    UpdateSubportsResult trunks.UpdateSubportsResult
}

func NewAddSubportsNetworkingV2ExtensionsTrunksResponse(updateSubportsResult trunks.UpdateSubportsResult,)*AddSubportsNetworkingV2ExtensionsTrunksResponse {
    return &AddSubportsNetworkingV2ExtensionsTrunksResponse{
            UpdateSubportsResult:updateSubportsResult,
    }
}

// action function
func (oc *OpenstackClient) AddSubportsNetworkingV2ExtensionsTrunks(req *AddSubportsNetworkingV2ExtensionsTrunksRequest)(*AddSubportsNetworkingV2ExtensionsTrunksResponse){
    return NewAddSubportsNetworkingV2ExtensionsTrunksResponse(trunks.AddSubports(oc.client,req.Id,req.Opts, ))

}
//request struct for the RemoveSubportsNetworkingV2ExtensionsTrunks
type RemoveSubportsNetworkingV2ExtensionsTrunksRequest struct{
    Id string
    Opts trunks.RemoveSubportsOptsBuilder
}

func NewRemoveSubportsNetworkingV2ExtensionsTrunksRequest()*RemoveSubportsNetworkingV2ExtensionsTrunksRequest{
    return &RemoveSubportsNetworkingV2ExtensionsTrunksRequest{}
}

//response struct for the RemoveSubportsNetworkingV2ExtensionsTrunks
type RemoveSubportsNetworkingV2ExtensionsTrunksResponse struct{
    UpdateSubportsResult trunks.UpdateSubportsResult
}

func NewRemoveSubportsNetworkingV2ExtensionsTrunksResponse(updateSubportsResult trunks.UpdateSubportsResult,)*RemoveSubportsNetworkingV2ExtensionsTrunksResponse {
    return &RemoveSubportsNetworkingV2ExtensionsTrunksResponse{
            UpdateSubportsResult:updateSubportsResult,
    }
}

// action function
func (oc *OpenstackClient) RemoveSubportsNetworkingV2ExtensionsTrunks(req *RemoveSubportsNetworkingV2ExtensionsTrunksRequest)(*RemoveSubportsNetworkingV2ExtensionsTrunksResponse){
    return NewRemoveSubportsNetworkingV2ExtensionsTrunksResponse(trunks.RemoveSubports(oc.client,req.Id,req.Opts, ))

}