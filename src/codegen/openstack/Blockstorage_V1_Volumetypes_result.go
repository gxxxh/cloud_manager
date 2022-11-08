package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumetypes"
)

// call result's extract function
func ExtractCreateBlockstorageV1VolumetypesResponse(response *CreateBlockstorageV1VolumetypesResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// call result's extract function
func ExtractDeleteBlockstorageV1VolumetypesResponse(response *DeleteBlockstorageV1VolumetypesResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}

// call result's extract function
func ExtractGetBlockstorageV1VolumetypesResponse(response *GetBlockstorageV1VolumetypesResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// extract response info from pager for ListBlockstorageV1Volumetypes
func ExtractListBlockstorageV1VolumetypesResponse(response *ListBlockstorageV1VolumetypesResponse) ([]volumetypes.VolumeType, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return volumetypes.ExtractVolumeTypes(page)
}