package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/baremetalintrospection/v1/introspection"
)

// extract response info from pager for ListIntrospectionsBaremetalintrospectionV1Introspection
func ExtractListIntrospectionsBaremetalintrospectionV1IntrospectionResponse(response *ListIntrospectionsBaremetalintrospectionV1IntrospectionResponse) ([]introspection.Introspection, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return introspection.ExtractIntrospections(page)
}

// call result's extract function
func ExtractGetIntrospectionStatusBaremetalintrospectionV1IntrospectionResponse(response *GetIntrospectionStatusBaremetalintrospectionV1IntrospectionResponse) (interface{}, error) {
	return response.GetIntrospectionStatusResult.Body, response.GetIntrospectionStatusResult.Err
}

// call result's extract function
func ExtractStartIntrospectionBaremetalintrospectionV1IntrospectionResponse(response *StartIntrospectionBaremetalintrospectionV1IntrospectionResponse) (interface{}, error) {
	return response.StartResult.Body, response.StartResult.Err
}

// call result's extract function
func ExtractAbortIntrospectionBaremetalintrospectionV1IntrospectionResponse(response *AbortIntrospectionBaremetalintrospectionV1IntrospectionResponse) (interface{}, error) {
	return response.AbortResult.Body, response.AbortResult.Err
}

// call result's extract function
func ExtractGetIntrospectionDataBaremetalintrospectionV1IntrospectionResponse(response *GetIntrospectionDataBaremetalintrospectionV1IntrospectionResponse) (interface{}, error) {
	return response.DataResult.Body, response.DataResult.Err
}

// call result's extract function
func ExtractReApplyIntrospectionBaremetalintrospectionV1IntrospectionResponse(response *ReApplyIntrospectionBaremetalintrospectionV1IntrospectionResponse) (interface{}, error) {
	return response.ApplyDataResult.Body, response.ApplyDataResult.Err
}