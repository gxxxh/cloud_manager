package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/aggregates"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the ListComputeV2ExtensionsAggregates
type ListComputeV2ExtensionsAggregatesRequest struct {
}

func NewListComputeV2ExtensionsAggregatesRequest() *ListComputeV2ExtensionsAggregatesRequest {
	return &ListComputeV2ExtensionsAggregatesRequest{}
}

// response struct for the ListComputeV2ExtensionsAggregates
type ListComputeV2ExtensionsAggregatesResponse struct {
	Pager pagination.Pager
}

func NewListComputeV2ExtensionsAggregatesResponse(pager pagination.Pager) *ListComputeV2ExtensionsAggregatesResponse {
	return &ListComputeV2ExtensionsAggregatesResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListComputeV2ExtensionsAggregates(req *ListComputeV2ExtensionsAggregatesRequest) *ListComputeV2ExtensionsAggregatesResponse {
	return NewListComputeV2ExtensionsAggregatesResponse(aggregates.List(oc.Client))

}

// request struct for the CreateComputeV2ExtensionsAggregates
type CreateComputeV2ExtensionsAggregatesRequest struct {
	Opts aggregates.CreateOpts
}

func NewCreateComputeV2ExtensionsAggregatesRequest() *CreateComputeV2ExtensionsAggregatesRequest {
	return &CreateComputeV2ExtensionsAggregatesRequest{}
}

// response struct for the CreateComputeV2ExtensionsAggregates
type CreateComputeV2ExtensionsAggregatesResponse struct {
	CreateResult aggregates.CreateResult
}

func NewCreateComputeV2ExtensionsAggregatesResponse(createResult aggregates.CreateResult) *CreateComputeV2ExtensionsAggregatesResponse {
	return &CreateComputeV2ExtensionsAggregatesResponse{
		CreateResult: createResult,
	}
}

// action function
func (oc *OpenstackClient) CreateComputeV2ExtensionsAggregates(req *CreateComputeV2ExtensionsAggregatesRequest) *CreateComputeV2ExtensionsAggregatesResponse {
	return NewCreateComputeV2ExtensionsAggregatesResponse(aggregates.Create(oc.Client, req.Opts))

}

// request struct for the DeleteComputeV2ExtensionsAggregates
type DeleteComputeV2ExtensionsAggregatesRequest struct {
	AggregateID int
}

func NewDeleteComputeV2ExtensionsAggregatesRequest() *DeleteComputeV2ExtensionsAggregatesRequest {
	return &DeleteComputeV2ExtensionsAggregatesRequest{}
}

// response struct for the DeleteComputeV2ExtensionsAggregates
type DeleteComputeV2ExtensionsAggregatesResponse struct {
	DeleteResult aggregates.DeleteResult
}

func NewDeleteComputeV2ExtensionsAggregatesResponse(deleteResult aggregates.DeleteResult) *DeleteComputeV2ExtensionsAggregatesResponse {
	return &DeleteComputeV2ExtensionsAggregatesResponse{
		DeleteResult: deleteResult,
	}
}

// action function
func (oc *OpenstackClient) DeleteComputeV2ExtensionsAggregates(req *DeleteComputeV2ExtensionsAggregatesRequest) *DeleteComputeV2ExtensionsAggregatesResponse {
	return NewDeleteComputeV2ExtensionsAggregatesResponse(aggregates.Delete(oc.Client, req.AggregateID))

}

// request struct for the GetComputeV2ExtensionsAggregates
type GetComputeV2ExtensionsAggregatesRequest struct {
	AggregateID int
}

func NewGetComputeV2ExtensionsAggregatesRequest() *GetComputeV2ExtensionsAggregatesRequest {
	return &GetComputeV2ExtensionsAggregatesRequest{}
}

// response struct for the GetComputeV2ExtensionsAggregates
type GetComputeV2ExtensionsAggregatesResponse struct {
	GetResult aggregates.GetResult
}

func NewGetComputeV2ExtensionsAggregatesResponse(getResult aggregates.GetResult) *GetComputeV2ExtensionsAggregatesResponse {
	return &GetComputeV2ExtensionsAggregatesResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetComputeV2ExtensionsAggregates(req *GetComputeV2ExtensionsAggregatesRequest) *GetComputeV2ExtensionsAggregatesResponse {
	return NewGetComputeV2ExtensionsAggregatesResponse(aggregates.Get(oc.Client, req.AggregateID))

}

// request struct for the UpdateComputeV2ExtensionsAggregates
type UpdateComputeV2ExtensionsAggregatesRequest struct {
	AggregateID int
	Opts        aggregates.UpdateOpts
}

func NewUpdateComputeV2ExtensionsAggregatesRequest() *UpdateComputeV2ExtensionsAggregatesRequest {
	return &UpdateComputeV2ExtensionsAggregatesRequest{}
}

// response struct for the UpdateComputeV2ExtensionsAggregates
type UpdateComputeV2ExtensionsAggregatesResponse struct {
	UpdateResult aggregates.UpdateResult
}

func NewUpdateComputeV2ExtensionsAggregatesResponse(updateResult aggregates.UpdateResult) *UpdateComputeV2ExtensionsAggregatesResponse {
	return &UpdateComputeV2ExtensionsAggregatesResponse{
		UpdateResult: updateResult,
	}
}

// action function
func (oc *OpenstackClient) UpdateComputeV2ExtensionsAggregates(req *UpdateComputeV2ExtensionsAggregatesRequest) *UpdateComputeV2ExtensionsAggregatesResponse {
	return NewUpdateComputeV2ExtensionsAggregatesResponse(aggregates.Update(oc.Client, req.AggregateID, req.Opts))

}

// request struct for the AddHostComputeV2ExtensionsAggregates
type AddHostComputeV2ExtensionsAggregatesRequest struct {
	AggregateID int
	Opts        aggregates.AddHostOpts
}

func NewAddHostComputeV2ExtensionsAggregatesRequest() *AddHostComputeV2ExtensionsAggregatesRequest {
	return &AddHostComputeV2ExtensionsAggregatesRequest{}
}

// response struct for the AddHostComputeV2ExtensionsAggregates
type AddHostComputeV2ExtensionsAggregatesResponse struct {
	ActionResult aggregates.ActionResult
}

func NewAddHostComputeV2ExtensionsAggregatesResponse(actionResult aggregates.ActionResult) *AddHostComputeV2ExtensionsAggregatesResponse {
	return &AddHostComputeV2ExtensionsAggregatesResponse{
		ActionResult: actionResult,
	}
}

// action function
func (oc *OpenstackClient) AddHostComputeV2ExtensionsAggregates(req *AddHostComputeV2ExtensionsAggregatesRequest) *AddHostComputeV2ExtensionsAggregatesResponse {
	return NewAddHostComputeV2ExtensionsAggregatesResponse(aggregates.AddHost(oc.Client, req.AggregateID, req.Opts))

}

// request struct for the RemoveHostComputeV2ExtensionsAggregates
type RemoveHostComputeV2ExtensionsAggregatesRequest struct {
	AggregateID int
	Opts        aggregates.RemoveHostOpts
}

func NewRemoveHostComputeV2ExtensionsAggregatesRequest() *RemoveHostComputeV2ExtensionsAggregatesRequest {
	return &RemoveHostComputeV2ExtensionsAggregatesRequest{}
}

// response struct for the RemoveHostComputeV2ExtensionsAggregates
type RemoveHostComputeV2ExtensionsAggregatesResponse struct {
	ActionResult aggregates.ActionResult
}

func NewRemoveHostComputeV2ExtensionsAggregatesResponse(actionResult aggregates.ActionResult) *RemoveHostComputeV2ExtensionsAggregatesResponse {
	return &RemoveHostComputeV2ExtensionsAggregatesResponse{
		ActionResult: actionResult,
	}
}

// action function
func (oc *OpenstackClient) RemoveHostComputeV2ExtensionsAggregates(req *RemoveHostComputeV2ExtensionsAggregatesRequest) *RemoveHostComputeV2ExtensionsAggregatesResponse {
	return NewRemoveHostComputeV2ExtensionsAggregatesResponse(aggregates.RemoveHost(oc.Client, req.AggregateID, req.Opts))

}

// request struct for the SetMetadataComputeV2ExtensionsAggregates
type SetMetadataComputeV2ExtensionsAggregatesRequest struct {
	AggregateID int
	Opts        aggregates.SetMetadataOpts
}

func NewSetMetadataComputeV2ExtensionsAggregatesRequest() *SetMetadataComputeV2ExtensionsAggregatesRequest {
	return &SetMetadataComputeV2ExtensionsAggregatesRequest{}
}

// response struct for the SetMetadataComputeV2ExtensionsAggregates
type SetMetadataComputeV2ExtensionsAggregatesResponse struct {
	ActionResult aggregates.ActionResult
}

func NewSetMetadataComputeV2ExtensionsAggregatesResponse(actionResult aggregates.ActionResult) *SetMetadataComputeV2ExtensionsAggregatesResponse {
	return &SetMetadataComputeV2ExtensionsAggregatesResponse{
		ActionResult: actionResult,
	}
}

// action function
func (oc *OpenstackClient) SetMetadataComputeV2ExtensionsAggregates(req *SetMetadataComputeV2ExtensionsAggregatesRequest) *SetMetadataComputeV2ExtensionsAggregatesResponse {
	return NewSetMetadataComputeV2ExtensionsAggregatesResponse(aggregates.SetMetadata(oc.Client, req.AggregateID, req.Opts))

}