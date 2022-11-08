package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/containerinfra/v1/clustertemplates"
)

// call result's extract function
func ExtractCreateContainerinfraV1ClustertemplatesResponse(response *CreateContainerinfraV1ClustertemplatesResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// call result's extract function
func ExtractDeleteContainerinfraV1ClustertemplatesResponse(response *DeleteContainerinfraV1ClustertemplatesResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}

// extract response info from pager for ListContainerinfraV1Clustertemplates
func ExtractListContainerinfraV1ClustertemplatesResponse(response *ListContainerinfraV1ClustertemplatesResponse) ([]clustertemplates.ClusterTemplate, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return clustertemplates.ExtractClusterTemplates(page)
}

// call result's extract function
func ExtractGetContainerinfraV1ClustertemplatesResponse(response *GetContainerinfraV1ClustertemplatesResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}