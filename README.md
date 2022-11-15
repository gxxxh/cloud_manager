
## todo 
### Operator
- 资源定义
  - [ ] AliyunECS
  - [ ] Openstack


### Aliyun
1. 创建项目:
```bash
kubebuilder init --domain doslab.io
```
2. 创建API
```bash
kubebuilder create api --group aliyun.ecs --kind VMInstance --version v1
```
### Example
#### AliyunECS
1. 认证
```json
{
  "apiVersion": "cloudplus.io/v1alpha3",
  "kind": "AliyunECS",
  "metadata": {
    "name": ""
  },
  "spec": {
    "lifecycle":{
      "commandName":{
        "parameterName": "parameterValue"
      }
    }
  }
}
```