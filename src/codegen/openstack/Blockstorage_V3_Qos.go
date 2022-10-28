package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/blockstorage/v3/qos"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateBlockstorageV3Qos
type CreateBlockstorageV3QosRequest struct{
    Opts qos.CreateOptsBuilder
}

func NewCreateBlockstorageV3QosRequest()*CreateBlockstorageV3QosRequest{
    return &CreateBlockstorageV3QosRequest{}
}

//response struct for the CreateBlockstorageV3Qos
type CreateBlockstorageV3QosResponse struct{
    CreateResult qos.CreateResult
}

func NewCreateBlockstorageV3QosResponse(createResult qos.CreateResult,)*CreateBlockstorageV3QosResponse {
    return &CreateBlockstorageV3QosResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateBlockstorageV3Qos(req *CreateBlockstorageV3QosRequest)(*CreateBlockstorageV3QosResponse){
    return NewCreateBlockstorageV3QosResponse(qos.Create(oc.client,req.Opts, ))

}
//request struct for the DeleteBlockstorageV3Qos
type DeleteBlockstorageV3QosRequest struct{
    Id string
    Opts qos.DeleteOptsBuilder
}

func NewDeleteBlockstorageV3QosRequest()*DeleteBlockstorageV3QosRequest{
    return &DeleteBlockstorageV3QosRequest{}
}

//response struct for the DeleteBlockstorageV3Qos
type DeleteBlockstorageV3QosResponse struct{
    DeleteResult qos.DeleteResult
}

func NewDeleteBlockstorageV3QosResponse(deleteResult qos.DeleteResult,)*DeleteBlockstorageV3QosResponse {
    return &DeleteBlockstorageV3QosResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteBlockstorageV3Qos(req *DeleteBlockstorageV3QosRequest)(*DeleteBlockstorageV3QosResponse){
    return NewDeleteBlockstorageV3QosResponse(qos.Delete(oc.client,req.Id,req.Opts, ))

}
//request struct for the ListBlockstorageV3Qos
type ListBlockstorageV3QosRequest struct{
    Opts qos.ListOptsBuilder
}

func NewListBlockstorageV3QosRequest()*ListBlockstorageV3QosRequest{
    return &ListBlockstorageV3QosRequest{}
}

//response struct for the ListBlockstorageV3Qos
type ListBlockstorageV3QosResponse struct{
    Pager pagination.Pager
}

func NewListBlockstorageV3QosResponse(pager pagination.Pager,)*ListBlockstorageV3QosResponse {
    return &ListBlockstorageV3QosResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListBlockstorageV3Qos(req *ListBlockstorageV3QosRequest)(*ListBlockstorageV3QosResponse){
    return NewListBlockstorageV3QosResponse(qos.List(oc.client,req.Opts, ))

}
//request struct for the GetBlockstorageV3Qos
type GetBlockstorageV3QosRequest struct{
    Id string
}

func NewGetBlockstorageV3QosRequest()*GetBlockstorageV3QosRequest{
    return &GetBlockstorageV3QosRequest{}
}

//response struct for the GetBlockstorageV3Qos
type GetBlockstorageV3QosResponse struct{
    GetResult qos.GetResult
}

func NewGetBlockstorageV3QosResponse(getResult qos.GetResult,)*GetBlockstorageV3QosResponse {
    return &GetBlockstorageV3QosResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetBlockstorageV3Qos(req *GetBlockstorageV3QosRequest)(*GetBlockstorageV3QosResponse){
    return NewGetBlockstorageV3QosResponse(qos.Get(oc.client,req.Id, ))

}
//request struct for the DeleteKeysBlockstorageV3Qos
type DeleteKeysBlockstorageV3QosRequest struct{
    QosID string
    Opts qos.DeleteKeysOptsBuilder
}

func NewDeleteKeysBlockstorageV3QosRequest()*DeleteKeysBlockstorageV3QosRequest{
    return &DeleteKeysBlockstorageV3QosRequest{}
}

//response struct for the DeleteKeysBlockstorageV3Qos
type DeleteKeysBlockstorageV3QosResponse struct{
    DeleteResult qos.DeleteResult
}

func NewDeleteKeysBlockstorageV3QosResponse(deleteResult qos.DeleteResult,)*DeleteKeysBlockstorageV3QosResponse {
    return &DeleteKeysBlockstorageV3QosResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteKeysBlockstorageV3Qos(req *DeleteKeysBlockstorageV3QosRequest)(*DeleteKeysBlockstorageV3QosResponse){
    return NewDeleteKeysBlockstorageV3QosResponse(qos.DeleteKeys(oc.client,req.QosID,req.Opts, ))

}
//request struct for the AssociateBlockstorageV3Qos
type AssociateBlockstorageV3QosRequest struct{
    QosID string
    Opts qos.AssociateOptsBuilder
}

func NewAssociateBlockstorageV3QosRequest()*AssociateBlockstorageV3QosRequest{
    return &AssociateBlockstorageV3QosRequest{}
}

//response struct for the AssociateBlockstorageV3Qos
type AssociateBlockstorageV3QosResponse struct{
    AssociateResult qos.AssociateResult
}

func NewAssociateBlockstorageV3QosResponse(associateResult qos.AssociateResult,)*AssociateBlockstorageV3QosResponse {
    return &AssociateBlockstorageV3QosResponse{
            AssociateResult:associateResult,
    }
}

// action function
func (oc *OpenstackClient) AssociateBlockstorageV3Qos(req *AssociateBlockstorageV3QosRequest)(*AssociateBlockstorageV3QosResponse){
    return NewAssociateBlockstorageV3QosResponse(qos.Associate(oc.client,req.QosID,req.Opts, ))

}
//request struct for the DisassociateBlockstorageV3Qos
type DisassociateBlockstorageV3QosRequest struct{
    QosID string
    Opts qos.DisassociateOptsBuilder
}

func NewDisassociateBlockstorageV3QosRequest()*DisassociateBlockstorageV3QosRequest{
    return &DisassociateBlockstorageV3QosRequest{}
}

//response struct for the DisassociateBlockstorageV3Qos
type DisassociateBlockstorageV3QosResponse struct{
    DisassociateResult qos.DisassociateResult
}

func NewDisassociateBlockstorageV3QosResponse(disassociateResult qos.DisassociateResult,)*DisassociateBlockstorageV3QosResponse {
    return &DisassociateBlockstorageV3QosResponse{
            DisassociateResult:disassociateResult,
    }
}

// action function
func (oc *OpenstackClient) DisassociateBlockstorageV3Qos(req *DisassociateBlockstorageV3QosRequest)(*DisassociateBlockstorageV3QosResponse){
    return NewDisassociateBlockstorageV3QosResponse(qos.Disassociate(oc.client,req.QosID,req.Opts, ))

}
//request struct for the DisassociateAllBlockstorageV3Qos
type DisassociateAllBlockstorageV3QosRequest struct{
    QosID string
}

func NewDisassociateAllBlockstorageV3QosRequest()*DisassociateAllBlockstorageV3QosRequest{
    return &DisassociateAllBlockstorageV3QosRequest{}
}

//response struct for the DisassociateAllBlockstorageV3Qos
type DisassociateAllBlockstorageV3QosResponse struct{
    DisassociateAllResult qos.DisassociateAllResult
}

func NewDisassociateAllBlockstorageV3QosResponse(disassociateAllResult qos.DisassociateAllResult,)*DisassociateAllBlockstorageV3QosResponse {
    return &DisassociateAllBlockstorageV3QosResponse{
            DisassociateAllResult:disassociateAllResult,
    }
}

// action function
func (oc *OpenstackClient) DisassociateAllBlockstorageV3Qos(req *DisassociateAllBlockstorageV3QosRequest)(*DisassociateAllBlockstorageV3QosResponse){
    return NewDisassociateAllBlockstorageV3QosResponse(qos.DisassociateAll(oc.client,req.QosID, ))

}
//request struct for the ListAssociationsBlockstorageV3Qos
type ListAssociationsBlockstorageV3QosRequest struct{
    QosID string
}

func NewListAssociationsBlockstorageV3QosRequest()*ListAssociationsBlockstorageV3QosRequest{
    return &ListAssociationsBlockstorageV3QosRequest{}
}

//response struct for the ListAssociationsBlockstorageV3Qos
type ListAssociationsBlockstorageV3QosResponse struct{
    Pager pagination.Pager
}

func NewListAssociationsBlockstorageV3QosResponse(pager pagination.Pager,)*ListAssociationsBlockstorageV3QosResponse {
    return &ListAssociationsBlockstorageV3QosResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListAssociationsBlockstorageV3Qos(req *ListAssociationsBlockstorageV3QosRequest)(*ListAssociationsBlockstorageV3QosResponse){
    return NewListAssociationsBlockstorageV3QosResponse(qos.ListAssociations(oc.client,req.QosID, ))

}