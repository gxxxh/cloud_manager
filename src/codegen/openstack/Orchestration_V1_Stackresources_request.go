package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/orchestration/v1/stackresources"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the FindOrchestrationV1Stackresources
type FindOrchestrationV1StackresourcesRequest struct {
	StackName string
}

func NewFindOrchestrationV1StackresourcesRequest() *FindOrchestrationV1StackresourcesRequest {
	return &FindOrchestrationV1StackresourcesRequest{}
}

// response struct for the FindOrchestrationV1Stackresources
type FindOrchestrationV1StackresourcesResponse struct {
	FindResult stackresources.FindResult
}

func NewFindOrchestrationV1StackresourcesResponse(findResult stackresources.FindResult) *FindOrchestrationV1StackresourcesResponse {
	return &FindOrchestrationV1StackresourcesResponse{
		FindResult: findResult,
	}
}

// action function
func (oc *OpenstackClient) FindOrchestrationV1Stackresources(req *FindOrchestrationV1StackresourcesRequest) *FindOrchestrationV1StackresourcesResponse {
	return NewFindOrchestrationV1StackresourcesResponse(stackresources.Find(oc.Client, req.StackName))

}

// request struct for the ListOrchestrationV1Stackresources
type ListOrchestrationV1StackresourcesRequest struct {
	StackName string
	StackID   string
	Opts      stackresources.ListOpts
}

func NewListOrchestrationV1StackresourcesRequest() *ListOrchestrationV1StackresourcesRequest {
	return &ListOrchestrationV1StackresourcesRequest{}
}

// response struct for the ListOrchestrationV1Stackresources
type ListOrchestrationV1StackresourcesResponse struct {
	Pager pagination.Pager
}

func NewListOrchestrationV1StackresourcesResponse(pager pagination.Pager) *ListOrchestrationV1StackresourcesResponse {
	return &ListOrchestrationV1StackresourcesResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListOrchestrationV1Stackresources(req *ListOrchestrationV1StackresourcesRequest) *ListOrchestrationV1StackresourcesResponse {
	return NewListOrchestrationV1StackresourcesResponse(stackresources.List(oc.Client, req.StackName, req.StackID, req.Opts))

}

// request struct for the GetOrchestrationV1Stackresources
type GetOrchestrationV1StackresourcesRequest struct {
	StackName    string
	StackID      string
	ResourceName string
}

func NewGetOrchestrationV1StackresourcesRequest() *GetOrchestrationV1StackresourcesRequest {
	return &GetOrchestrationV1StackresourcesRequest{}
}

// response struct for the GetOrchestrationV1Stackresources
type GetOrchestrationV1StackresourcesResponse struct {
	GetResult stackresources.GetResult
}

func NewGetOrchestrationV1StackresourcesResponse(getResult stackresources.GetResult) *GetOrchestrationV1StackresourcesResponse {
	return &GetOrchestrationV1StackresourcesResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetOrchestrationV1Stackresources(req *GetOrchestrationV1StackresourcesRequest) *GetOrchestrationV1StackresourcesResponse {
	return NewGetOrchestrationV1StackresourcesResponse(stackresources.Get(oc.Client, req.StackName, req.StackID, req.ResourceName))

}

// request struct for the MetadataOrchestrationV1Stackresources
type MetadataOrchestrationV1StackresourcesRequest struct {
	StackName    string
	StackID      string
	ResourceName string
}

func NewMetadataOrchestrationV1StackresourcesRequest() *MetadataOrchestrationV1StackresourcesRequest {
	return &MetadataOrchestrationV1StackresourcesRequest{}
}

// response struct for the MetadataOrchestrationV1Stackresources
type MetadataOrchestrationV1StackresourcesResponse struct {
	MetadataResult stackresources.MetadataResult
}

func NewMetadataOrchestrationV1StackresourcesResponse(metadataResult stackresources.MetadataResult) *MetadataOrchestrationV1StackresourcesResponse {
	return &MetadataOrchestrationV1StackresourcesResponse{
		MetadataResult: metadataResult,
	}
}

// action function
func (oc *OpenstackClient) MetadataOrchestrationV1Stackresources(req *MetadataOrchestrationV1StackresourcesRequest) *MetadataOrchestrationV1StackresourcesResponse {
	return NewMetadataOrchestrationV1StackresourcesResponse(stackresources.Metadata(oc.Client, req.StackName, req.StackID, req.ResourceName))

}

// request struct for the ListTypesOrchestrationV1Stackresources
type ListTypesOrchestrationV1StackresourcesRequest struct {
}

func NewListTypesOrchestrationV1StackresourcesRequest() *ListTypesOrchestrationV1StackresourcesRequest {
	return &ListTypesOrchestrationV1StackresourcesRequest{}
}

// response struct for the ListTypesOrchestrationV1Stackresources
type ListTypesOrchestrationV1StackresourcesResponse struct {
	Pager pagination.Pager
}

func NewListTypesOrchestrationV1StackresourcesResponse(pager pagination.Pager) *ListTypesOrchestrationV1StackresourcesResponse {
	return &ListTypesOrchestrationV1StackresourcesResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListTypesOrchestrationV1Stackresources(req *ListTypesOrchestrationV1StackresourcesRequest) *ListTypesOrchestrationV1StackresourcesResponse {
	return NewListTypesOrchestrationV1StackresourcesResponse(stackresources.ListTypes(oc.Client))

}

// request struct for the SchemaOrchestrationV1Stackresources
type SchemaOrchestrationV1StackresourcesRequest struct {
	ResourceType string
}

func NewSchemaOrchestrationV1StackresourcesRequest() *SchemaOrchestrationV1StackresourcesRequest {
	return &SchemaOrchestrationV1StackresourcesRequest{}
}

// response struct for the SchemaOrchestrationV1Stackresources
type SchemaOrchestrationV1StackresourcesResponse struct {
	SchemaResult stackresources.SchemaResult
}

func NewSchemaOrchestrationV1StackresourcesResponse(schemaResult stackresources.SchemaResult) *SchemaOrchestrationV1StackresourcesResponse {
	return &SchemaOrchestrationV1StackresourcesResponse{
		SchemaResult: schemaResult,
	}
}

// action function
func (oc *OpenstackClient) SchemaOrchestrationV1Stackresources(req *SchemaOrchestrationV1StackresourcesRequest) *SchemaOrchestrationV1StackresourcesResponse {
	return NewSchemaOrchestrationV1StackresourcesResponse(stackresources.Schema(oc.Client, req.ResourceType))

}

// request struct for the TemplateOrchestrationV1Stackresources
type TemplateOrchestrationV1StackresourcesRequest struct {
	ResourceType string
}

func NewTemplateOrchestrationV1StackresourcesRequest() *TemplateOrchestrationV1StackresourcesRequest {
	return &TemplateOrchestrationV1StackresourcesRequest{}
}

// response struct for the TemplateOrchestrationV1Stackresources
type TemplateOrchestrationV1StackresourcesResponse struct {
	TemplateResult stackresources.TemplateResult
}

func NewTemplateOrchestrationV1StackresourcesResponse(templateResult stackresources.TemplateResult) *TemplateOrchestrationV1StackresourcesResponse {
	return &TemplateOrchestrationV1StackresourcesResponse{
		TemplateResult: templateResult,
	}
}

// action function
func (oc *OpenstackClient) TemplateOrchestrationV1Stackresources(req *TemplateOrchestrationV1StackresourcesRequest) *TemplateOrchestrationV1StackresourcesResponse {
	return NewTemplateOrchestrationV1StackresourcesResponse(stackresources.Template(oc.Client, req.ResourceType))

}

// request struct for the MarkUnhealthyOrchestrationV1Stackresources
type MarkUnhealthyOrchestrationV1StackresourcesRequest struct {
	StackName    string
	StackID      string
	ResourceName string
	Opts         stackresources.MarkUnhealthyOpts
}

func NewMarkUnhealthyOrchestrationV1StackresourcesRequest() *MarkUnhealthyOrchestrationV1StackresourcesRequest {
	return &MarkUnhealthyOrchestrationV1StackresourcesRequest{}
}

// response struct for the MarkUnhealthyOrchestrationV1Stackresources
type MarkUnhealthyOrchestrationV1StackresourcesResponse struct {
	MarkUnhealthyResult stackresources.MarkUnhealthyResult
}

func NewMarkUnhealthyOrchestrationV1StackresourcesResponse(markUnhealthyResult stackresources.MarkUnhealthyResult) *MarkUnhealthyOrchestrationV1StackresourcesResponse {
	return &MarkUnhealthyOrchestrationV1StackresourcesResponse{
		MarkUnhealthyResult: markUnhealthyResult,
	}
}

// action function
func (oc *OpenstackClient) MarkUnhealthyOrchestrationV1Stackresources(req *MarkUnhealthyOrchestrationV1StackresourcesRequest) *MarkUnhealthyOrchestrationV1StackresourcesResponse {
	return NewMarkUnhealthyOrchestrationV1StackresourcesResponse(stackresources.MarkUnhealthy(oc.Client, req.StackName, req.StackID, req.ResourceName, req.Opts))

}