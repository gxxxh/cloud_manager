package analyzer

import (
	"cloud_manager/src/utils"
	"strings"
)

// resource in openstack
// server, img ...
type OpenstackResourceInfo struct {
	ResourcePackageName string
	ResourceName        string
	ResourcePath        string //dir to save the resource Code
	Actions             []*OpenStackActionInfo
	ImportPaths         utils.Set
}

func NewOpenstackResourceInfo(resourcePackageName string, resourcePath string) *OpenstackResourceInfo {
	ri := &OpenstackResourceInfo{
		ResourcePackageName: resourcePackageName,
		ResourcePath:        resourcePath,
	}
	ri.ResourceName = utils.JoinName(resourcePath, "openstack", "")
	ri.Actions = make([]*OpenStackActionInfo, 0)
	ri.ImportPaths = utils.NewSet()
	ri.ImportPaths.Insert(resourcePath)
	return ri
}

func (ri *OpenstackResourceInfo) AddAction(actionInfo *OpenStackActionInfo) {
	ri.Actions = append(ri.Actions, actionInfo)
}

type VarInfo struct {
	Name     string
	TypeName string
}

// describe a action to the resouce
// list, get, create ...
type OpenStackActionInfo struct {
	ActionName       string
	ActionParameters []VarInfo //TypeName, name
	ActionReturns    []VarInfo //TypeName, name
}

func NewOpenstackActionInfo(actionName string) *OpenStackActionInfo {
	ai := &OpenStackActionInfo{
		ActionName: actionName,
	}
	ai.ActionParameters = make([]VarInfo, 0)
	ai.ActionReturns = make([]VarInfo, 0)
	return ai
}

/*
add parameters/return variable name and typeName
*/
func (ai *OpenStackActionInfo) AddVarInfo(name, typeName, kind string) {
	varInfo := VarInfo{
		Name:     name,
		TypeName: typeName,
	}
	if kind == "parameters" {
		ai.ActionParameters = append(ai.ActionParameters, varInfo)
	} else {
		ai.ActionReturns = append(ai.ActionReturns, varInfo)
	}
}

func GetParasList(paraInfo []VarInfo) string {
	var paras = ""
	for i := 0; i < len(paraInfo); i++ {
		name := paraInfo[i].Name
		typeName := paraInfo[i].TypeName
		paras += name + " " + typeName + ","
	}
	return paras[:len(paras)-1]
}

func GetParasCallList(paraInfo []VarInfo) string {
	var paras = ""
	for i := 0; i < len(paraInfo); i++ {
		name := paraInfo[i].Name
		paras += name + ","
	}
	return paras[:len(paras)-1]
}

func GetReturnsList(returnInfo []VarInfo) string {
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
	return paras[:len(paras)-1]
}

func TypeName2MemberName(typeName string) string {
	return utils.UpperFirst(TypeName2LocalVarName(typeName))
}

func TypeName2LocalVarName(typeName string) string {
	//todo check basic type
	localVarName := typeName
	if strings.Contains(typeName, ".") {
		tmp := strings.Split(typeName, ".")
		localVarName = tmp[len(tmp)-1]
	}
	return utils.LowerFirst(localVarName)
}
