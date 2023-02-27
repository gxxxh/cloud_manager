package utils

import (
	"strings"
	"unicode"
)

func UpperFirst(s string) string {
	first := strings.ToUpper(s[0:1])
	if len(s) == 1 {
		return first
	}
	return first + s[1:]
}

func LowerFirst(s string) string {
	first := strings.ToLower(s[0:1])
	if len(s) == 1 {
		return first
	}
	return first + s[1:]
}

// 删除函数名称的actionName(list, Create)等，若没有宾语则返回原名称
func ParseResourceName(funcName string, actionName string) string {
	if strings.HasPrefix(funcName, actionName) && funcName != actionName {
		return strings.ToLower(funcName[len(actionName):])
	}
	return strings.ToLower(funcName)
}

// 删除复数的s，java中类名称，也是crd的名称
func GetJavaResourceName(cloudType, resourceName string) string {
	return UpperFirst(cloudType) + UpperFirst(resourceName)
}

// 从request名称中获取资源名称
// requestName中第一个单词表示动作，删除即可
func GetResourceNameFromRequestName(requestName string) string {
	for i := 1; i < len(requestName); i++ {
		if unicode.IsUpper(rune(requestName[i])) {
			return requestName[i:]
		}
	}
	return ""
}

// CreateComputeV2ServersRequest
// 返回actionName="Create", resourceName = "server", cloudResourceName="ComputeV2Servers"
func ParseRequestName(name string) (actionName, cloudResourceName, resourceName string) {
	if strings.HasSuffix(name, "Request") {
		name = name[0 : len(name)-len("Request")]
	}
	pathPrefixes := [...]string{"Baremetal", "Baremetalintrospection", "Blockstorage", "Cdn", "Clustering", "Common", "Compute", "Container", "Containerinfra", "Db", "Dns", "Identity", "Imageservice", "Keymanager", "Loadbalancer", "Messaging", "Networking", "Objectstorage", "Orchestration", "Placement", "Sharedfilesystems", "Testing", "Utils", "Workflow"}
	for _, pathPrefix := range pathPrefixes {
		if pos := strings.Index(name, pathPrefix); pos != -1 {
			actionName = name[:pos]
			cloudResourceName = name[pos:]
			break
		}
	}
	j := 0
	for i := 0; i < len(name); i++ {
		if unicode.IsUpper(rune(name[i])) {
			j = i
		}
	}
	resourceName = strings.ToLower(name[j : len(name)-1])
	return
}

// 获取ComputeV2Servers中的Servers，即最后一个单词
func GetResourcePackageName(resourceName string) string {
	j := 0
	for i := 0; i < len(resourceName); i++ {
		if unicode.IsUpper(rune(resourceName[i])) {
			j = i
		}
	}
	return strings.ToLower(resourceName[j:])
}

// 获取A.B中的A
func GetPackageName(name string) string {
	packageName := ""
	for _, ch := range name {
		if ch == '.' {
			break
		}
		packageName += string(ch)
	}
	return packageName
}

/*
check if the first character is lower
*/
func IsLower(s string) bool {
	return s[0] >= 'a' && s[0] <= 'z'
}
func IsBasicType(typeName string) bool {
	tmp := typeName
	if strings.HasPrefix(typeName, "[]") {
		tmp = tmp[2:]
	}
	if strings.HasPrefix(typeName, "*") {
		tmp = tmp[1:]
	}
	if tmp == "string" ||
		tmp == "int" ||
		tmp == "bool" ||
		tmp == "byte" ||
		tmp == "chan" ||
		tmp == "error" {
		return true
	}
	return false
}

func IsExportedStruct(typeName string) bool {
	names := strings.Split(typeName, ".")
	structName := names[len(names)-1]
	return IsLower(structName)
}

func TypeName2MemberName(typeName string) string {
	return UpperFirst(TypeName2LocalVarName(typeName))
}

// remove package info from typename
func GetStructName(typeName string) string {
	if strings.Contains(typeName, ".") {
		tmp := strings.Split(typeName, ".")
		return tmp[len(tmp)-1]
	}
	return typeName
}
func TypeName2LocalVarName(typeName string) string {
	//todo check basic type
	localVarName := GetStructName(typeName)
	return LowerFirst(localVarName)
}

func CompareSlice(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	tmp := NewSet()
	for _, s := range s1 {
		tmp.Insert(s)
	}
	for _, s := range s2 {
		if tmp.Has(s) {
			tmp.Delete(s)
		} else {
			return false
		}
	}
	return len(tmp) == 0
}

func DiffSlice(s1, s2 []string) Set {
	res := NewSet()
	for _, s := range s1 {
		res.Insert(s)
	}
	for _, s := range s2 {
		if res.Has(s) {
			res.Delete(s)
		} else {
			res.Insert(s)
		}
	}
	return res
}
