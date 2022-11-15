package gen

import (
	"bytes"
	"github.com/kube-stack/multicloud_service/src/analyzer"
	"github.com/kube-stack/multicloud_service/src/utils"
	"log"
	"strings"
	"text/template"
)

/*
templateData is data used for template
params are extra info used by param function
*/
func GenerateTemplate(templateText string, templateData interface{}, params map[string]interface{}) ([]byte, error) {
	t, err := template.New("tableTemplate").Funcs(template.FuncMap{
		"Replace": func(old, new, src string) string {
			return strings.ReplaceAll(src, old, new)
		},
		"GetCodeHeader": func() string {
			return templateHeader
		},
		"GetParamsList":     analyzer.GetParamsLsit,
		"GetReturnsList":    analyzer.GetReturnsList,
		"GetParamsCallList": analyzer.GetParamsCallList,
		"GenMemberName":     utils.TypeName2MemberName,   // 大写首字母作为成员变量
		"GenLocalVarName":   utils.TypeName2LocalVarName, // 针对返回值生成针对类型的成员变量名称
		"UpperFirst":        utils.UpperFirst,
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
