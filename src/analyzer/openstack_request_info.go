package analyzer

import (
	"cloud_manager/src/utils"
	"log"
	"strings"
)

// todo log format
// resource in openstack
// server, img ...
// 一个request file对应多个action
type OpenstackRequestInfo struct {
	ResourcePackageName string //package name
	ResourceName        string //resouce name of the request file
	ResourcePath        string //dir to save the resource Code
	ActionInfos         []*OpenStackActionInfo
	HasResultFile       bool
	RequestImportPaths  utils.Set
	ResultImportPaths   utils.Set
}

func NewOpenstackRequestInfo(requestPackageName string, requestPath string) *OpenstackRequestInfo {
	ri := &OpenstackRequestInfo{
		ResourcePackageName: requestPackageName,
		ResourcePath:        requestPath,
		HasResultFile:       false,
	}
	ri.ResourceName = utils.JoinName(requestPath, "openstack", "")
	ri.ActionInfos = make([]*OpenStackActionInfo, 0)
	ri.RequestImportPaths = utils.NewSet()
	ri.ResultImportPaths = utils.NewSet()
	ri.RequestImportPaths.Insert(requestPath)
	ri.ResultImportPaths.Insert(requestPath)
	return ri
}

// get import path for request file
func (ori *OpenstackRequestInfo) GenRequestImportPaths() {
	ori.RequestImportPaths.Insert(ori.ResourcePath)
	for _, actionInfo := range ori.ActionInfos {
		paramsImportPath := actionInfo.ActionParameters.GetImportPaths()
		paramsImportPath.Delete("github.com/gophercloud/gophercloud")
		ori.RequestImportPaths.Add(paramsImportPath)
		ori.RequestImportPaths.Add(actionInfo.ActionReturns.GetImportPaths())
	}
	ori.RequestImportPaths.Delete("")
}

// get import path for result files
func (ori *OpenstackRequestInfo) GenResultImportPaths() {
	ori.ResultImportPaths.Insert(ori.ResourcePath)
	for _, actionInfo := range ori.ActionInfos {
		if actionInfo.PageExtractInfo != nil {
			ori.ResultImportPaths.Add(actionInfo.PageExtractInfo.ReturnInfo.GetImportPaths())
		}
		if actionInfo.ResultExtractInfo != nil {
			ori.ResultImportPaths.Add(actionInfo.ResultExtractInfo.ReturnInfo.GetImportPaths())
		}
	}
	ori.ResultImportPaths.Delete("")
}

// delete unsupport action
func (ri *OpenstackRequestInfo) checkValidAction(actionInfo *OpenStackActionInfo) bool {
	checkVarInfo := func(varInfos VarInfos) bool {
		for _, varInfo := range varInfos {
			//the action should be exported
			if utils.IsBasicType(varInfo.TypeName) {
				continue
			}
			if utils.IsExportedStruct(varInfo.TypeName) {
				return false
			}
			//todo support action with array parameter
			if strings.HasPrefix(varInfo.TypeName, "[]") {
				return false
			}
		}
		return true
	}
	return checkVarInfo(actionInfo.ActionParameters) && checkVarInfo(actionInfo.ActionReturns)
}

func (ri *OpenstackRequestInfo) AddAction(actionInfo *OpenStackActionInfo) bool {
	if ri.checkValidAction(actionInfo) {
		ri.ActionInfos = append(ri.ActionInfos, actionInfo)
		return true
	} else {
		log.Println("invalid action: ", actionInfo)
	}
	return false
}

//func (ri *OpenstackRequestInfo) AddImportPaths(packagePaths utils.Set) {
//	for packagePath, _ := range packagePaths {
//		if packagePath != "" {
//			ri.RequestImportPaths.Insert(packagePath)
//		}
//	}
//}

type VarInfo struct {
	Name       string
	TypeName   string
	ImportPath string
}

func NewVarInfo(name, typeName, importPath string) VarInfo {
	return VarInfo{
		Name:       name,
		TypeName:   typeName,
		ImportPath: importPath,
	}
}

type VarInfos []VarInfo

func NewVarInfos() VarInfos {
	return make([]VarInfo, 0)
}

func (vi *VarInfos) AddVarInfo(varInfo VarInfo) {
	*vi = append(*vi, varInfo)
}
func (vi *VarInfos) Add(names []string, typeName string, importPath string) {
	for _, name := range names {
		*vi = append(*vi, NewVarInfo(name, typeName, importPath))
	}
}

func (vi *VarInfos) GetImportPaths() utils.Set {
	importPaths := utils.NewSet()
	for _, varinfo := range *vi {
		importPaths.Insert(varinfo.ImportPath)
	}
	return importPaths
}

// describe an action to the resource
// list, get, create ...
type OpenStackActionInfo struct {
	ActionName        string
	ActionParameters  VarInfos           //TypeName, name
	ActionReturns     VarInfos           //TypeName, name
	PageExtractInfo   *PageExtractInfo   //for action start with list
	ResultExtractInfo *ResultExtractInfo // for action return a result type with an extract method
}

func NewOpenstackActionInfo(actionName string) *OpenStackActionInfo {
	ai := &OpenStackActionInfo{
		ActionName:        actionName,
		PageExtractInfo:   nil,
		ResultExtractInfo: nil,
	}
	ai.ActionParameters = NewVarInfos()
	ai.ActionReturns = NewVarInfos()
	return ai
}

func (ai *OpenStackActionInfo) AddVarInfos(varInfos VarInfos, kind string) {
	for _, varInfo := range varInfos {
		if kind == "parameters" {
			if varInfo.TypeName == "*gophercloud.ServiceClient" {
				continue
			}
			ai.ActionParameters.AddVarInfo(varInfo)
		} else {
			ai.ActionReturns.AddVarInfo(varInfo)
		}
	}

}

func GetParamsLsit(paraInfo VarInfos) string {
	var paras = ""
	for i := 0; i < len(paraInfo); i++ {
		name := paraInfo[i].Name
		typeName := paraInfo[i].TypeName
		if typeName == "*gophercloud.ServiceClient" {
			continue
		}
		paras += name + " " + typeName + ","
	}
	if paras == "" {
		return paras
	}
	return paras[:len(paras)-1]
}

func GetParamsCallList(paraInfo VarInfos) string {
	var paras = ""
	for i := 0; i < len(paraInfo); i++ {
		name := paraInfo[i].Name
		typeName := paraInfo[i].TypeName
		if typeName == "*gophercloud.ServiceClient" {
			continue
		}
		paras += name + ","
	}
	if paras == "" {
		return paras
	}
	return paras[:len(paras)-1]
}

func GetReturnsList(returnInfo VarInfos) string {
	var paras = ""
	for i := 0; i < len(returnInfo); i++ {
		name := returnInfo[i].Name
		typeName := returnInfo[i].TypeName
		if name == "" {
			paras += typeName + ","
		} else {
			paras += name + " " + typeName + ","
		}
	}
	if paras == "" {
		return paras
	}
	return paras[:len(paras)-1]
}
