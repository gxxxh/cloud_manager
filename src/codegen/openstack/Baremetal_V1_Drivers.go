package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/baremetal/v1/drivers"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListDriversBaremetalV1Drivers
type ListDriversBaremetalV1DriversRequest struct{
    Opts drivers.ListDriversOptsBuilder
}

func NewListDriversBaremetalV1DriversRequest()*ListDriversBaremetalV1DriversRequest{
    return &ListDriversBaremetalV1DriversRequest{}
}

//response struct for the ListDriversBaremetalV1Drivers
type ListDriversBaremetalV1DriversResponse struct{
    Pager pagination.Pager
}

func NewListDriversBaremetalV1DriversResponse(pager pagination.Pager,)*ListDriversBaremetalV1DriversResponse {
    return &ListDriversBaremetalV1DriversResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListDriversBaremetalV1Drivers(req *ListDriversBaremetalV1DriversRequest)(*ListDriversBaremetalV1DriversResponse){
    return NewListDriversBaremetalV1DriversResponse(drivers.ListDrivers(oc.client,req.Opts, ))

}
//request struct for the GetDriverDetailsBaremetalV1Drivers
type GetDriverDetailsBaremetalV1DriversRequest struct{
    DriverName string
}

func NewGetDriverDetailsBaremetalV1DriversRequest()*GetDriverDetailsBaremetalV1DriversRequest{
    return &GetDriverDetailsBaremetalV1DriversRequest{}
}

//response struct for the GetDriverDetailsBaremetalV1Drivers
type GetDriverDetailsBaremetalV1DriversResponse struct{
    GetDriverResult drivers.GetDriverResult
}

func NewGetDriverDetailsBaremetalV1DriversResponse(getDriverResult drivers.GetDriverResult,)*GetDriverDetailsBaremetalV1DriversResponse {
    return &GetDriverDetailsBaremetalV1DriversResponse{
            GetDriverResult:getDriverResult,
    }
}

// action function
func (oc *OpenstackClient) GetDriverDetailsBaremetalV1Drivers(req *GetDriverDetailsBaremetalV1DriversRequest)(*GetDriverDetailsBaremetalV1DriversResponse){
    return NewGetDriverDetailsBaremetalV1DriversResponse(drivers.GetDriverDetails(oc.client,req.DriverName, ))

}
//request struct for the GetDriverPropertiesBaremetalV1Drivers
type GetDriverPropertiesBaremetalV1DriversRequest struct{
    DriverName string
}

func NewGetDriverPropertiesBaremetalV1DriversRequest()*GetDriverPropertiesBaremetalV1DriversRequest{
    return &GetDriverPropertiesBaremetalV1DriversRequest{}
}

//response struct for the GetDriverPropertiesBaremetalV1Drivers
type GetDriverPropertiesBaremetalV1DriversResponse struct{
    GetPropertiesResult drivers.GetPropertiesResult
}

func NewGetDriverPropertiesBaremetalV1DriversResponse(getPropertiesResult drivers.GetPropertiesResult,)*GetDriverPropertiesBaremetalV1DriversResponse {
    return &GetDriverPropertiesBaremetalV1DriversResponse{
            GetPropertiesResult:getPropertiesResult,
    }
}

// action function
func (oc *OpenstackClient) GetDriverPropertiesBaremetalV1Drivers(req *GetDriverPropertiesBaremetalV1DriversRequest)(*GetDriverPropertiesBaremetalV1DriversResponse){
    return NewGetDriverPropertiesBaremetalV1DriversResponse(drivers.GetDriverProperties(oc.client,req.DriverName, ))

}
//request struct for the GetDriverDiskPropertiesBaremetalV1Drivers
type GetDriverDiskPropertiesBaremetalV1DriversRequest struct{
    DriverName string
}

func NewGetDriverDiskPropertiesBaremetalV1DriversRequest()*GetDriverDiskPropertiesBaremetalV1DriversRequest{
    return &GetDriverDiskPropertiesBaremetalV1DriversRequest{}
}

//response struct for the GetDriverDiskPropertiesBaremetalV1Drivers
type GetDriverDiskPropertiesBaremetalV1DriversResponse struct{
    GetDiskPropertiesResult drivers.GetDiskPropertiesResult
}

func NewGetDriverDiskPropertiesBaremetalV1DriversResponse(getDiskPropertiesResult drivers.GetDiskPropertiesResult,)*GetDriverDiskPropertiesBaremetalV1DriversResponse {
    return &GetDriverDiskPropertiesBaremetalV1DriversResponse{
            GetDiskPropertiesResult:getDiskPropertiesResult,
    }
}

// action function
func (oc *OpenstackClient) GetDriverDiskPropertiesBaremetalV1Drivers(req *GetDriverDiskPropertiesBaremetalV1DriversRequest)(*GetDriverDiskPropertiesBaremetalV1DriversResponse){
    return NewGetDriverDiskPropertiesBaremetalV1DriversResponse(drivers.GetDriverDiskProperties(oc.client,req.DriverName, ))

}