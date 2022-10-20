package code_generator

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"
)

func GenerateTemplate(templateText string, templateData interface{}, params map[string]interface{}) ([]byte, error) {
	t, err := template.New("tableTemplate").Funcs(template.FuncMap{
		"Replace": func(old, new, src string) string {
			return strings.ReplaceAll(src, old, new)
		},
		"Add": func(a, b int) int {
			return a + b
		},
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
		fmt.Println("GenTemplateError: ", err)
		return nil, err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, templateData); err != nil {
		fmt.Println("GenTemplateError: ", err)
		return nil, err
	}

	return buf.Bytes(), nil
}
