package openstack

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

type OpenstackClient struct {
	ProviderClient *gophercloud.ProviderClient
	Client         *gophercloud.ServiceClient
	Kind           string //compute...
}

func NewOpenstackClient(params map[string]string) (oc *OpenstackClient, err error) {
	oc = &OpenstackClient{
		ProviderClient: nil,
		Client:         nil,
		Kind:           "",
	}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: params["IdentityEndpoint"],
		Username:         params["username"],
		Password:         params["password"],
	}
	oc.ProviderClient, err = openstack.AuthenticatedClient(opts)
	return
}

func (oc *OpenstackClient) InitClient(kind string, eo gophercloud.EndpointOpts) (err error) {

	switch kind {
	case "baremetal":
		oc.Client, err = openstack.NewBareMetalV1(oc.ProviderClient, eo)
	case "baremetal-inspector":
		oc.Client, err = openstack.NewBareMetalIntrospectionV1(oc.ProviderClient, eo)
	case "object-store":
		oc.Client, err = openstack.NewObjectStorageV1(oc.ProviderClient, eo)
	case "compute":
		oc.Client, err = openstack.NewComputeV2(oc.ProviderClient, eo)
	case "network":
		oc.Client, err = openstack.NewNetworkV2(oc.ProviderClient, eo)
	case "volume":
		oc.Client, err = openstack.NewBlockStorageV1(oc.ProviderClient, eo)
	case "volumev2":
		oc.Client, err = openstack.NewBlockStorageV2(oc.ProviderClient, eo)
	case "volumev3":
		oc.Client, err = openstack.NewBlockStorageV3(oc.ProviderClient, eo)
	case "sharev2":
		oc.Client, err = openstack.NewSharedFileSystemV2(oc.ProviderClient, eo)
	case "cdn":
		oc.Client, err = openstack.NewCDNV1(oc.ProviderClient, eo)
	case "orchestration":
		oc.Client, err = openstack.NewOrchestrationV1(oc.ProviderClient, eo)
	case "database":
		oc.Client, err = openstack.NewDBV1(oc.ProviderClient, eo)
	case "dns":
		oc.Client, err = openstack.NewDNSV2(oc.ProviderClient, eo)
	case "image":
		oc.Client, err = openstack.NewImageServiceV2(oc.ProviderClient, eo)
	case "load-balancer":
		oc.Client, err = openstack.NewLoadBalancerV2(oc.ProviderClient, eo)
	case "clustering":
		oc.Client, err = openstack.NewClusteringV1(oc.ProviderClient, eo)
		//todo need Client id
	//case "messaging":
	//	oc.Client, err = openstack.NewMessagingV2(Client, eo)
	case "container":
		oc.Client, err = openstack.NewContainerV1(oc.ProviderClient, eo)
	case "key-manager":
		oc.Client, err = openstack.NewKeyManagerV1(oc.ProviderClient, eo)
	case "container-infra":
		oc.Client, err = openstack.NewContainerInfraV1(oc.ProviderClient, eo)
	case "workflowv2":
		oc.Client, err = openstack.NewWorkflowV2(oc.ProviderClient, eo)
	case "placement":
		oc.Client, err = openstack.NewPlacementV1(oc.ProviderClient, eo)
	case "identityv2":
		oc.Client, err = openstack.NewIdentityV2(oc.ProviderClient, eo)
	case "identityv3":
		oc.Client, err = openstack.NewIdentityV3(oc.ProviderClient, eo)
	default:
		err = fmt.Errorf("no such kind in openstack")
	}
	return
}
