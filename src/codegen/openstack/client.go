package openstack

import "github.com/gophercloud/gophercloud"

type OpenstackClient struct {
	*gophercloud.ServiceClient
	Kind string //compute...
}
