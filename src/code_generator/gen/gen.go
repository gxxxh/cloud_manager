package gen

import (
	"bytes"
	"fmt"
	"github.com/kube-stack/multicloud_service/src/analyzer"
	"github.com/kube-stack/multicloud_service/src/utils"
	"log"
	"strings"
	"text/template"
	"time"
)

func GenAndSaveCode(templatePath, codePath string, data, params map[string]interface{}) {
	code, err := GenCode(templatePath, data, params)
	if err != nil {
		log.Fatalln("Gen Code error, ", err)
	}
	utils.Save(code, codePath)
}

func GenCode(templatePath string, data map[string]interface{}, params map[string]interface{}) ([]byte, error) {
	createRequestRegistryTemplate, err := NewCustomerTemplate(templatePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	code, err := GenerateTemplate(createRequestRegistryTemplate.GetTemplateBody(),
		data, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return code, err
}

/*
templateData is data used for template
params are extra info used by param function
*/
func GenerateTemplate(templateText string, templateData interface{}, params map[string]interface{}) ([]byte, error) {
	// functions used in template
	t, err := template.New("tableTemplate").Funcs(template.FuncMap{
		"GetParamsList":     analyzer.GetParamsLsit,
		"GetReturnsList":    analyzer.GetReturnsList,
		"GetParamsCallList": analyzer.GetParamsCallList,
		"GenMemberName":     utils.TypeName2MemberName,   // 大写首字母作为成员变量
		"GenLocalVarName":   utils.TypeName2LocalVarName, // 针对返回值生成针对类型的成员变量名称
		"UpperFirst":        utils.UpperFirst,
		"LowerFirst":        utils.LowerFirst,
		"ToLower":           strings.ToLower,
		"GetPreDepth": func(depth int) int {
			return depth - 1
		},
		"GetTab": func(depth int) string {
			res := ""
			for i := 0; i < depth; i++ {
				res += "    "
			}
			return res
		},
		"GenJavaType": func(typeName string, memberKind MemberKind) string {
			javaTypeName := utils.TypeConvert(typeName)
			switch memberKind {
			case Array:
				return fmt.Sprintf("List<%s>", javaTypeName)
			case Map:
				return fmt.Sprintf("HashMap<String, %s>", javaTypeName)
			default:
				return javaTypeName
			}
		},
		"Date": func() string {
			now := time.Now()
			return fmt.Sprintf("%v/%v/%v %d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
		},
		"Replace": func(old, new, src string) string {
			return strings.ReplaceAll(src, old, new)
		},
		"GetCodeHeader": func() string {
			return templateHeader
		},
		"RemoveRequestSuffix": func(key string) string {
			return key[0 : len(key)-len("Request")]
		},
		"Param": func(name string) interface{} {
			if v, ok := params[name]; ok {
				return v
			}
			return ""
		},
	}).Parse(templateText)
	if err != nil {
		log.Println("GenTemplateError: ", err)
		return nil, err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, templateData); err != nil {
		log.Println("GenTemplateError: ", err)
		return nil, err
	}

	return buf.Bytes(), nil
}
