package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/vpnaas/ikepolicies"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the CreateNetworkingV2ExtensionsVpnaasIkepolicies
type CreateNetworkingV2ExtensionsVpnaasIkepoliciesRequest struct {
	Opts ikepolicies.CreateOpts
}

func NewCreateNetworkingV2ExtensionsVpnaasIkepoliciesRequest() *CreateNetworkingV2ExtensionsVpnaasIkepoliciesRequest {
	return &CreateNetworkingV2ExtensionsVpnaasIkepoliciesRequest{}
}

// response struct for the CreateNetworkingV2ExtensionsVpnaasIkepolicies
type CreateNetworkingV2ExtensionsVpnaasIkepoliciesResponse struct {
	CreateResult ikepolicies.CreateResult
}

func NewCreateNetworkingV2ExtensionsVpnaasIkepoliciesResponse(createResult ikepolicies.CreateResult) *CreateNetworkingV2ExtensionsVpnaasIkepoliciesResponse {
	return &CreateNetworkingV2ExtensionsVpnaasIkepoliciesResponse{
		CreateResult: createResult,
	}
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsVpnaasIkepolicies(req *CreateNetworkingV2ExtensionsVpnaasIkepoliciesRequest) *CreateNetworkingV2ExtensionsVpnaasIkepoliciesResponse {
	return NewCreateNetworkingV2ExtensionsVpnaasIkepoliciesResponse(ikepolicies.Create(oc.Client, req.Opts))

}

// request struct for the GetNetworkingV2ExtensionsVpnaasIkepolicies
type GetNetworkingV2ExtensionsVpnaasIkepoliciesRequest struct {
	Id string
}

func NewGetNetworkingV2ExtensionsVpnaasIkepoliciesRequest() *GetNetworkingV2ExtensionsVpnaasIkepoliciesRequest {
	return &GetNetworkingV2ExtensionsVpnaasIkepoliciesRequest{}
}

// response struct for the GetNetworkingV2ExtensionsVpnaasIkepolicies
type GetNetworkingV2ExtensionsVpnaasIkepoliciesResponse struct {
	GetResult ikepolicies.GetResult
}

func NewGetNetworkingV2ExtensionsVpnaasIkepoliciesResponse(getResult ikepolicies.GetResult) *GetNetworkingV2ExtensionsVpnaasIkepoliciesResponse {
	return &GetNetworkingV2ExtensionsVpnaasIkepoliciesResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsVpnaasIkepolicies(req *GetNetworkingV2ExtensionsVpnaasIkepoliciesRequest) *GetNetworkingV2ExtensionsVpnaasIkepoliciesResponse {
	return NewGetNetworkingV2ExtensionsVpnaasIkepoliciesResponse(ikepolicies.Get(oc.Client, req.Id))

}

// request struct for the DeleteNetworkingV2ExtensionsVpnaasIkepolicies
type DeleteNetworkingV2ExtensionsVpnaasIkepoliciesRequest struct {
	Id string
}

func NewDeleteNetworkingV2ExtensionsVpnaasIkepoliciesRequest() *DeleteNetworkingV2ExtensionsVpnaasIkepoliciesRequest {
	return &DeleteNetworkingV2ExtensionsVpnaasIkepoliciesRequest{}
}

// response struct for the DeleteNetworkingV2ExtensionsVpnaasIkepolicies
type DeleteNetworkingV2ExtensionsVpnaasIkepoliciesResponse struct {
	DeleteResult ikepolicies.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsVpnaasIkepoliciesResponse(deleteResult ikepolicies.DeleteResult) *DeleteNetworkingV2ExtensionsVpnaasIkepoliciesResponse {
	return &DeleteNetworkingV2ExtensionsVpnaasIkepoliciesResponse{
		DeleteResult: deleteResult,
	}
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsVpnaasIkepolicies(req *DeleteNetworkingV2ExtensionsVpnaasIkepoliciesRequest) *DeleteNetworkingV2ExtensionsVpnaasIkepoliciesResponse {
	return NewDeleteNetworkingV2ExtensionsVpnaasIkepoliciesResponse(ikepolicies.Delete(oc.Client, req.Id))

}

// request struct for the ListNetworkingV2ExtensionsVpnaasIkepolicies
type ListNetworkingV2ExtensionsVpnaasIkepoliciesRequest struct {
	Opts ikepolicies.ListOpts
}

func NewListNetworkingV2ExtensionsVpnaasIkepoliciesRequest() *ListNetworkingV2ExtensionsVpnaasIkepoliciesRequest {
	return &ListNetworkingV2ExtensionsVpnaasIkepoliciesRequest{}
}

// response struct for the ListNetworkingV2ExtensionsVpnaasIkepolicies
type ListNetworkingV2ExtensionsVpnaasIkepoliciesResponse struct {
	Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsVpnaasIkepoliciesResponse(pager pagination.Pager) *ListNetworkingV2ExtensionsVpnaasIkepoliciesResponse {
	return &ListNetworkingV2ExtensionsVpnaasIkepoliciesResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsVpnaasIkepolicies(req *ListNetworkingV2ExtensionsVpnaasIkepoliciesRequest) *ListNetworkingV2ExtensionsVpnaasIkepoliciesResponse {
	return NewListNetworkingV2ExtensionsVpnaasIkepoliciesResponse(ikepolicies.List(oc.Client, req.Opts))

}

// request struct for the UpdateNetworkingV2ExtensionsVpnaasIkepolicies
type UpdateNetworkingV2ExtensionsVpnaasIkepoliciesRequest struct {
	Id   string
	Opts ikepolicies.UpdateOpts
}

func NewUpdateNetworkingV2ExtensionsVpnaasIkepoliciesRequest() *UpdateNetworkingV2ExtensionsVpnaasIkepoliciesRequest {
	return &UpdateNetworkingV2ExtensionsVpnaasIkepoliciesRequest{}
}

// response struct for the UpdateNetworkingV2ExtensionsVpnaasIkepolicies
type UpdateNetworkingV2ExtensionsVpnaasIkepoliciesResponse struct {
	UpdateResult ikepolicies.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsVpnaasIkepoliciesResponse(updateResult ikepolicies.UpdateResult) *UpdateNetworkingV2ExtensionsVpnaasIkepoliciesResponse {
	return &UpdateNetworkingV2ExtensionsVpnaasIkepoliciesResponse{
		UpdateResult: updateResult,
	}
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsVpnaasIkepolicies(req *UpdateNetworkingV2ExtensionsVpnaasIkepoliciesRequest) *UpdateNetworkingV2ExtensionsVpnaasIkepoliciesResponse {
	return NewUpdateNetworkingV2ExtensionsVpnaasIkepoliciesResponse(ikepolicies.Update(oc.Client, req.Id, req.Opts))

}