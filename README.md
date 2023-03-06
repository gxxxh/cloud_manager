# The Go Multi Cloud Development API

## What is it?

This golang library aims to provide an abstraction API of different cloud providers, such as Aliyun and Openstack. Instead of writing code to call different API, this package call those API by code analyzing and reflection, which saves a lot of development work.

Different providers have different utilities and API interfaces to manage their cloud infrastructure, but the abstraction of their interfaces are quiet similar. Using the method of code analyzing and code generation. We format those cloud code into the same pattern, which is easy to unstandstand and use.

For users who want to call the sdk, they only need to provide a json file with details about the cloud resource, operation and parameters, and the underlying function will be called using reflection, the parameters can be find from document of this project and corresponding cloud's document. 

## Design
For details of this code, see [design](https://github.com/kube-stack/multicloud_service/blob/main/doc/design.md). 

## Code Genrate Usage

In order to call the cloud sdk api in the same way, we format different cloud code into the same pattern. For example, in openstack, there are cloud resources like server and image, and each resource is related to a lot of operation(create, update, delete...). So we generate a function for each operation separately, the function takes a xxxRequest struct as input
 and a xxxResponse struct as returns. 

```go
// action function
func (oc *OpenstackClient) CreateComputeV2Servers(req *CreateComputeV2ServersRequest)(*CreateComputeV2ServersResponse){
    return NewCreateComputeV2ServersResponse(servers.Create(oc.Client, req.Opts, ))
}
```
We offer a command tool to generate both golang and java cloud sdk. You can build the tool with 
the following command(requires golang > 1.19)
```shell
make cloudcodegen
```
or
```shell
go build -o cloudcodegen ./src/code_generator/main.go
```
Then you can generate code by
```shell
cloudcodegen gen -f configFile.json
```
### Generate Golang SDK Config
#### Aliyun
Here is an example of json config file for aliyun. RegistryConfigs is used to genereate request registry for the 
cloud. The `CodeGenConfig` include the [registry template](https://github.com/kube-stack/multicloud_service/blob/main/src/code_generator/templates/registry.tmpl) path(`CodeGenConfig`) and the directory path(`CodePath`) you 
want to save the registry code generated. 
```json
{
  "CloudType": "Aliyun",
  "RegistryConfigs": [
    {
      "CodeGenConfig": {
        "TemplatePath":"",
        "CodePath":""
      }
    }
  ]
}
```

#### Openstack Config
1. For Openstack, there are two kind of registry(request and response),
2. We need to generate new openstack sdk from [gophercloud](https://github.com/gophercloud/gophercloud), so the user must provide APICodeConfigs
   1. The first one is used to generate `Request` code, and the second one is used to generate `Result` code. User only need to change the file path and source code path in this config.
   2. `SourceCodePath` is where you save the gophercloud code in your local file system.
   2. Note that the first TemplatePath of the APICodeConfigs should be [openstack_request.tmpl](https://github.com/kube-stack/multicloud_service/blob/main/src/code_generator/templates/openstack_request.tmpl), and the second should be [openstack_result.tmpl](https://github.com/kube-stack/multicloud_service/blob/main/src/code_generator/templates/openstack_result.tmpl). Similarly, the 
      CodeType should be `Request` and `Result`.
```json
{
  "CloudType": "Openstack",
  "RegistryConfigs": [
    {
      "CodeGenConfig": {
        "TemplatePath": "",
        "CodePath": ""
      }
    }
  ],
  "APICodeConfigs": [
    {
      "CodeGenConfig": {
        "TemplatePath": "",
        "CodePath": ""
      },
      "SourceCodePath": "",
      "CodeType": ""
    },
    {
      "CodeGenConfig": {
        "TemplatePath": "",
        "CodePath": ""
      },
      "SourceCodePath": "",
      "CodeType": ""
    }
  ]
}
```
### Generate JAVA SDK
#### Config
We Provide a java-sdk project, which will use this package indirectly with k8s. The code in that package is also 
generate by this project. The config file is listed as below. User can get the template file from [java templates](https://github.com/kube-stack/multicloud_service/tree/main/src/code_generator/templates), and the code path need to 
be [kubestack](https://github.com/kube-stack/java-sdk/tree/master/src/main/java/io/github/kubestack) ot the 
[java-sdk](https://github.com/kube-stack/java-sdk) project

```json
  "JavaCodeConfig": {
    "Class": {
      "TemplatePath": "",
      "CodePath": ""
    },
    "Domain": {
      "TemplatePath": "",
      "CodePath": ""
    },
    "Impl": {
      "TemplatePath": "",
      "CodePath": ""
    },
    "LifecycleHeader": {
      "TemplatePath": "",
      "CodePath": ""
    },
    "LifecycleClass": {
      "TemplatePath": "",
      "CodePath": ""
    },
    "Spec": {
      "TemplatePath": "",
      "CodePath": ""
    }
  }
```

## Cloud Service Usage
This abstraction of API only accept json format of input. The support action and corresponding prarmeters can be find from the doc of this repository. The detail meaning and correct values of those parameters can be find from those cloud providers' document website.

### Openstack

```go
func main(){
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
}
```

### Aliyun

This example call Aliyun DescribeInstances API.

```go
func main(){
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
}
```

## Supported Providers
| cloud                 | version                                                                           | document       |
|-----------------------|-----------------------------------------------------------------------------------|----------------|
| Aliyun                | [alibaba-cloud-sdk-go(\>=1.62.0)](https://github.com/aliyun/alibaba-cloud-sdk-go) | [Aliyun ECS](https://help.aliyun.com/document_detail/25485.html) |
| Openstack  | [gophercloud](https://github.com/gophercloud/gophercloud)                                                                   |https://docs.openstack.org/zed/api/|

## 下一步计划
1. support other cloud services