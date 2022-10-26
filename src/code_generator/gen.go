package code_generator

import (
	"bytes"
	"cloud_manager/src/analyzer"
	"log"
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
		"GenParams":     analyzer.GetParas,
		"GenReturns":    analyzer.GetReturns,
		"GenParamsCall": analyzer.GetParasCall,
		"now": func() string {
			return time.Now().Format(time.RFC3339)
		},
		"param": func(name string) interface{} {
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

func GenCode(templatePath string, data map[string]interface{}, packageName string) ([]byte, error) {
	createRequestRegistryTemplate, err := NewCustomerTemplate(templatePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	code, err := GenerateTemplate(createRequestRegistryTemplate.GetTemplateBody(),
		data,
		map[string]interface{}{
			"packageName": packageName,
		})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return code, err
}
