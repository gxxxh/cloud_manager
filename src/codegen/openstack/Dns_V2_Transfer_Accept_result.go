package openstack

// Code generated by cloud manager.

import (
	"github.com/gophercloud/gophercloud/openstack/dns/v2/transfer/accept"
)

// extract response info from pager for ListDnsV2TransferAccept
func ExtractListDnsV2TransferAcceptResponse(response *ListDnsV2TransferAcceptResponse) ([]accept.TransferAccept, error) {
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return accept.ExtractTransferAccepts(page)
}

// call result's extract function
func ExtractGetDnsV2TransferAcceptResponse(response *GetDnsV2TransferAcceptResponse) (interface{}, error) {
	return response.GetResult.Body, response.GetResult.Err
}

// call result's extract function
func ExtractCreateDnsV2TransferAcceptResponse(response *CreateDnsV2TransferAcceptResponse) (interface{}, error) {
	return response.CreateResult.Body, response.CreateResult.Err
}