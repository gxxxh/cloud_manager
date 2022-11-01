package code_generator

import (
	"bytes"
	"cloud_manager/src/analyzer"
	"cloud_manager/src/log"
	"cloud_manager/src/utils"
	"strings"
	"text/template"
	"time"
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
		"GenParams":     analyzer.GetParasList,
		"GenReturns":    analyzer.GetReturnsList,
		"GenParamsCall": analyzer.GetParasCallList,
		"now": func() string {
			return time.Now().Format(time.RFC3339)
		},
		"GenMemberName":   utils.TypeName2MemberName,   // 大写首字母作为成员变量
		"GenLocalVarName": utils.TypeName2LocalVarName, // 针对返回值生成针对类型的成员变量名称
		"UpperFirst":      utils.UpperFirst,
		"param": func(name string) interface{} {
			if v, ok := params[name]; ok {
				return v
			}
			return ""
		},
	}).Parse(templateText)
	if err != nil {
		log.ERROR("%v\n ", err)
		return nil, err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, templateData); err != nil {
		log.ERROR("%v\n", err)
		return nil, err
	}

	return buf.Bytes(), nil
}

func GenCode(templatePath string, data map[string]interface{}, packageName string) ([]byte, error) {
	createRequestRegistryTemplate, err := NewCustomerTemplate(templatePath)
	if err != nil {
		log.ERROR("%v\n", err)
		return nil, err
	}
	code, err := GenerateTemplate(createRequestRegistryTemplate.GetTemplateBody(),
		data,
		map[string]interface{}{
			"packageName": packageName,
		})
	if err != nil {
		log.ERROR("%v\n", err)
		return nil, err
	}
	return code, err
}
