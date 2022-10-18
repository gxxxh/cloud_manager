package aliyun

import (
	"cloud_manager/src/analyzer"
	"cloud_manager/src/code_generator"
	"log"
)

func GenCreateRequestRegistry(templatePath string, requestInfos []analyzer.RequestInfo, importPath string) ([]byte, error) {
	createRequestRegistryTemplate, err := code_generator.NewCustomerTemplate(templatePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	code, err := code_generator.GenerateTemplate(createRequestRegistryTemplate.GetTemplateBody(),
		map[string]interface{}{
			"requestInfos": requestInfos,
			"importPath":   importPath,
		},
		map[string]interface{}{
			"packageName":    "test",
			"templateHeader": createRequestRegistryTemplate.GetTemplateHeader(),
		})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return code, err
}
