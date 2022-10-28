package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/messages"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the DeleteSharedfilesystemsV2Messages
type DeleteSharedfilesystemsV2MessagesRequest struct{
    Id string
}

func NewDeleteSharedfilesystemsV2MessagesRequest()*DeleteSharedfilesystemsV2MessagesRequest{
    return &DeleteSharedfilesystemsV2MessagesRequest{}
}

//response struct for the DeleteSharedfilesystemsV2Messages
type DeleteSharedfilesystemsV2MessagesResponse struct{
    DeleteResult messages.DeleteResult
}

func NewDeleteSharedfilesystemsV2MessagesResponse(deleteResult messages.DeleteResult,)*DeleteSharedfilesystemsV2MessagesResponse {
    return &DeleteSharedfilesystemsV2MessagesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteSharedfilesystemsV2Messages(req *DeleteSharedfilesystemsV2MessagesRequest)(*DeleteSharedfilesystemsV2MessagesResponse){
    return NewDeleteSharedfilesystemsV2MessagesResponse(messages.Delete(oc.client,req.Id, ))

}
//request struct for the ListSharedfilesystemsV2Messages
type ListSharedfilesystemsV2MessagesRequest struct{
    Opts messages.ListOptsBuilder
}

func NewListSharedfilesystemsV2MessagesRequest()*ListSharedfilesystemsV2MessagesRequest{
    return &ListSharedfilesystemsV2MessagesRequest{}
}

//response struct for the ListSharedfilesystemsV2Messages
type ListSharedfilesystemsV2MessagesResponse struct{
    Pager pagination.Pager
}

func NewListSharedfilesystemsV2MessagesResponse(pager pagination.Pager,)*ListSharedfilesystemsV2MessagesResponse {
    return &ListSharedfilesystemsV2MessagesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListSharedfilesystemsV2Messages(req *ListSharedfilesystemsV2MessagesRequest)(*ListSharedfilesystemsV2MessagesResponse){
    return NewListSharedfilesystemsV2MessagesResponse(messages.List(oc.client,req.Opts, ))

}
//request struct for the GetSharedfilesystemsV2Messages
type GetSharedfilesystemsV2MessagesRequest struct{
    Id string
}

func NewGetSharedfilesystemsV2MessagesRequest()*GetSharedfilesystemsV2MessagesRequest{
    return &GetSharedfilesystemsV2MessagesRequest{}
}

//response struct for the GetSharedfilesystemsV2Messages
type GetSharedfilesystemsV2MessagesResponse struct{
    GetResult messages.GetResult
}

func NewGetSharedfilesystemsV2MessagesResponse(getResult messages.GetResult,)*GetSharedfilesystemsV2MessagesResponse {
    return &GetSharedfilesystemsV2MessagesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetSharedfilesystemsV2Messages(req *GetSharedfilesystemsV2MessagesRequest)(*GetSharedfilesystemsV2MessagesResponse){
    return NewGetSharedfilesystemsV2MessagesResponse(messages.Get(oc.client,req.Id, ))

}