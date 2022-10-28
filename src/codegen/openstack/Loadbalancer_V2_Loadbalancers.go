package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/loadbalancers"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListLoadbalancerV2Loadbalancers
type ListLoadbalancerV2LoadbalancersRequest struct{
    Opts loadbalancers.ListOptsBuilder
}

func NewListLoadbalancerV2LoadbalancersRequest()*ListLoadbalancerV2LoadbalancersRequest{
    return &ListLoadbalancerV2LoadbalancersRequest{}
}

//response struct for the ListLoadbalancerV2Loadbalancers
type ListLoadbalancerV2LoadbalancersResponse struct{
    Pager pagination.Pager
}

func NewListLoadbalancerV2LoadbalancersResponse(pager pagination.Pager,)*ListLoadbalancerV2LoadbalancersResponse {
    return &ListLoadbalancerV2LoadbalancersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListLoadbalancerV2Loadbalancers(req *ListLoadbalancerV2LoadbalancersRequest)(*ListLoadbalancerV2LoadbalancersResponse){
    return NewListLoadbalancerV2LoadbalancersResponse(loadbalancers.List(oc.client,req.Opts, ))

}
//request struct for the CreateLoadbalancerV2Loadbalancers
type CreateLoadbalancerV2LoadbalancersRequest struct{
    Opts loadbalancers.CreateOptsBuilder
}

func NewCreateLoadbalancerV2LoadbalancersRequest()*CreateLoadbalancerV2LoadbalancersRequest{
    return &CreateLoadbalancerV2LoadbalancersRequest{}
}

//response struct for the CreateLoadbalancerV2Loadbalancers
type CreateLoadbalancerV2LoadbalancersResponse struct{
    CreateResult loadbalancers.CreateResult
}

func NewCreateLoadbalancerV2LoadbalancersResponse(createResult loadbalancers.CreateResult,)*CreateLoadbalancerV2LoadbalancersResponse {
    return &CreateLoadbalancerV2LoadbalancersResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateLoadbalancerV2Loadbalancers(req *CreateLoadbalancerV2LoadbalancersRequest)(*CreateLoadbalancerV2LoadbalancersResponse){
    return NewCreateLoadbalancerV2LoadbalancersResponse(loadbalancers.Create(oc.client,req.Opts, ))

}
//request struct for the GetLoadbalancerV2Loadbalancers
type GetLoadbalancerV2LoadbalancersRequest struct{
    Id string
}

func NewGetLoadbalancerV2LoadbalancersRequest()*GetLoadbalancerV2LoadbalancersRequest{
    return &GetLoadbalancerV2LoadbalancersRequest{}
}

//response struct for the GetLoadbalancerV2Loadbalancers
type GetLoadbalancerV2LoadbalancersResponse struct{
    GetResult loadbalancers.GetResult
}

func NewGetLoadbalancerV2LoadbalancersResponse(getResult loadbalancers.GetResult,)*GetLoadbalancerV2LoadbalancersResponse {
    return &GetLoadbalancerV2LoadbalancersResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetLoadbalancerV2Loadbalancers(req *GetLoadbalancerV2LoadbalancersRequest)(*GetLoadbalancerV2LoadbalancersResponse){
    return NewGetLoadbalancerV2LoadbalancersResponse(loadbalancers.Get(oc.client,req.Id, ))

}
//request struct for the UpdateLoadbalancerV2Loadbalancers
type UpdateLoadbalancerV2LoadbalancersRequest struct{
    Id string
    Opts loadbalancers.UpdateOpts
}

func NewUpdateLoadbalancerV2LoadbalancersRequest()*UpdateLoadbalancerV2LoadbalancersRequest{
    return &UpdateLoadbalancerV2LoadbalancersRequest{}
}

//response struct for the UpdateLoadbalancerV2Loadbalancers
type UpdateLoadbalancerV2LoadbalancersResponse struct{
    UpdateResult loadbalancers.UpdateResult
}

func NewUpdateLoadbalancerV2LoadbalancersResponse(updateResult loadbalancers.UpdateResult,)*UpdateLoadbalancerV2LoadbalancersResponse {
    return &UpdateLoadbalancerV2LoadbalancersResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateLoadbalancerV2Loadbalancers(req *UpdateLoadbalancerV2LoadbalancersRequest)(*UpdateLoadbalancerV2LoadbalancersResponse){
    return NewUpdateLoadbalancerV2LoadbalancersResponse(loadbalancers.Update(oc.client,req.Id,req.Opts, ))

}
//request struct for the DeleteLoadbalancerV2Loadbalancers
type DeleteLoadbalancerV2LoadbalancersRequest struct{
    Id string
    Opts loadbalancers.DeleteOptsBuilder
}

func NewDeleteLoadbalancerV2LoadbalancersRequest()*DeleteLoadbalancerV2LoadbalancersRequest{
    return &DeleteLoadbalancerV2LoadbalancersRequest{}
}

//response struct for the DeleteLoadbalancerV2Loadbalancers
type DeleteLoadbalancerV2LoadbalancersResponse struct{
    DeleteResult loadbalancers.DeleteResult
}

func NewDeleteLoadbalancerV2LoadbalancersResponse(deleteResult loadbalancers.DeleteResult,)*DeleteLoadbalancerV2LoadbalancersResponse {
    return &DeleteLoadbalancerV2LoadbalancersResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteLoadbalancerV2Loadbalancers(req *DeleteLoadbalancerV2LoadbalancersRequest)(*DeleteLoadbalancerV2LoadbalancersResponse){
    return NewDeleteLoadbalancerV2LoadbalancersResponse(loadbalancers.Delete(oc.client,req.Id,req.Opts, ))

}
//request struct for the GetStatusesLoadbalancerV2Loadbalancers
type GetStatusesLoadbalancerV2LoadbalancersRequest struct{
    Id string
}

func NewGetStatusesLoadbalancerV2LoadbalancersRequest()*GetStatusesLoadbalancerV2LoadbalancersRequest{
    return &GetStatusesLoadbalancerV2LoadbalancersRequest{}
}

//response struct for the GetStatusesLoadbalancerV2Loadbalancers
type GetStatusesLoadbalancerV2LoadbalancersResponse struct{
    GetStatusesResult loadbalancers.GetStatusesResult
}

func NewGetStatusesLoadbalancerV2LoadbalancersResponse(getStatusesResult loadbalancers.GetStatusesResult,)*GetStatusesLoadbalancerV2LoadbalancersResponse {
    return &GetStatusesLoadbalancerV2LoadbalancersResponse{
            GetStatusesResult:getStatusesResult,
    }
}

// action function
func (oc *OpenstackClient) GetStatusesLoadbalancerV2Loadbalancers(req *GetStatusesLoadbalancerV2LoadbalancersRequest)(*GetStatusesLoadbalancerV2LoadbalancersResponse){
    return NewGetStatusesLoadbalancerV2LoadbalancersResponse(loadbalancers.GetStatuses(oc.client,req.Id, ))

}
//request struct for the GetStatsLoadbalancerV2Loadbalancers
type GetStatsLoadbalancerV2LoadbalancersRequest struct{
    Id string
}

func NewGetStatsLoadbalancerV2LoadbalancersRequest()*GetStatsLoadbalancerV2LoadbalancersRequest{
    return &GetStatsLoadbalancerV2LoadbalancersRequest{}
}

//response struct for the GetStatsLoadbalancerV2Loadbalancers
type GetStatsLoadbalancerV2LoadbalancersResponse struct{
    StatsResult loadbalancers.StatsResult
}

func NewGetStatsLoadbalancerV2LoadbalancersResponse(statsResult loadbalancers.StatsResult,)*GetStatsLoadbalancerV2LoadbalancersResponse {
    return &GetStatsLoadbalancerV2LoadbalancersResponse{
            StatsResult:statsResult,
    }
}

// action function
func (oc *OpenstackClient) GetStatsLoadbalancerV2Loadbalancers(req *GetStatsLoadbalancerV2LoadbalancersRequest)(*GetStatsLoadbalancerV2LoadbalancersResponse){
    return NewGetStatsLoadbalancerV2LoadbalancersResponse(loadbalancers.GetStats(oc.client,req.Id, ))

}
//request struct for the FailoverLoadbalancerV2Loadbalancers
type FailoverLoadbalancerV2LoadbalancersRequest struct{
    Id string
}

func NewFailoverLoadbalancerV2LoadbalancersRequest()*FailoverLoadbalancerV2LoadbalancersRequest{
    return &FailoverLoadbalancerV2LoadbalancersRequest{}
}

//response struct for the FailoverLoadbalancerV2Loadbalancers
type FailoverLoadbalancerV2LoadbalancersResponse struct{
    FailoverResult loadbalancers.FailoverResult
}

func NewFailoverLoadbalancerV2LoadbalancersResponse(failoverResult loadbalancers.FailoverResult,)*FailoverLoadbalancerV2LoadbalancersResponse {
    return &FailoverLoadbalancerV2LoadbalancersResponse{
            FailoverResult:failoverResult,
    }
}

// action function
func (oc *OpenstackClient) FailoverLoadbalancerV2Loadbalancers(req *FailoverLoadbalancerV2LoadbalancersRequest)(*FailoverLoadbalancerV2LoadbalancersResponse){
    return NewFailoverLoadbalancerV2LoadbalancersResponse(loadbalancers.Failover(oc.client,req.Id, ))

}