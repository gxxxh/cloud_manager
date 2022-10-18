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
