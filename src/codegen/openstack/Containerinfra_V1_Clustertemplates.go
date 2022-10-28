package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/containerinfra/v1/clustertemplates"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateContainerinfraV1Clustertemplates
type CreateContainerinfraV1ClustertemplatesRequest struct{
    Opts clustertemplates.CreateOptsBuilder
}

func NewCreateContainerinfraV1ClustertemplatesRequest()*CreateContainerinfraV1ClustertemplatesRequest{
    return &CreateContainerinfraV1ClustertemplatesRequest{}
}

//response struct for the CreateContainerinfraV1Clustertemplates
type CreateContainerinfraV1ClustertemplatesResponse struct{
    CreateResult clustertemplates.CreateResult
}

func NewCreateContainerinfraV1ClustertemplatesResponse(createResult clustertemplates.CreateResult,)*CreateContainerinfraV1ClustertemplatesResponse {
    return &CreateContainerinfraV1ClustertemplatesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateContainerinfraV1Clustertemplates(request *CreateContainerinfraV1ClustertemplatesRequest)(*CreateContainerinfraV1ClustertemplatesResponse){
    return NewCreateContainerinfraV1ClustertemplatesResponse(clustertemplates.Create(oc.client,request.Opts, ))

}
//request struct for the DeleteContainerinfraV1Clustertemplates
type DeleteContainerinfraV1ClustertemplatesRequest struct{
    Id string
}

func NewDeleteContainerinfraV1ClustertemplatesRequest()*DeleteContainerinfraV1ClustertemplatesRequest{
    return &DeleteContainerinfraV1ClustertemplatesRequest{}
}

//response struct for the DeleteContainerinfraV1Clustertemplates
type DeleteContainerinfraV1ClustertemplatesResponse struct{
    DeleteResult clustertemplates.DeleteResult
}

func NewDeleteContainerinfraV1ClustertemplatesResponse(deleteResult clustertemplates.DeleteResult,)*DeleteContainerinfraV1ClustertemplatesResponse {
    return &DeleteContainerinfraV1ClustertemplatesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteContainerinfraV1Clustertemplates(request *DeleteContainerinfraV1ClustertemplatesRequest)(*DeleteContainerinfraV1ClustertemplatesResponse){
    return NewDeleteContainerinfraV1ClustertemplatesResponse(clustertemplates.Delete(oc.client,request.Id, ))

}
//request struct for the ListContainerinfraV1Clustertemplates
type ListContainerinfraV1ClustertemplatesRequest struct{
    Opts clustertemplates.ListOptsBuilder
}

func NewListContainerinfraV1ClustertemplatesRequest()*ListContainerinfraV1ClustertemplatesRequest{
    return &ListContainerinfraV1ClustertemplatesRequest{}
}

//response struct for the ListContainerinfraV1Clustertemplates
type ListContainerinfraV1ClustertemplatesResponse struct{
    Pager pagination.Pager
}

func NewListContainerinfraV1ClustertemplatesResponse(pager pagination.Pager,)*ListContainerinfraV1ClustertemplatesResponse {
    return &ListContainerinfraV1ClustertemplatesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListContainerinfraV1Clustertemplates(request *ListContainerinfraV1ClustertemplatesRequest)(*ListContainerinfraV1ClustertemplatesResponse){
    return NewListContainerinfraV1ClustertemplatesResponse(clustertemplates.List(oc.client,request.Opts, ))

}
//request struct for the GetContainerinfraV1Clustertemplates
type GetContainerinfraV1ClustertemplatesRequest struct{
    Id string
}

func NewGetContainerinfraV1ClustertemplatesRequest()*GetContainerinfraV1ClustertemplatesRequest{
    return &GetContainerinfraV1ClustertemplatesRequest{}
}

//response struct for the GetContainerinfraV1Clustertemplates
type GetContainerinfraV1ClustertemplatesResponse struct{
    GetResult clustertemplates.GetResult
}

func NewGetContainerinfraV1ClustertemplatesResponse(getResult clustertemplates.GetResult,)*GetContainerinfraV1ClustertemplatesResponse {
    return &GetContainerinfraV1ClustertemplatesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetContainerinfraV1Clustertemplates(request *GetContainerinfraV1ClustertemplatesRequest)(*GetContainerinfraV1ClustertemplatesResponse){
    return NewGetContainerinfraV1ClustertemplatesResponse(clustertemplates.Get(oc.client,request.Id, ))

}
//request struct for the UpdateContainerinfraV1Clustertemplates
type UpdateContainerinfraV1ClustertemplatesRequest struct{
    Id string
    Opts []clustertemplates.UpdateOptsBuilder
}

func NewUpdateContainerinfraV1ClustertemplatesRequest()*UpdateContainerinfraV1ClustertemplatesRequest{
    return &UpdateContainerinfraV1ClustertemplatesRequest{}
}

//response struct for the UpdateContainerinfraV1Clustertemplates
type UpdateContainerinfraV1ClustertemplatesResponse struct{
    UpdateResult clustertemplates.UpdateResult
}

func NewUpdateContainerinfraV1ClustertemplatesResponse(updateResult clustertemplates.UpdateResult,)*UpdateContainerinfraV1ClustertemplatesResponse {
    return &UpdateContainerinfraV1ClustertemplatesResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateContainerinfraV1Clustertemplates(request *UpdateContainerinfraV1ClustertemplatesRequest)(*UpdateContainerinfraV1ClustertemplatesResponse){
    return NewUpdateContainerinfraV1ClustertemplatesResponse(clustertemplates.Update(oc.client,request.Id,request.Opts, ))

}