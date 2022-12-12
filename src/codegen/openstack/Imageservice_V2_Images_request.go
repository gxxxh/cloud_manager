package openstack

// Code generated by cloud manager.

import (
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	"github.com/gophercloud/gophercloud/pagination"
)

// request struct for the ListImageserviceV2Images
type ListImageserviceV2ImagesRequest struct {
	Opts images.ListOpts
}

func NewListImageserviceV2ImagesRequest() *ListImageserviceV2ImagesRequest {
	return &ListImageserviceV2ImagesRequest{}
}

// response struct for the ListImageserviceV2Images
type ListImageserviceV2ImagesResponse struct {
	Pager pagination.Pager
}

func NewListImageserviceV2ImagesResponse(pager pagination.Pager) *ListImageserviceV2ImagesResponse {
	return &ListImageserviceV2ImagesResponse{
		Pager: pager,
	}
}

// action function
func (oc *OpenstackClient) ListImageserviceV2Images(req *ListImageserviceV2ImagesRequest) *ListImageserviceV2ImagesResponse {
	return NewListImageserviceV2ImagesResponse(images.List(oc.Client, req.Opts))

}

// request struct for the CreateImageserviceV2Images
type CreateImageserviceV2ImagesRequest struct {
	Opts images.CreateOpts
}

func NewCreateImageserviceV2ImagesRequest() *CreateImageserviceV2ImagesRequest {
	return &CreateImageserviceV2ImagesRequest{}
}

// response struct for the CreateImageserviceV2Images
type CreateImageserviceV2ImagesResponse struct {
	CreateResult images.CreateResult
}

func NewCreateImageserviceV2ImagesResponse(createResult images.CreateResult) *CreateImageserviceV2ImagesResponse {
	return &CreateImageserviceV2ImagesResponse{
		CreateResult: createResult,
	}
}

// action function
func (oc *OpenstackClient) CreateImageserviceV2Images(req *CreateImageserviceV2ImagesRequest) *CreateImageserviceV2ImagesResponse {
	return NewCreateImageserviceV2ImagesResponse(images.Create(oc.Client, req.Opts))

}

// request struct for the DeleteImageserviceV2Images
type DeleteImageserviceV2ImagesRequest struct {
	Id string
}

func NewDeleteImageserviceV2ImagesRequest() *DeleteImageserviceV2ImagesRequest {
	return &DeleteImageserviceV2ImagesRequest{}
}

// response struct for the DeleteImageserviceV2Images
type DeleteImageserviceV2ImagesResponse struct {
	DeleteResult images.DeleteResult
}

func NewDeleteImageserviceV2ImagesResponse(deleteResult images.DeleteResult) *DeleteImageserviceV2ImagesResponse {
	return &DeleteImageserviceV2ImagesResponse{
		DeleteResult: deleteResult,
	}
}

// action function
func (oc *OpenstackClient) DeleteImageserviceV2Images(req *DeleteImageserviceV2ImagesRequest) *DeleteImageserviceV2ImagesResponse {
	return NewDeleteImageserviceV2ImagesResponse(images.Delete(oc.Client, req.Id))

}

// request struct for the GetImageserviceV2Images
type GetImageserviceV2ImagesRequest struct {
	Id string
}

func NewGetImageserviceV2ImagesRequest() *GetImageserviceV2ImagesRequest {
	return &GetImageserviceV2ImagesRequest{}
}

// response struct for the GetImageserviceV2Images
type GetImageserviceV2ImagesResponse struct {
	GetResult images.GetResult
}

func NewGetImageserviceV2ImagesResponse(getResult images.GetResult) *GetImageserviceV2ImagesResponse {
	return &GetImageserviceV2ImagesResponse{
		GetResult: getResult,
	}
}

// action function
func (oc *OpenstackClient) GetImageserviceV2Images(req *GetImageserviceV2ImagesRequest) *GetImageserviceV2ImagesResponse {
	return NewGetImageserviceV2ImagesResponse(images.Get(oc.Client, req.Id))

}

type UpdateImageserviceV2ImagesProperty struct {
	Op    images.UpdateOp
	Name  string
	Value interface{}
}

func (r UpdateImageserviceV2ImagesProperty) ToImagePatchMap() map[string]interface{} {
	updateMap := map[string]interface{}{
		"op":   r.Op,
		"path": fmt.Sprintf("/%s", r.Name),
	}

	if r.Op != images.RemoveOp {
		updateMap["value"] = r.Value
	}

	return updateMap
}

type UpdateImageserviceV2ImagesOpts []UpdateImageserviceV2ImagesProperty

func (opts UpdateImageserviceV2ImagesOpts) ToImageUpdateMap() ([]interface{}, error) {
	m := make([]interface{}, len(opts))
	for i, patch := range opts {
		patchJSON := patch.ToImagePatchMap()
		m[i] = patchJSON
	}
	return m, nil
}

// request struct for the UpdateImageserviceV2Images
type UpdateImageserviceV2ImagesRequest struct {
	Id   string
	Opts UpdateImageserviceV2ImagesOpts
}

func NewUpdateImageserviceV2ImagesRequest() *UpdateImageserviceV2ImagesRequest {
	return &UpdateImageserviceV2ImagesRequest{}
}

// response struct for the UpdateImageserviceV2Images
type UpdateImageserviceV2ImagesResponse struct {
	UpdateResult images.UpdateResult
}

func NewUpdateImageserviceV2ImagesResponse(updateResult images.UpdateResult) *UpdateImageserviceV2ImagesResponse {
	return &UpdateImageserviceV2ImagesResponse{
		UpdateResult: updateResult,
	}
}

// action function
func (oc *OpenstackClient) UpdateImageserviceV2Images(req *UpdateImageserviceV2ImagesRequest) *UpdateImageserviceV2ImagesResponse {
	return NewUpdateImageserviceV2ImagesResponse(images.Update(oc.Client, req.Id, req.Opts))

}
