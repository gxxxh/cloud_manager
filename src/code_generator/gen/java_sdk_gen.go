package gen

import (
	multicloud_service "github.com/kube-stack/multicloud_service/src/analyzer"
	"github.com/kube-stack/multicloud_service/src/utils"
	"log"
	"path/filepath"
)

type JavaSDKGenerator struct {
	Config               *CloudConfig
	ModuleAnalyzer       *multicloud_service.ModuleAnalyzer
	Analyzer             *multicloud_service.CloudAPIAnalyzer
	RequestResourceInfos []*multicloud_service.OpenstackResourceInfo
	ResultResourceInfos  []*multicloud_service.OpenstackResourceInfo
}

func NewJavaSDKGenerator(config *CloudConfig) *JavaSDKGenerator {
	cloudAPIAnalyzer := multicloud_service.NewCloudAPIAnalyzer()
	javaSDKGenerator := &JavaSDKGenerator{
		Config:   config,
		Analyzer: cloudAPIAnalyzer,
	}
	ma := multicloud_service.NewModuleAnalyzer()
	for _, apiConfig := range config.APICodeConfigs {
		resourceInfos, err := ma.DoAnalyze(apiConfig.SourceCodePath)
		if err != nil {
			log.Fatalln("analyze source code error, ", err)
		}
		if apiConfig.CodeType == "Request" {
			javaSDKGenerator.RequestResourceInfos = resourceInfos
		} else if apiConfig.CodeType == "Result" {
			javaSDKGenerator.ResultResourceInfos = resourceInfos
		}
	}
	return javaSDKGenerator
}

// class code → java_resource_class.tmpl
func (j *JavaSDKGenerator) GenResourceClass() {
	data := make(map[string]interface{})
	for _, requestResourceInfo := range j.RequestResourceInfos {
		javaResourceName := utils.GetJavaResourceName(j.Config.CloudType, requestResourceInfo.ResourcePackageName)
		data["JavaResourceName"] = javaResourceName
		codePath := filepath.Join(j.Config.JavaCodeConfig.Class.CodePath, javaResourceName+".java")
		GenAndSaveCode(j.Config.JavaCodeConfig.Class.TemplatePath, codePath, data, nil)
	}

}

// spec code → java_resource_spec.tmpl
func (j *JavaSDKGenerator) GenResourceSpec() {
	data := make(map[string]interface{})
	for _, requestResourceInfo := range j.RequestResourceInfos {
		javaResourceName := utils.GetJavaResourceName(j.Config.CloudType, requestResourceInfo.ResourcePackageName)
		data["JavaResourceName"] = javaResourceName
		codePath := filepath.Join(j.Config.JavaCodeConfig.Spec.CodePath, javaResourceName+"Spec.java")
		GenAndSaveCode(j.Config.JavaCodeConfig.Spec.TemplatePath, codePath, data, nil)
	}
}

func (j *JavaSDKGenerator) GenResourceDomain() {

}

func (j *JavaSDKGenerator) GenResourceLifecycle() {

}

func (j *JavaSDKGenerator) GenResourceImpl() {

}
