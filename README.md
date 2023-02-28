# The Go Multi Cloud Development API

## What is it?

This is a golang library to provide an abstraction API of different
cloud providers, such as Aliyun and Openstack. Instead of
writing code to call different API, this package call those API by code analyzing
and reflection, which saves a lot of development work.

Different providers have different utilities and API interfaces to manage their cloud infrastructure, but the
abstraction
of their interfaces are quiet similar. Using the method of code analyzing and code generation. We format those cloud
into
the following pattern. The request struct contains the parameters of the cloud API, and the returns of those API will be
written into the response struct.

```go
// action function
func (oc *OpenstackClient) CreateComputeV2Servers(req *CreateComputeV2ServersRequest)(*CreateComputeV2ServersResponse){
    return NewCreateComputeV2ServersResponse(servers.Create(oc.Client, req.Opts, ))
}
```

For users of this pacakge, the only work they need to do is writing a right json of the underlying API, which can be
find from the document of those cloud SDKs.

## Code Generate
### Golang SDK

### Java SDK

## Usage
This abstraction of API only accept json format of input. The support action and corresponding prarmeters can be find
from the doc of this repository. The detail meaning and correct values of those parameters can be find from those cloud
providers' document website.

### Openstack

```go
//openstack glance get image
//openstack authentication info
authInfo := map[string]string{
"projectName":         "",
"domainName":          "",
"identityEndpoint":    "",
"username":            "",
"password":            "",
"Region":              "",
"openstackClientType": "image",
"cloudType":           "openstack",
}
service, _ := service.NewMultiCloudService(authInfo)
requestByte := `{"Id":"e7db3b45-8db7-47ad-8109-3fb55c2c24fe"}`
resp, err := service.CallCloudAPI("GetImageserviceV2Images", requestByte)
if err != nil {
    t.Error(err)
}
```

### Aliyun

This example call Aliyun DescribeInstances API.

```go
//aliyun authentication info 
authInfo := map[string]string{
"regionId":"",
"accessId":"",
"accessKeySecret":"",
"cloudType":"aliyun"
}
mcm, _ := service.NewMultiCloudService(params)
//request json
requestJson := `{
        "InstanceIds":"[\"i-2zegiq87g0txkt1bvrb5\"]",
    }`
//call underlying cloud API    `	
ret, err := mcm.CallCloudAPI("DescribeInstances", requestJson)
if err != nil {
    t.Error(err)
}
```

## Supported Providers
| cloud                 | version                                                                           | document       |
|-----------------------|-----------------------------------------------------------------------------------|----------------|
| Aliyun                | [alibaba-cloud-sdk-go(\>=1.62.0)](https://github.com/aliyun/alibaba-cloud-sdk-go) | [Aliyun ECS](https://help.aliyun.com/document_detail/25485.html) |
| Openstack  | [gophercloud](https://github.com/gophercloud/gophercloud)                                                                   |https://docs.openstack.org/zed/api/|

## 下一步计划
1. 对compute/extentsion中不包含list/get/create等方法的，归类到server方法中
2. 提供cli代码生成接口
3. 代码重构，注释添加