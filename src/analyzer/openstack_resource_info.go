package analyzer

import (
	"cloud_manager/src/utils"
	"strings"
)

// resource in openstack
// server, img ...
type OpenstackResourceInfo struct {
	ResourceName        string //packageName is the resource Name
	ResourcePackageName string
	ResourcePath        string //dir to save the resource Code
	Actions             []*OpenStackActionInfo
	ImportPaths         utils.Set
}

func NewOpenstackResourceInfo(resourcePackageName string, resourcePath string) *OpenstackResourceInfo {
	resourceName := strings.ToUpper(resourcePackageName[0:1])
	ri := &OpenstackResourceInfo{
		ResourceName:        resourceName + resourcePackageName[1:],
		ResourcePackageName: resourcePackageName,
		ResourcePath:        resourcePath,
	}
	ri.Actions = make([]*OpenStackActionInfo, 0)
	ri.ImportPaths = utils.NewSet()
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

func GetParas(actionInfo *OpenStackActionInfo) string {
	var paras = ""
	for i := 0; i < len(actionInfo.ActionParameters); i++ {
		name := actionInfo.ActionParameters[i].Name
		typeName := actionInfo.ActionParameters[i].TypeName
		paras += name + " " + typeName + ","
	}
	return paras[:len(paras)-1]
}

func GetParasCall(actionInfo *OpenStackActionInfo) string {
	var paras = ""
	for i := 0; i < len(actionInfo.ActionParameters); i++ {
		name := actionInfo.ActionParameters[i].Name
		paras += name + ","
	}
	return paras[:len(paras)-1]
}

func GetReturns(actionInfo *OpenStackActionInfo) string {
	var paras = ""
	for i := 0; i < len(actionInfo.ActionReturns); i++ {
		name := actionInfo.ActionReturns[i].Name
		typeName := actionInfo.ActionReturns[i].TypeName
		if name == "" {
			paras += typeName + ","
		} else {
			paras += name + " " + typeName + ","
		}
	}
	return paras[:len(paras)-1]
}
