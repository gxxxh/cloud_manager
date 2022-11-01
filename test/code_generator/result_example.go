package code_generator

//func ExtractListDetailComputeV2ImagesResponse(response *openstack.ListDetailComputeV2ImagesResponse) ([]images.Image, error) {
//	res := make([]images.Image, 0)
//	page, err := response.Pager.AllPages()
//	if err != nil {
//		return res, err
//	}
//	tmpList, err := images.ExtractImages(page)
//	if err != nil {
//		return res, err
//	}
//	res = append(res, tmpList...)
//	return res, nil
//}
//
//func ExtractGetComputeV2ImagesResponse(response *openstack.GetComputeV2ImagesResponse) (*images.Image, error) {
//	return response.GetResult.Extract()
//}
