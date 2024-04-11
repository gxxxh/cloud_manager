package gen

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	multicloud_service "github.com/kube-stack/multicloud_service/src/analyzer"
	openstack "github.com/kube-stack/multicloud_service/src/codegen/openstack"
	"github.com/kube-stack/multicloud_service/src/utils"
	"log"
	"path/filepath"
	"strings"
)

type CloudAPIGenerator struct {
	CloudType        string
	CloudAPIAnalyzer *multicloud_service.CloudAPIAnalyzer
	CodeGenerator    *CodeGenerator
}

func NewCloudAPIGenerator(cloudType string) *CloudAPIGenerator {
	return &CloudAPIGenerator{
		CloudType:        cloudType,
		CloudAPIAnalyzer: multicloud_service.NewCloudAPIAnalyzer(),
	}
}

func (g *CloudAPIGenerator) DoGen(config *CloudConfig) error {
	switch config.CloudType {
	case "Aliyun":
		client := ecs.Client{}
		g.CloudAPIAnalyzer.ExtractCloudAPIs(client)
		for _, registryConfig := range config.RegistryConfigs {
			g.GenRegistryCode(config.CloudType, registryConfig)
		}
	case "Openstack":
		g.GenRequestCode(config)
		g.GenBasicCode(config)
	default:
		log.Fatalln("unsupport cloud kind, ", config.CloudType)
	}
	//log.Printf("generate all cloud code for %s done", config.CloudType)
	return nil
}

func (g *CloudAPIGenerator) GenRequestCode(config *CloudConfig) error {
	//log.Printf("start to generate request code for %s", config.CloudType)
	analyzer := multicloud_service.NewCloudAPIAnalyzer()
	switch config.CloudType {
	case "Aliyun":
		client := ecs.Client{}
		analyzer.ExtractCloudAPIs(client)
		fmt.Println("Gen Aliyun Request Code Done")
	case "Openstack":
		for _, apiCodeConfig := range config.APICodeConfigs {
			g.GenAPICode(apiCodeConfig)
		}
	default:
		log.Fatalln("unsupport cloud kind, ", config.CloudType)
	}
	//log.Printf("generate request code for %s done", config.CloudType)
	return nil
}

func (g *CloudAPIGenerator) GenBasicCode(config *CloudConfig) error {
	analyzer := multicloud_service.NewCloudAPIAnalyzer()
	//log.Printf("Gen Basic for %s Code", config.CloudType)
	switch config.CloudType {
	case "Aliyun":
		client := ecs.Client{}
		analyzer.ExtractCloudAPIs(client)
		for _, registryConfig := range config.RegistryConfigs {
			g.GenRegistryCode(config.CloudType, registryConfig)
		}
	case "Openstack":
		client := openstack.OpenstackClient{}
		g.CloudAPIAnalyzer.ExtractCloudAPIs(client)
		for _, registryConfig := range config.RegistryConfigs {
			g.GenRegistryCode(config.CloudType, registryConfig)
		}

	default:
		log.Fatalln("unsupport cloud kind, ", config.CloudType)
	}
	//log.Printf("Gen Basic for %s Code done", config.CloudType)
	return nil

}

func (g *CloudAPIGenerator) GenRegistryCode(cloudKind string, registryConfig *RegistryConfig) {
	requestRegistryInfo := g.CloudAPIAnalyzer.ExtractRequestInfos(registryConfig.CreateFuncPre)
	if cloudKind == "Openstack" && registryConfig.CodeType == "Response" {
		for _, requestInfo := range requestRegistryInfo.RequestInfos {
			requestInfo.CreateFunctionName = strings.Replace(requestInfo.CreateFunctionName, "New", "Extract", -1)
			requestInfo.CreateFunctionName = requestInfo.CreateFunctionName[0:len(requestInfo.CreateFunctionName)-len("Request")] + "Response"
		}
	}
	requestRegistryInfo.ImportPaths = append(requestRegistryInfo.ImportPaths, registryConfig.RegistryImportPath)
	data := utils.Struct2Map(requestRegistryInfo)
	params := make(map[string]interface{})
	params["packageName"] = "registry"
	params["kind"] = cloudKind
	params["action"] = registryConfig.FuncAction
	params["type"] = registryConfig.CodeType
	codePath := filepath.Join(registryConfig.CodeGenConfig.CodePath, strings.ToLower(cloudKind)+"_"+strings.ToLower(registryConfig.FuncAction)+"_"+strings.ToLower(registryConfig.CodeType)+"_registry.go")
	g.CodeGenerator.GenAndSaveCode(registryConfig.CodeGenConfig.TemplatePath, codePath, data, params)
}

func (g *CloudAPIGenerator) GenAPICode(config *APICodeConfig) {
	ma := multicloud_service.NewModuleAnalyzer()
	resourceInfos, err := ma.DoAnalyze(config.SourceCodePath)
	if err != nil {
		log.Fatalln("analyze source code error, ", err)
	}
	for _, resourceInfo := range resourceInfos {
		if len(resourceInfo.ActionInfos) == 0 {
			continue
		}
		//log.Printf("cmd.go %s code for actions in resource %s\n", config.CodeType, resourceInfo.ResourcePackageName)
		data := utils.Struct2Map(resourceInfo)
		params := map[string]interface{}{
			"packageName": "openstack",
		}
		filenName := utils.JoinName(resourceInfo.ResourcePath, "openstack", "_") + "_" + strings.ToLower(config.CodeType) + ".go"
		codePath := filepath.Join(config.CodeGenConfig.CodePath, filenName)
		g.CodeGenerator.GenAndSaveCode(config.CodeGenConfig.TemplatePath, codePath, data, params)
	}
}
