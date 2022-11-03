package utils

import "strings"

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
