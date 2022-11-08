package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/common/extensions"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the GetCommonExtensions
type GetCommonExtensionsRequest struct {
	Alias string
}

func NewGetCommonExtensionsRequest() *GetCommonExtensionsRequest {
	return &GetCommonExtensionsRequest{}
}

// response struct for the GetCommonExtensions
type GetCommonExtensionsResponse struct {
	GetResult extensions.GetResult
}

func NewGetCommonExtensionsResponse(getResult extensions.GetResult) *GetCommonExtensionsResponse {
	return &GetCommonExtensionsResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetCommonExtensions(req *GetCommonExtensionsRequest) *GetCommonExtensionsResponse {
	return NewGetCommonExtensionsResponse(extensions.Get(oc.Client, req.Alias))

}

// request struct for the ListCommonExtensions
type ListCommonExtensionsRequest struct {
}

func NewListCommonExtensionsRequest() *ListCommonExtensionsRequest {
	return &ListCommonExtensionsRequest{}
}

// response struct for the ListCommonExtensions
type ListCommonExtensionsResponse struct {
	Pager pagination.Pager
}

func NewListCommonExtensionsResponse(pager pagination.Pager) *ListCommonExtensionsResponse {
	return &ListCommonExtensionsResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListCommonExtensions(req *ListCommonExtensionsRequest) *ListCommonExtensionsResponse {
	return NewListCommonExtensionsResponse(extensions.List(oc.Client))

}