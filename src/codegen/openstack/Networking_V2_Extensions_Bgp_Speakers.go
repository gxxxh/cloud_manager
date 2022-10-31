package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/bgp/speakers"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListNetworkingV2ExtensionsBgpSpeakers
type ListNetworkingV2ExtensionsBgpSpeakersRequest struct{
}

func NewListNetworkingV2ExtensionsBgpSpeakersRequest()*ListNetworkingV2ExtensionsBgpSpeakersRequest{
    return &ListNetworkingV2ExtensionsBgpSpeakersRequest{}
}

//response struct for the ListNetworkingV2ExtensionsBgpSpeakers
type ListNetworkingV2ExtensionsBgpSpeakersResponse struct{
    Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsBgpSpeakersResponse(pager pagination.Pager,)*ListNetworkingV2ExtensionsBgpSpeakersResponse {
    return &ListNetworkingV2ExtensionsBgpSpeakersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsBgpSpeakers(req *ListNetworkingV2ExtensionsBgpSpeakersRequest)(*ListNetworkingV2ExtensionsBgpSpeakersResponse){
    return NewListNetworkingV2ExtensionsBgpSpeakersResponse(speakers.List(oc.Client, ))

}
//request struct for the GetNetworkingV2ExtensionsBgpSpeakers
type GetNetworkingV2ExtensionsBgpSpeakersRequest struct{
    Id string
}

func NewGetNetworkingV2ExtensionsBgpSpeakersRequest()*GetNetworkingV2ExtensionsBgpSpeakersRequest{
    return &GetNetworkingV2ExtensionsBgpSpeakersRequest{}
}

//response struct for the GetNetworkingV2ExtensionsBgpSpeakers
type GetNetworkingV2ExtensionsBgpSpeakersResponse struct{
    GetResult speakers.GetResult
}

func NewGetNetworkingV2ExtensionsBgpSpeakersResponse(getResult speakers.GetResult,)*GetNetworkingV2ExtensionsBgpSpeakersResponse {
    return &GetNetworkingV2ExtensionsBgpSpeakersResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsBgpSpeakers(req *GetNetworkingV2ExtensionsBgpSpeakersRequest)(*GetNetworkingV2ExtensionsBgpSpeakersResponse){
    return NewGetNetworkingV2ExtensionsBgpSpeakersResponse(speakers.Get(oc.Client,req.Id, ))

}
//request struct for the CreateNetworkingV2ExtensionsBgpSpeakers
type CreateNetworkingV2ExtensionsBgpSpeakersRequest struct{
    Opts speakers.CreateOpts
}

func NewCreateNetworkingV2ExtensionsBgpSpeakersRequest()*CreateNetworkingV2ExtensionsBgpSpeakersRequest{
    return &CreateNetworkingV2ExtensionsBgpSpeakersRequest{}
}

//response struct for the CreateNetworkingV2ExtensionsBgpSpeakers
type CreateNetworkingV2ExtensionsBgpSpeakersResponse struct{
    CreateResult speakers.CreateResult
}

func NewCreateNetworkingV2ExtensionsBgpSpeakersResponse(createResult speakers.CreateResult,)*CreateNetworkingV2ExtensionsBgpSpeakersResponse {
    return &CreateNetworkingV2ExtensionsBgpSpeakersResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsBgpSpeakers(req *CreateNetworkingV2ExtensionsBgpSpeakersRequest)(*CreateNetworkingV2ExtensionsBgpSpeakersResponse){
    return NewCreateNetworkingV2ExtensionsBgpSpeakersResponse(speakers.Create(oc.Client,req.Opts, ))

}
//request struct for the DeleteNetworkingV2ExtensionsBgpSpeakers
type DeleteNetworkingV2ExtensionsBgpSpeakersRequest struct{
    SpeakerID string
}

func NewDeleteNetworkingV2ExtensionsBgpSpeakersRequest()*DeleteNetworkingV2ExtensionsBgpSpeakersRequest{
    return &DeleteNetworkingV2ExtensionsBgpSpeakersRequest{}
}

//response struct for the DeleteNetworkingV2ExtensionsBgpSpeakers
type DeleteNetworkingV2ExtensionsBgpSpeakersResponse struct{
    DeleteResult speakers.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsBgpSpeakersResponse(deleteResult speakers.DeleteResult,)*DeleteNetworkingV2ExtensionsBgpSpeakersResponse {
    return &DeleteNetworkingV2ExtensionsBgpSpeakersResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsBgpSpeakers(req *DeleteNetworkingV2ExtensionsBgpSpeakersRequest)(*DeleteNetworkingV2ExtensionsBgpSpeakersResponse){
    return NewDeleteNetworkingV2ExtensionsBgpSpeakersResponse(speakers.Delete(oc.Client,req.SpeakerID, ))

}
//request struct for the UpdateNetworkingV2ExtensionsBgpSpeakers
type UpdateNetworkingV2ExtensionsBgpSpeakersRequest struct{
    SpeakerID string
    Opts speakers.UpdateOpts
}

func NewUpdateNetworkingV2ExtensionsBgpSpeakersRequest()*UpdateNetworkingV2ExtensionsBgpSpeakersRequest{
    return &UpdateNetworkingV2ExtensionsBgpSpeakersRequest{}
}

//response struct for the UpdateNetworkingV2ExtensionsBgpSpeakers
type UpdateNetworkingV2ExtensionsBgpSpeakersResponse struct{
    UpdateResult speakers.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsBgpSpeakersResponse(updateResult speakers.UpdateResult,)*UpdateNetworkingV2ExtensionsBgpSpeakersResponse {
    return &UpdateNetworkingV2ExtensionsBgpSpeakersResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsBgpSpeakers(req *UpdateNetworkingV2ExtensionsBgpSpeakersRequest)(*UpdateNetworkingV2ExtensionsBgpSpeakersResponse){
    return NewUpdateNetworkingV2ExtensionsBgpSpeakersResponse(speakers.Update(oc.Client,req.SpeakerID,req.Opts, ))

}
//request struct for the AddBGPPeerNetworkingV2ExtensionsBgpSpeakers
type AddBGPPeerNetworkingV2ExtensionsBgpSpeakersRequest struct{
    BgpSpeakerID string
    Opts *speakers.AddBGPPeerOpts
}

func NewAddBGPPeerNetworkingV2ExtensionsBgpSpeakersRequest()*AddBGPPeerNetworkingV2ExtensionsBgpSpeakersRequest{
    return &AddBGPPeerNetworkingV2ExtensionsBgpSpeakersRequest{}
}

//response struct for the AddBGPPeerNetworkingV2ExtensionsBgpSpeakers
type AddBGPPeerNetworkingV2ExtensionsBgpSpeakersResponse struct{
    AddBGPPeerResult speakers.AddBGPPeerResult
}

func NewAddBGPPeerNetworkingV2ExtensionsBgpSpeakersResponse(addBGPPeerResult speakers.AddBGPPeerResult,)*AddBGPPeerNetworkingV2ExtensionsBgpSpeakersResponse {
    return &AddBGPPeerNetworkingV2ExtensionsBgpSpeakersResponse{
            AddBGPPeerResult:addBGPPeerResult,
    }
}

// action function
func (oc *OpenstackClient) AddBGPPeerNetworkingV2ExtensionsBgpSpeakers(req *AddBGPPeerNetworkingV2ExtensionsBgpSpeakersRequest)(*AddBGPPeerNetworkingV2ExtensionsBgpSpeakersResponse){
    return NewAddBGPPeerNetworkingV2ExtensionsBgpSpeakersResponse(speakers.AddBGPPeer(oc.Client,req.BgpSpeakerID,req.Opts, ))

}
//request struct for the RemoveBGPPeerNetworkingV2ExtensionsBgpSpeakers
type RemoveBGPPeerNetworkingV2ExtensionsBgpSpeakersRequest struct{
    BgpSpeakerID string
    Opts speakers.RemoveBGPPeerOpts
}

func NewRemoveBGPPeerNetworkingV2ExtensionsBgpSpeakersRequest()*RemoveBGPPeerNetworkingV2ExtensionsBgpSpeakersRequest{
    return &RemoveBGPPeerNetworkingV2ExtensionsBgpSpeakersRequest{}
}

//response struct for the RemoveBGPPeerNetworkingV2ExtensionsBgpSpeakers
type RemoveBGPPeerNetworkingV2ExtensionsBgpSpeakersResponse struct{
    RemoveBGPPeerResult speakers.RemoveBGPPeerResult
}

func NewRemoveBGPPeerNetworkingV2ExtensionsBgpSpeakersResponse(removeBGPPeerResult speakers.RemoveBGPPeerResult,)*RemoveBGPPeerNetworkingV2ExtensionsBgpSpeakersResponse {
    return &RemoveBGPPeerNetworkingV2ExtensionsBgpSpeakersResponse{
            RemoveBGPPeerResult:removeBGPPeerResult,
    }
}

// action function
func (oc *OpenstackClient) RemoveBGPPeerNetworkingV2ExtensionsBgpSpeakers(req *RemoveBGPPeerNetworkingV2ExtensionsBgpSpeakersRequest)(*RemoveBGPPeerNetworkingV2ExtensionsBgpSpeakersResponse){
    return NewRemoveBGPPeerNetworkingV2ExtensionsBgpSpeakersResponse(speakers.RemoveBGPPeer(oc.Client,req.BgpSpeakerID,req.Opts, ))

}
//request struct for the GetAdvertisedRoutesNetworkingV2ExtensionsBgpSpeakers
type GetAdvertisedRoutesNetworkingV2ExtensionsBgpSpeakersRequest struct{
    BgpSpeakerID string
}

func NewGetAdvertisedRoutesNetworkingV2ExtensionsBgpSpeakersRequest()*GetAdvertisedRoutesNetworkingV2ExtensionsBgpSpeakersRequest{
    return &GetAdvertisedRoutesNetworkingV2ExtensionsBgpSpeakersRequest{}
}

//response struct for the GetAdvertisedRoutesNetworkingV2ExtensionsBgpSpeakers
type GetAdvertisedRoutesNetworkingV2ExtensionsBgpSpeakersResponse struct{
    Pager pagination.Pager
}

func NewGetAdvertisedRoutesNetworkingV2ExtensionsBgpSpeakersResponse(pager pagination.Pager,)*GetAdvertisedRoutesNetworkingV2ExtensionsBgpSpeakersResponse {
    return &GetAdvertisedRoutesNetworkingV2ExtensionsBgpSpeakersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) GetAdvertisedRoutesNetworkingV2ExtensionsBgpSpeakers(req *GetAdvertisedRoutesNetworkingV2ExtensionsBgpSpeakersRequest)(*GetAdvertisedRoutesNetworkingV2ExtensionsBgpSpeakersResponse){
    return NewGetAdvertisedRoutesNetworkingV2ExtensionsBgpSpeakersResponse(speakers.GetAdvertisedRoutes(oc.Client,req.BgpSpeakerID, ))

}
//request struct for the AddGatewayNetworkNetworkingV2ExtensionsBgpSpeakers
type AddGatewayNetworkNetworkingV2ExtensionsBgpSpeakersRequest struct{
    BgpSpeakerID string
    Opts *speakers.AddGatewayNetworkOpts
}

func NewAddGatewayNetworkNetworkingV2ExtensionsBgpSpeakersRequest()*AddGatewayNetworkNetworkingV2ExtensionsBgpSpeakersRequest{
    return &AddGatewayNetworkNetworkingV2ExtensionsBgpSpeakersRequest{}
}

//response struct for the AddGatewayNetworkNetworkingV2ExtensionsBgpSpeakers
type AddGatewayNetworkNetworkingV2ExtensionsBgpSpeakersResponse struct{
    AddGatewayNetworkResult speakers.AddGatewayNetworkResult
}

func NewAddGatewayNetworkNetworkingV2ExtensionsBgpSpeakersResponse(addGatewayNetworkResult speakers.AddGatewayNetworkResult,)*AddGatewayNetworkNetworkingV2ExtensionsBgpSpeakersResponse {
    return &AddGatewayNetworkNetworkingV2ExtensionsBgpSpeakersResponse{
            AddGatewayNetworkResult:addGatewayNetworkResult,
    }
}

// action function
func (oc *OpenstackClient) AddGatewayNetworkNetworkingV2ExtensionsBgpSpeakers(req *AddGatewayNetworkNetworkingV2ExtensionsBgpSpeakersRequest)(*AddGatewayNetworkNetworkingV2ExtensionsBgpSpeakersResponse){
    return NewAddGatewayNetworkNetworkingV2ExtensionsBgpSpeakersResponse(speakers.AddGatewayNetwork(oc.Client,req.BgpSpeakerID,req.Opts, ))

}
//request struct for the RemoveGatewayNetworkNetworkingV2ExtensionsBgpSpeakers
type RemoveGatewayNetworkNetworkingV2ExtensionsBgpSpeakersRequest struct{
    BgpSpeakerID string
    Opts speakers.RemoveGatewayNetworkOpts
}

func NewRemoveGatewayNetworkNetworkingV2ExtensionsBgpSpeakersRequest()*RemoveGatewayNetworkNetworkingV2ExtensionsBgpSpeakersRequest{
    return &RemoveGatewayNetworkNetworkingV2ExtensionsBgpSpeakersRequest{}
}

//response struct for the RemoveGatewayNetworkNetworkingV2ExtensionsBgpSpeakers
type RemoveGatewayNetworkNetworkingV2ExtensionsBgpSpeakersResponse struct{
    RemoveGatewayNetworkResult speakers.RemoveGatewayNetworkResult
}

func NewRemoveGatewayNetworkNetworkingV2ExtensionsBgpSpeakersResponse(removeGatewayNetworkResult speakers.RemoveGatewayNetworkResult,)*RemoveGatewayNetworkNetworkingV2ExtensionsBgpSpeakersResponse {
    return &RemoveGatewayNetworkNetworkingV2ExtensionsBgpSpeakersResponse{
            RemoveGatewayNetworkResult:removeGatewayNetworkResult,
    }
}

// action function
func (oc *OpenstackClient) RemoveGatewayNetworkNetworkingV2ExtensionsBgpSpeakers(req *RemoveGatewayNetworkNetworkingV2ExtensionsBgpSpeakersRequest)(*RemoveGatewayNetworkNetworkingV2ExtensionsBgpSpeakersResponse){
    return NewRemoveGatewayNetworkNetworkingV2ExtensionsBgpSpeakersResponse(speakers.RemoveGatewayNetwork(oc.Client,req.BgpSpeakerID,req.Opts, ))

}