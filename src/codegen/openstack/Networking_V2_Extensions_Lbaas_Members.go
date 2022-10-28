package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas/members"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListNetworkingV2ExtensionsLbaasMembers
type ListNetworkingV2ExtensionsLbaasMembersRequest struct{
    Opts members.ListOpts
}

func NewListNetworkingV2ExtensionsLbaasMembersRequest()*ListNetworkingV2ExtensionsLbaasMembersRequest{
    return &ListNetworkingV2ExtensionsLbaasMembersRequest{}
}

//response struct for the ListNetworkingV2ExtensionsLbaasMembers
type ListNetworkingV2ExtensionsLbaasMembersResponse struct{
    Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsLbaasMembersResponse(pager pagination.Pager,)*ListNetworkingV2ExtensionsLbaasMembersResponse {
    return &ListNetworkingV2ExtensionsLbaasMembersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsLbaasMembers(req *ListNetworkingV2ExtensionsLbaasMembersRequest)(*ListNetworkingV2ExtensionsLbaasMembersResponse){
    return NewListNetworkingV2ExtensionsLbaasMembersResponse(members.List(oc.client,req.Opts, ))

}
//request struct for the CreateNetworkingV2ExtensionsLbaasMembers
type CreateNetworkingV2ExtensionsLbaasMembersRequest struct{
    Opts members.CreateOptsBuilder
}

func NewCreateNetworkingV2ExtensionsLbaasMembersRequest()*CreateNetworkingV2ExtensionsLbaasMembersRequest{
    return &CreateNetworkingV2ExtensionsLbaasMembersRequest{}
}

//response struct for the CreateNetworkingV2ExtensionsLbaasMembers
type CreateNetworkingV2ExtensionsLbaasMembersResponse struct{
    CreateResult members.CreateResult
}

func NewCreateNetworkingV2ExtensionsLbaasMembersResponse(createResult members.CreateResult,)*CreateNetworkingV2ExtensionsLbaasMembersResponse {
    return &CreateNetworkingV2ExtensionsLbaasMembersResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsLbaasMembers(req *CreateNetworkingV2ExtensionsLbaasMembersRequest)(*CreateNetworkingV2ExtensionsLbaasMembersResponse){
    return NewCreateNetworkingV2ExtensionsLbaasMembersResponse(members.Create(oc.client,req.Opts, ))

}
//request struct for the GetNetworkingV2ExtensionsLbaasMembers
type GetNetworkingV2ExtensionsLbaasMembersRequest struct{
    Id string
}

func NewGetNetworkingV2ExtensionsLbaasMembersRequest()*GetNetworkingV2ExtensionsLbaasMembersRequest{
    return &GetNetworkingV2ExtensionsLbaasMembersRequest{}
}

//response struct for the GetNetworkingV2ExtensionsLbaasMembers
type GetNetworkingV2ExtensionsLbaasMembersResponse struct{
    GetResult members.GetResult
}

func NewGetNetworkingV2ExtensionsLbaasMembersResponse(getResult members.GetResult,)*GetNetworkingV2ExtensionsLbaasMembersResponse {
    return &GetNetworkingV2ExtensionsLbaasMembersResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsLbaasMembers(req *GetNetworkingV2ExtensionsLbaasMembersRequest)(*GetNetworkingV2ExtensionsLbaasMembersResponse){
    return NewGetNetworkingV2ExtensionsLbaasMembersResponse(members.Get(oc.client,req.Id, ))

}
//request struct for the UpdateNetworkingV2ExtensionsLbaasMembers
type UpdateNetworkingV2ExtensionsLbaasMembersRequest struct{
    Id string
    Opts members.UpdateOptsBuilder
}

func NewUpdateNetworkingV2ExtensionsLbaasMembersRequest()*UpdateNetworkingV2ExtensionsLbaasMembersRequest{
    return &UpdateNetworkingV2ExtensionsLbaasMembersRequest{}
}

//response struct for the UpdateNetworkingV2ExtensionsLbaasMembers
type UpdateNetworkingV2ExtensionsLbaasMembersResponse struct{
    UpdateResult members.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsLbaasMembersResponse(updateResult members.UpdateResult,)*UpdateNetworkingV2ExtensionsLbaasMembersResponse {
    return &UpdateNetworkingV2ExtensionsLbaasMembersResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsLbaasMembers(req *UpdateNetworkingV2ExtensionsLbaasMembersRequest)(*UpdateNetworkingV2ExtensionsLbaasMembersResponse){
    return NewUpdateNetworkingV2ExtensionsLbaasMembersResponse(members.Update(oc.client,req.Id,req.Opts, ))

}
//request struct for the DeleteNetworkingV2ExtensionsLbaasMembers
type DeleteNetworkingV2ExtensionsLbaasMembersRequest struct{
    Id string
}

func NewDeleteNetworkingV2ExtensionsLbaasMembersRequest()*DeleteNetworkingV2ExtensionsLbaasMembersRequest{
    return &DeleteNetworkingV2ExtensionsLbaasMembersRequest{}
}

//response struct for the DeleteNetworkingV2ExtensionsLbaasMembers
type DeleteNetworkingV2ExtensionsLbaasMembersResponse struct{
    DeleteResult members.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsLbaasMembersResponse(deleteResult members.DeleteResult,)*DeleteNetworkingV2ExtensionsLbaasMembersResponse {
    return &DeleteNetworkingV2ExtensionsLbaasMembersResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsLbaasMembers(req *DeleteNetworkingV2ExtensionsLbaasMembersRequest)(*DeleteNetworkingV2ExtensionsLbaasMembersResponse){
    return NewDeleteNetworkingV2ExtensionsLbaasMembersResponse(members.Delete(oc.client,req.Id, ))

}