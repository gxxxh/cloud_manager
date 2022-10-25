package code_generator

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
)

type OpenstackClient struct {
	*gophercloud.ServiceClient
	Kind string //compute...
}

func (oc *OpenstackClient) ListServes(opts servers.ListOptsBuilder) pagination.Pager {
	return servers.List(oc.ServiceClient, opts)
}
