package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/regions"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListIdentityV3Regions
type ListIdentityV3RegionsRequest struct{
    Opts regions.ListOptsBuilder
}

func NewListIdentityV3RegionsRequest()*ListIdentityV3RegionsRequest{
    return &ListIdentityV3RegionsRequest{}
}

//response struct for the ListIdentityV3Regions
type ListIdentityV3RegionsResponse struct{
    Pager pagination.Pager
}

func NewListIdentityV3RegionsResponse(pager pagination.Pager,)*ListIdentityV3RegionsResponse {
    return &ListIdentityV3RegionsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListIdentityV3Regions(req *ListIdentityV3RegionsRequest)(*ListIdentityV3RegionsResponse){
    return NewListIdentityV3RegionsResponse(regions.List(oc.client,req.Opts, ))

}
//request struct for the GetIdentityV3Regions
type GetIdentityV3RegionsRequest struct{
    Id string
}

func NewGetIdentityV3RegionsRequest()*GetIdentityV3RegionsRequest{
    return &GetIdentityV3RegionsRequest{}
}

//response struct for the GetIdentityV3Regions
type GetIdentityV3RegionsResponse struct{
    GetResult regions.GetResult
}

func NewGetIdentityV3RegionsResponse(getResult regions.GetResult,)*GetIdentityV3RegionsResponse {
    return &GetIdentityV3RegionsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetIdentityV3Regions(req *GetIdentityV3RegionsRequest)(*GetIdentityV3RegionsResponse){
    return NewGetIdentityV3RegionsResponse(regions.Get(oc.client,req.Id, ))

}
//request struct for the CreateIdentityV3Regions
type CreateIdentityV3RegionsRequest struct{
    Opts regions.CreateOptsBuilder
}

func NewCreateIdentityV3RegionsRequest()*CreateIdentityV3RegionsRequest{
    return &CreateIdentityV3RegionsRequest{}
}

//response struct for the CreateIdentityV3Regions
type CreateIdentityV3RegionsResponse struct{
    CreateResult regions.CreateResult
}

func NewCreateIdentityV3RegionsResponse(createResult regions.CreateResult,)*CreateIdentityV3RegionsResponse {
    return &CreateIdentityV3RegionsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateIdentityV3Regions(req *CreateIdentityV3RegionsRequest)(*CreateIdentityV3RegionsResponse){
    return NewCreateIdentityV3RegionsResponse(regions.Create(oc.client,req.Opts, ))

}
//request struct for the UpdateIdentityV3Regions
type UpdateIdentityV3RegionsRequest struct{
    RegionID string
    Opts regions.UpdateOptsBuilder
}

func NewUpdateIdentityV3RegionsRequest()*UpdateIdentityV3RegionsRequest{
    return &UpdateIdentityV3RegionsRequest{}
}

//response struct for the UpdateIdentityV3Regions
type UpdateIdentityV3RegionsResponse struct{
    UpdateResult regions.UpdateResult
}

func NewUpdateIdentityV3RegionsResponse(updateResult regions.UpdateResult,)*UpdateIdentityV3RegionsResponse {
    return &UpdateIdentityV3RegionsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateIdentityV3Regions(req *UpdateIdentityV3RegionsRequest)(*UpdateIdentityV3RegionsResponse){
    return NewUpdateIdentityV3RegionsResponse(regions.Update(oc.client,req.RegionID,req.Opts, ))

}
//request struct for the DeleteIdentityV3Regions
type DeleteIdentityV3RegionsRequest struct{
    RegionID string
}

func NewDeleteIdentityV3RegionsRequest()*DeleteIdentityV3RegionsRequest{
    return &DeleteIdentityV3RegionsRequest{}
}

//response struct for the DeleteIdentityV3Regions
type DeleteIdentityV3RegionsResponse struct{
    DeleteResult regions.DeleteResult
}

func NewDeleteIdentityV3RegionsResponse(deleteResult regions.DeleteResult,)*DeleteIdentityV3RegionsResponse {
    return &DeleteIdentityV3RegionsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteIdentityV3Regions(req *DeleteIdentityV3RegionsRequest)(*DeleteIdentityV3RegionsResponse){
    return NewDeleteIdentityV3RegionsResponse(regions.Delete(oc.client,req.RegionID, ))

}