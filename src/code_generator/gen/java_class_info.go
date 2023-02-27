package gen

import (
	"fmt"
	"log"
	"strings"
)

const (
	JSON_INCLOUDE_ANNOTATION          = "@JsonInclude(JsonInclude.Include.NON_NULL)"
	JSON_DESERIALIZE_ANNOTATION       = "@JsonDeserialize(using = com.fasterxml.jackson.databind.JsonDeserializer.None.class)"
	JSON_IGNORE_Properties_ANNOTATION = "@JsonIgnoreProperties(ignoreUnknown = true)"
	JSON_ROOT_NAME_ANNOTATION         = "@JsonRootName(\"%s\")"
	JSON_PROPERTY_ANNOTATION          = "@JsonProperty(\"%s\")"
)

// 用于描述一个java资源，比如OpenstackServer
type JavaResource struct {
	Name         string //crd name, eg: OpenstackServer
	ResourceName string //resource name in code gen, eg: ComputeV2Servers
	JavaClass
}

func NewJavaResource(name, resourceName string) *JavaResource {
	jr := &JavaResource{
		Name:         name,
		ResourceName: resourceName,
	}
	return jr
}

type JavaClass struct {
	ClassName       string
	Depth           int //used to generated \t in template
	Annotations     []string
	MemberVariables []*MemberVariable
	MemberClasses   []*JavaClass
}

func NewJavaClass(name, jsonRootName string, depth int) *JavaClass {
	jc := &JavaClass{
		ClassName:       name,
		Depth:           depth,
		Annotations:     make([]string, 0, 0),
		MemberVariables: make([]*MemberVariable, 0, 0),
		MemberClasses:   make([]*JavaClass, 0, 0),
	}
	jc.Annotations = append(jc.Annotations, JSON_INCLOUDE_ANNOTATION)
	jc.Annotations = append(jc.Annotations, JSON_DESERIALIZE_ANNOTATION)
	jc.Annotations = append(jc.Annotations, JSON_IGNORE_Properties_ANNOTATION)
	if jsonRootName != "" {
		jc.Annotations = append(jc.Annotations, fmt.Sprintf(JSON_ROOT_NAME_ANNOTATION, jsonRootName))
	} else {
		jc.Annotations = append(jc.Annotations, fmt.Sprintf(JSON_ROOT_NAME_ANNOTATION, name))
	}
	return jc
}

type MemberKind uint

const (
	Basic MemberKind = iota //基础类型和自定义类型，不需要翻译
	Array
	Map
	Time
)

// 若为map,保存value的类型，默认key为string
type MemberVariable struct {
	Name       string
	TypeName   string
	JsonName   string //@JsonProperty
	MemberKind MemberKind
}

func NewMemberVariable(name, typename, jsonName string, memberKind MemberKind) *MemberVariable {
	if jsonName == "" || jsonName == "-" {
		jsonName = name
	}
	//"tag, omitempty" → "tag"
	if strings.Contains(jsonName, ",") {
		jsonName = strings.Split(jsonName, ",")[0]
	}
	mv := &MemberVariable{
		Name:       name,
		TypeName:   typename,
		JsonName:   jsonName,
		MemberKind: memberKind,
	}
	return mv
}

func (jc *JavaClass) Add(mv *MemberVariable, mc *JavaClass) {
	if mv.TypeName == "" {
		log.Printf("member %v type is interface, deleted\n", mv.Name)
		mv.TypeName = "Object"
		//return
	}
	if mv != nil {
		jc.MemberVariables = append(jc.MemberVariables, mv)
	}
	if mc != nil {
		jc.MemberClasses = append(jc.MemberClasses, mc)
	}
	return
}
func (jc *JavaClass) AddMemberVariabel(mv *MemberVariable) {

	return
}
func (jc *JavaClass) AddMemberClass(memberClass *JavaClass) {

	return
}
