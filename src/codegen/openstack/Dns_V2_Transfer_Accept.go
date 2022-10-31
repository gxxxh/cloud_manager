package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/dns/v2/transfer/accept"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListDnsV2TransferAccept
type ListDnsV2TransferAcceptRequest struct{
    Opts accept.ListOpts
}

func NewListDnsV2TransferAcceptRequest()*ListDnsV2TransferAcceptRequest{
    return &ListDnsV2TransferAcceptRequest{}
}

//response struct for the ListDnsV2TransferAccept
type ListDnsV2TransferAcceptResponse struct{
    Pager pagination.Pager
}

func NewListDnsV2TransferAcceptResponse(pager pagination.Pager,)*ListDnsV2TransferAcceptResponse {
    return &ListDnsV2TransferAcceptResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListDnsV2TransferAccept(req *ListDnsV2TransferAcceptRequest)(*ListDnsV2TransferAcceptResponse){
    return NewListDnsV2TransferAcceptResponse(accept.List(oc.Client,req.Opts, ))

}
//request struct for the GetDnsV2TransferAccept
type GetDnsV2TransferAcceptRequest struct{
    TransferAcceptID string
}

func NewGetDnsV2TransferAcceptRequest()*GetDnsV2TransferAcceptRequest{
    return &GetDnsV2TransferAcceptRequest{}
}

//response struct for the GetDnsV2TransferAccept
type GetDnsV2TransferAcceptResponse struct{
    GetResult accept.GetResult
}

func NewGetDnsV2TransferAcceptResponse(getResult accept.GetResult,)*GetDnsV2TransferAcceptResponse {
    return &GetDnsV2TransferAcceptResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetDnsV2TransferAccept(req *GetDnsV2TransferAcceptRequest)(*GetDnsV2TransferAcceptResponse){
    return NewGetDnsV2TransferAcceptResponse(accept.Get(oc.Client,req.TransferAcceptID, ))

}
//request struct for the CreateDnsV2TransferAccept
type CreateDnsV2TransferAcceptRequest struct{
    Opts accept.CreateOpts
}

func NewCreateDnsV2TransferAcceptRequest()*CreateDnsV2TransferAcceptRequest{
    return &CreateDnsV2TransferAcceptRequest{}
}

//response struct for the CreateDnsV2TransferAccept
type CreateDnsV2TransferAcceptResponse struct{
    CreateResult accept.CreateResult
}

func NewCreateDnsV2TransferAcceptResponse(createResult accept.CreateResult,)*CreateDnsV2TransferAcceptResponse {
    return &CreateDnsV2TransferAcceptResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateDnsV2TransferAccept(req *CreateDnsV2TransferAcceptRequest)(*CreateDnsV2TransferAcceptResponse){
    return NewCreateDnsV2TransferAcceptResponse(accept.Create(oc.Client,req.Opts, ))

}