package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/availabilityzones"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListComputeV2ExtensionsAvailabilityzones
type ListComputeV2ExtensionsAvailabilityzonesRequest struct{
}

func NewListComputeV2ExtensionsAvailabilityzonesRequest()*ListComputeV2ExtensionsAvailabilityzonesRequest{
    return &ListComputeV2ExtensionsAvailabilityzonesRequest{}
}

//response struct for the ListComputeV2ExtensionsAvailabilityzones
type ListComputeV2ExtensionsAvailabilityzonesResponse struct{
    Pager pagination.Pager
}

func NewListComputeV2ExtensionsAvailabilityzonesResponse(pager pagination.Pager,)*ListComputeV2ExtensionsAvailabilityzonesResponse {
    return &ListComputeV2ExtensionsAvailabilityzonesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListComputeV2ExtensionsAvailabilityzones(req *ListComputeV2ExtensionsAvailabilityzonesRequest)(*ListComputeV2ExtensionsAvailabilityzonesResponse){
    return NewListComputeV2ExtensionsAvailabilityzonesResponse(availabilityzones.List(oc.client, ))

}
//request struct for the ListDetailComputeV2ExtensionsAvailabilityzones
type ListDetailComputeV2ExtensionsAvailabilityzonesRequest struct{
}

func NewListDetailComputeV2ExtensionsAvailabilityzonesRequest()*ListDetailComputeV2ExtensionsAvailabilityzonesRequest{
    return &ListDetailComputeV2ExtensionsAvailabilityzonesRequest{}
}

//response struct for the ListDetailComputeV2ExtensionsAvailabilityzones
type ListDetailComputeV2ExtensionsAvailabilityzonesResponse struct{
    Pager pagination.Pager
}

func NewListDetailComputeV2ExtensionsAvailabilityzonesResponse(pager pagination.Pager,)*ListDetailComputeV2ExtensionsAvailabilityzonesResponse {
    return &ListDetailComputeV2ExtensionsAvailabilityzonesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListDetailComputeV2ExtensionsAvailabilityzones(req *ListDetailComputeV2ExtensionsAvailabilityzonesRequest)(*ListDetailComputeV2ExtensionsAvailabilityzonesResponse){
    return NewListDetailComputeV2ExtensionsAvailabilityzonesResponse(availabilityzones.ListDetail(oc.client, ))

}