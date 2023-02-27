

### 部分复杂的操作无法生成，需要手动添加，
以imageService的updateProperties为例，需要手动实现为所需的形式,如何实现参考gophercloud的源代码，手动添加的主要原因是原有方法的参数为interface类型，但无法定位到
实现该类型的类
```go
// 手动添加
type UpdateImageserviceV2ImagesProperty struct {
	Op    images.UpdateOp
	Name  string
	Value string
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

```