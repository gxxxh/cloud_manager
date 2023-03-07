# Names in Openstack
In this project, in order to generate openstack code, there are many variables named 'XXName', which is really confused. So Here gives a example to explain it. 

Take openstack server as example, the source code can be find from [here](https://github.com/gophercloud/gophercloud/blob/master/openstack/compute/v2/servers/requests.go)

#### PackageName:
It means the package name defined in the source code, which is `server`

#### ResourceName
The Resource name is `OpenstackComputeV2Servers`. Cause some of the packages of gophercloud has the same name, we add the file path(`ComputeV2`) to the resource name, and `Openstack` is the cloud type.