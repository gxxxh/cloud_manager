package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/clustering/v1/events"
)

// extract response info from pager for ListClusteringV1Events
func ExtractListClusteringV1EventsResponse(response *ListClusteringV1EventsResponse) ([]events.Event, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return events.ExtractEvents(page)
}

// call result's extract function
func ExtractGetClusteringV1EventsResponse(response *GetClusteringV1EventsResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}