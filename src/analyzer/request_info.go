package analyzer

import (
	"cloud_manager/src/utils"
	"reflect"
)

var basePath = "test"

// 用于生成函数registry
type RequestInfo struct {
	RequestName        string
	CreateFunctionName string
}

func NewRequestInfo(requestTypeName string, requestType reflect.Type) *RequestInfo {
	packageName := requestType.String()
	return &RequestInfo{
		RequestName:        requestTypeName,
		CreateFunctionName: utils.GetPackageName(packageName) + ".Create" + requestTypeName,
	}
}

type RequestRegistryInfo struct {
	RequestInfos []*RequestInfo
	ImportPaths  []string
}

func NewRequestRegistryInfo() *RequestRegistryInfo {
	return &RequestRegistryInfo{
		RequestInfos: make([]*RequestInfo, 0),
		ImportPaths:  make([]string, 0),
	}
}
