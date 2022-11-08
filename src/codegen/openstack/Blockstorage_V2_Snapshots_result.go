package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/snapshots"
)

// call result's extract function
func ExtractCreateBlockstorageV2SnapshotsResponse(response *CreateBlockstorageV2SnapshotsResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}

// call result's extract function
func ExtractDeleteBlockstorageV2SnapshotsResponse(response *DeleteBlockstorageV2SnapshotsResponse) (interface{}, error) {
	return response.DeleteResult.Body, response.DeleteResult.Err
}

// call result's extract function
func ExtractGetBlockstorageV2SnapshotsResponse(response *GetBlockstorageV2SnapshotsResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// extract response info from pager for ListBlockstorageV2Snapshots
func ExtractListBlockstorageV2SnapshotsResponse(response *ListBlockstorageV2SnapshotsResponse) ([]snapshots.Snapshot, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return snapshots.ExtractSnapshots(page)
}

// call result's extract function
func ExtractUpdateMetadataBlockstorageV2SnapshotsResponse(response *UpdateMetadataBlockstorageV2SnapshotsResponse) (interface{}, error) {
	return response.UpdateMetadataResult.Body, response.UpdateMetadataResult.Err
}