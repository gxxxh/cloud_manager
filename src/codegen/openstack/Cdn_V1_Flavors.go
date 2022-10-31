package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/cdn/v1/flavors"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListCdnV1Flavors
type ListCdnV1FlavorsRequest struct{
}

func NewListCdnV1FlavorsRequest()*ListCdnV1FlavorsRequest{
    return &ListCdnV1FlavorsRequest{}
}

//response struct for the ListCdnV1Flavors
type ListCdnV1FlavorsResponse struct{
    Pager pagination.Pager
}

func NewListCdnV1FlavorsResponse(pager pagination.Pager,)*ListCdnV1FlavorsResponse {
    return &ListCdnV1FlavorsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListCdnV1Flavors(req *ListCdnV1FlavorsRequest)(*ListCdnV1FlavorsResponse){
    return NewListCdnV1FlavorsResponse(flavors.List(oc.Client, ))

}
//request struct for the GetCdnV1Flavors
type GetCdnV1FlavorsRequest struct{
    Id string
}

func NewGetCdnV1FlavorsRequest()*GetCdnV1FlavorsRequest{
    return &GetCdnV1FlavorsRequest{}
}

//response struct for the GetCdnV1Flavors
type GetCdnV1FlavorsResponse struct{
    GetResult flavors.GetResult
}

func NewGetCdnV1FlavorsResponse(getResult flavors.GetResult,)*GetCdnV1FlavorsResponse {
    return &GetCdnV1FlavorsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetCdnV1Flavors(req *GetCdnV1FlavorsRequest)(*GetCdnV1FlavorsResponse){
    return NewGetCdnV1FlavorsResponse(flavors.Get(oc.Client,req.Id, ))

}