package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas/vips"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the ListNetworkingV2ExtensionsLbaasVips
type ListNetworkingV2ExtensionsLbaasVipsRequest struct {
	Opts vips.ListOpts
}

func NewListNetworkingV2ExtensionsLbaasVipsRequest() *ListNetworkingV2ExtensionsLbaasVipsRequest {
	return &ListNetworkingV2ExtensionsLbaasVipsRequest{}
}

// response struct for the ListNetworkingV2ExtensionsLbaasVips
type ListNetworkingV2ExtensionsLbaasVipsResponse struct {
	Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsLbaasVipsResponse(pager pagination.Pager) *ListNetworkingV2ExtensionsLbaasVipsResponse {
	return &ListNetworkingV2ExtensionsLbaasVipsResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsLbaasVips(req *ListNetworkingV2ExtensionsLbaasVipsRequest) *ListNetworkingV2ExtensionsLbaasVipsResponse {
	return NewListNetworkingV2ExtensionsLbaasVipsResponse(vips.List(oc.Client, req.Opts))

}

// request struct for the CreateNetworkingV2ExtensionsLbaasVips
type CreateNetworkingV2ExtensionsLbaasVipsRequest struct {
	Opts vips.CreateOpts
}

func NewCreateNetworkingV2ExtensionsLbaasVipsRequest() *CreateNetworkingV2ExtensionsLbaasVipsRequest {
	return &CreateNetworkingV2ExtensionsLbaasVipsRequest{}
}

// response struct for the CreateNetworkingV2ExtensionsLbaasVips
type CreateNetworkingV2ExtensionsLbaasVipsResponse struct {
	CreateResult vips.CreateResult
}

func NewCreateNetworkingV2ExtensionsLbaasVipsResponse(createResult vips.CreateResult) *CreateNetworkingV2ExtensionsLbaasVipsResponse {
	return &CreateNetworkingV2ExtensionsLbaasVipsResponse{
		CreateResult: createResult,
	}
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsLbaasVips(req *CreateNetworkingV2ExtensionsLbaasVipsRequest) *CreateNetworkingV2ExtensionsLbaasVipsResponse {
	return NewCreateNetworkingV2ExtensionsLbaasVipsResponse(vips.Create(oc.Client, req.Opts))

}

// request struct for the GetNetworkingV2ExtensionsLbaasVips
type GetNetworkingV2ExtensionsLbaasVipsRequest struct {
	Id string
}

func NewGetNetworkingV2ExtensionsLbaasVipsRequest() *GetNetworkingV2ExtensionsLbaasVipsRequest {
	return &GetNetworkingV2ExtensionsLbaasVipsRequest{}
}

// response struct for the GetNetworkingV2ExtensionsLbaasVips
type GetNetworkingV2ExtensionsLbaasVipsResponse struct {
	GetResult vips.GetResult
}

func NewGetNetworkingV2ExtensionsLbaasVipsResponse(getResult vips.GetResult) *GetNetworkingV2ExtensionsLbaasVipsResponse {
	return &GetNetworkingV2ExtensionsLbaasVipsResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsLbaasVips(req *GetNetworkingV2ExtensionsLbaasVipsRequest) *GetNetworkingV2ExtensionsLbaasVipsResponse {
	return NewGetNetworkingV2ExtensionsLbaasVipsResponse(vips.Get(oc.Client, req.Id))

}

// request struct for the UpdateNetworkingV2ExtensionsLbaasVips
type UpdateNetworkingV2ExtensionsLbaasVipsRequest struct {
	Id   string
	Opts vips.UpdateOpts
}

func NewUpdateNetworkingV2ExtensionsLbaasVipsRequest() *UpdateNetworkingV2ExtensionsLbaasVipsRequest {
	return &UpdateNetworkingV2ExtensionsLbaasVipsRequest{}
}

// response struct for the UpdateNetworkingV2ExtensionsLbaasVips
type UpdateNetworkingV2ExtensionsLbaasVipsResponse struct {
	UpdateResult vips.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsLbaasVipsResponse(updateResult vips.UpdateResult) *UpdateNetworkingV2ExtensionsLbaasVipsResponse {
	return &UpdateNetworkingV2ExtensionsLbaasVipsResponse{
		UpdateResult: updateResult,
	}
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsLbaasVips(req *UpdateNetworkingV2ExtensionsLbaasVipsRequest) *UpdateNetworkingV2ExtensionsLbaasVipsResponse {
	return NewUpdateNetworkingV2ExtensionsLbaasVipsResponse(vips.Update(oc.Client, req.Id, req.Opts))

}

// request struct for the DeleteNetworkingV2ExtensionsLbaasVips
type DeleteNetworkingV2ExtensionsLbaasVipsRequest struct {
	Id string
}

func NewDeleteNetworkingV2ExtensionsLbaasVipsRequest() *DeleteNetworkingV2ExtensionsLbaasVipsRequest {
	return &DeleteNetworkingV2ExtensionsLbaasVipsRequest{}
}

// response struct for the DeleteNetworkingV2ExtensionsLbaasVips
type DeleteNetworkingV2ExtensionsLbaasVipsResponse struct {
	DeleteResult vips.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsLbaasVipsResponse(deleteResult vips.DeleteResult) *DeleteNetworkingV2ExtensionsLbaasVipsResponse {
	return &DeleteNetworkingV2ExtensionsLbaasVipsResponse{
		DeleteResult: deleteResult,
	}
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsLbaasVips(req *DeleteNetworkingV2ExtensionsLbaasVipsRequest) *DeleteNetworkingV2ExtensionsLbaasVipsResponse {
	return NewDeleteNetworkingV2ExtensionsLbaasVipsResponse(vips.Delete(oc.Client, req.Id))

}