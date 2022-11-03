package analyzer

import "cloud_manager/src/utils"

// info using to generate extract function

//1. Action return page should call extract, hard

//todo analyzer results.to
//1. CheckFunction: function, not method; start with extract; parameter type is pagination.Pate

// 2. struct, endwith result, has method extract
// 2. GenDecl, specs is typespce, specs name is endwith result. get the type and methods
// todo 抽象出统一的接口
type OpenstackResultInfo struct {
	PackageName      string
	ResourceName     string
	ResourcePath     string
	ImportPaths      utils.Set
	pageExtractInfos []*PageExtractInfo
}

func NewOpenstackResultInfo(packageName, packagePath string) *OpenstackResultInfo {
	ori := &OpenstackResultInfo{
		PackageName:  packageName,
		ResourcePath: packagePath,
	}
	ori.ResourceName = utils.JoinName(packagePath, "openstack", "")
	ori.ImportPaths = utils.NewSet()
	ori.pageExtractInfos = make([]*PageExtractInfo, 0)
	return ori
}
func (ori *OpenstackResultInfo) AddImportPaths(packagePaths utils.Set) {
	for packagePath, _ := range packagePaths {
		if packagePath != "" {
			ori.ImportPaths.Insert(packagePath)
		}
	}
}
func (ori *OpenstackResultInfo) AddPageExtractInfos(pageExtractInfo *PageExtractInfo) {
	ori.pageExtractInfos = append(ori.pageExtractInfos, pageExtractInfo)
}

type PageExtractInfo struct {
	FuncName   string
	ReturnInfo []VarInfo
}

func NewPageExtractInfo(funcName string) *PageExtractInfo {
	pei := &PageExtractInfo{
		FuncName:   funcName,
		ReturnInfo: nil,
	}
	pei.ReturnInfo = make([]VarInfo, 0)
	return pei
}
