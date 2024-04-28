package gen

import "os"

var templateHeader = `
// Code generated by cloud manager.
`

type ITemplate interface {
	GetTemplateBody() string
}
type BaseTemplate struct {
	TemplateBody string
}

func (t *BaseTemplate) GetTemplateBody() string {
	return t.TemplateBody
}
func (t *BaseTemplate) GetTemplateHeader() string {
	return templateHeader
}

type CustomerTemplate struct {
	*BaseTemplate
}

func NewCustomerTemplate(path string) (*CustomerTemplate, error) {
	templateBody, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &CustomerTemplate{
		&BaseTemplate{TemplateBody: string(templateBody)},
	}, nil
}