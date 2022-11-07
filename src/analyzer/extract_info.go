package analyzer

import (
	"cloud_manager/src/utils"
)

// info using to generate extract function
//1. CheckFunction: function, not method; start with extract; parameter type is pagination.Pate

type OpenstackResultInfo struct {
	PackageName      string
	ResourceName     string
	ResourcePath     string
	PageExtractInfos []*PageExtractInfo
}

func NewOpenstackResultInfo(packageName, packagePath string) *OpenstackResultInfo {
	ori := &OpenstackResultInfo{
		PackageName:  packageName,
		ResourcePath: packagePath,
	}
	ori.ResourceName = utils.JoinName(packagePath, "openstack", "")
	ori.PageExtractInfos = make([]*PageExtractInfo, 0)
	return ori
}

//	func (ori *OpenstackResultInfo) AddImportPaths(packagePaths utils.Set) {
//		for packagePath, _ := range packagePaths {
//			if packagePath != "" {
//				ori.ImportPaths.Insert(packagePath)
//			}
//		}
//	}
func (ori *OpenstackResultInfo) AddPageExtractInfos(pageExtractInfo *PageExtractInfo) {
	ori.PageExtractInfos = append(ori.PageExtractInfos, pageExtractInfo)
}

// info using to describe page extract function
type PageExtractInfo struct {
	FuncName   string
	ReturnInfo VarInfos
}

func NewPageExtractInfo(funcName string) *PageExtractInfo {
	pei := &PageExtractInfo{
		FuncName:   funcName,
		ReturnInfo: nil,
	}
	pei.ReturnInfo = NewVarInfos()
	return pei
}

//info using to describe result.extract function
// 1. 类型以Result结尾
// 2. 类型存在Extract()函数
// 3. 分析 Extract()函数，进行封装

type ResultExtractInfo struct {
	FuncName   string // Extract or ExtractErr
	ReturnInfo VarInfos
}

func NewResultExtractInfo() *ResultExtractInfo {
	rei := &ResultExtractInfo{}
	return rei
}
