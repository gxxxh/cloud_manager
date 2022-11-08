package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/orchestration/v1/stackresources"
)

// call result's extract function
func ExtractFindOrchestrationV1StackresourcesResponse(response *FindOrchestrationV1StackresourcesResponse) (interface{}, error) {
	return response.FindResult.Body, response.FindResult.Err
}

// extract response info from pager for ListOrchestrationV1Stackresources
func ExtractListOrchestrationV1StackresourcesResponse(response *ListOrchestrationV1StackresourcesResponse) ([]stackresources.Resource, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return stackresources.ExtractResources(page)
}

// call result's extract function
func ExtractGetOrchestrationV1StackresourcesResponse(response *GetOrchestrationV1StackresourcesResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractMetadataOrchestrationV1StackresourcesResponse(response *MetadataOrchestrationV1StackresourcesResponse) (interface{}, error) {
	return response.MetadataResult.Body, response.MetadataResult.Err
}

// extract response info from pager for ListTypesOrchestrationV1Stackresources
func ExtractListTypesOrchestrationV1StackresourcesResponse(response *ListTypesOrchestrationV1StackresourcesResponse) (stackresources.ResourceTypes, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return stackresources.ExtractResourceTypes(page)
}

// call result's extract function
func ExtractSchemaOrchestrationV1StackresourcesResponse(response *SchemaOrchestrationV1StackresourcesResponse) (interface{}, error) {
	return response.SchemaResult.Body, response.SchemaResult.Err
}

// call result's extract function
func ExtractTemplateOrchestrationV1StackresourcesResponse(response *TemplateOrchestrationV1StackresourcesResponse) (interface{}, error) {
	return response.TemplateResult.Body, response.TemplateResult.Err
}

// call result's extract function
func ExtractMarkUnhealthyOrchestrationV1StackresourcesResponse(response *MarkUnhealthyOrchestrationV1StackresourcesResponse) (interface{}, error) {
	return response.MarkUnhealthyResult.Body, response.MarkUnhealthyResult.Err
}