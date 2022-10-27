package openstack

import "github.com/gophercloud/gophercloud"

type OpenstackClient struct {
	client *gophercloud.ServiceClient
	Kind   string //compute...
}
