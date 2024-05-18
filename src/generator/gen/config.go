package gen

import (
	"encoding/json"
	"fmt"
	"github.com/kube-stack/multicloud_service/src/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// tempalte path

const TemplatePath = "template"
const CodePath = ""

// java template
const (
	JavaResourceClassTemplate           = "java_resource_class.tmpl"
	JavaResourceDomainHeaderTemplate    = "java_resource_domain_header.tmpl"
	JavaResourceLifecycleHeaderTemplate = "java_resource_lifecycle_header.tmpl"
	JavaResourceSpecTemplate            = "java_resource_spec.tmpl"
	JavaResourceImalTemplate            = "java_resource_impl.tmpl"
)

// go template
const ()

// 递归生成Class
const JavaClassTemplate = "java_class.tmpl"

type CloudConfig struct {
	CloudType       string            `json:"CloudType"`
	SourceCodePath  string            `json:"SourceCodePath"`
	JavaCodeConfig  CodeGenConfig     `json:"JavaCodeConfig"`
	GoCodeConfig    CodeGenConfig     `json:"GoCodeConfig"`
	RegistryConfigs []*RegistryConfig `json:"RegistryConfigs"`
	APICodeConfigs  []*APICodeConfig  `json:"APICodeConfigs"`
}

type RegistryConfig struct {
	CodeGenConfig      CodeGenConfig `json:"CodeGenConfig"`
	RegistryImportPath string        `json:"RegistryImportPath"`
	CreateFuncPre      string        `json:"CreateFuncPre"`
	FuncAction         string        `json:"FuncAction"`
	CodeType           string        `json:"CodeType"` //Request or Response
}

type APICodeConfig struct {
	CodeGenConfig  CodeGenConfig `json:"CodeGenConfig"`
	SourceCodePath string        `json:"SourceCodePath"`
	CodeType       string        `json:"CodeType"` //Request or Result
}
type CodeGenConfig struct {
	TemplatePath string `json:"TemplatePath"`
	CodePath     string `json:"CodePath"`
}

func LoadCloudConfig(configPath string) *CloudConfig {
	config := &CloudConfig{}
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("read config error, err=", err)
	}
	err = json.Unmarshal(bytes, config)
	if err != nil {
		log.Fatal("Wrong Format: json unmarshal error, ", err)
	}
	CheckSDK(config)
	CheckFormat(config)
	CheckParameters(config)
	InitCloudConfig(config)
	existed, err := utils.PathExists(config.SourceCodePath)
	if err != nil {
		log.Fatal("check source code path error, err=", err)
	}
	if config.SourceCodePath == "" || !existed {
		log.Fatal("source code path is empty")
	}
	return config
}
func CheckSDK(config *CloudConfig) {
	existed, err := utils.PathExists(config.SourceCodePath)
	if err != nil {
		log.Fatal("Wrong SDK, check SDK path error, err=", err)
	}
	if !existed {
		log.Fatal("Wrong SDK: SDK code path is not existed, ", config.SourceCodePath)
	}

}

func CheckFormat(config *CloudConfig) {
	if config.CloudType == "" {
		log.Fatal("Wrong Format: cloud type is empty")
	}
	if config.GoCodeConfig.TemplatePath == "" {
		log.Fatal("Wrong Format: go code template path is empty")
	}
	if config.GoCodeConfig.CodePath == "" {
		log.Fatal("Wrong Format: go code path is empty")
	}
	if config.SourceCodePath == "" {
		log.Fatal("Wrong Format: source code path is empty")
	}
}

func CheckParameters(config *CloudConfig) {
	if config.CloudType != "Aliyun" && config.CloudType != "Openstack" {
		log.Fatal(fmt.Sprintf("Wrong Parameters: cloud type is not supported, %s", config.CloudType))
	}
	existed, _ := utils.PathExists(config.GoCodeConfig.TemplatePath)
	if !existed {
		log.Fatal("Wrong Parameters: go code template path is not existed, ", config.GoCodeConfig.TemplatePath)
	}
	existed, _ = utils.PathExists(config.GoCodeConfig.CodePath)
	if !existed {
		log.Fatal("Wrong Parameters: go code path is not existed, ", config.GoCodeConfig.CodePath)
	}

	if len(config.RegistryConfigs) != len(config.APICodeConfigs) {
		log.Fatal("Wrong Parameters: registry config and api code config is not matched")
	}
	if config.RegistryConfigs == nil || config.APICodeConfigs == nil {
		log.Fatal("Wrong Parameters: registry config or api code config is nil")
	}
}

// 填充一些需要的字段
func InitCloudConfig(config *CloudConfig) *CloudConfig {

	// init registry path
	for _, registryConfig := range config.RegistryConfigs {
		registryConfig.CodeGenConfig = CodeGenConfig{
			TemplatePath: filepath.Join(config.GoCodeConfig.TemplatePath, "registry.tmpl"),
			CodePath:     filepath.Join(config.GoCodeConfig.CodePath, "registry"),
		}
	}

	//init api code path
	for _, apiCodeConfig := range config.APICodeConfigs {
		apiCodeConfig.CodeGenConfig = CodeGenConfig{
			TemplatePath: filepath.Join(config.GoCodeConfig.TemplatePath, strings.ToLower(config.CloudType)+"_"+strings.ToLower(apiCodeConfig.CodeType)+".tmpl"),
			CodePath:     filepath.Join(config.GoCodeConfig.CodePath, strings.ToLower(config.CloudType)),
		}
		apiCodeConfig.SourceCodePath = config.SourceCodePath
	}
	//switch config.CloudType {
	//case "Aliyun":
	//	config.RegistryConfigs[0].RegistryImportPaths = append(config.RegistryConfigs[0].RegistryImportPaths, "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs")
	//	config.RegistryConfigs[0].CreateFuncPre = "Create"
	//	config.RegistryConfigs[0].FuncAction = "Create"
	//	config.RegistryConfigs[0].CodeType = "Request"
	//case "Openstack":
	//	// Request registry config
	//	config.RegistryConfigs[0].RegistryImportPaths = append(config.RegistryConfigs[0].RegistryImportPaths, "github.com/kube-stack/multicloud_service/src/codegen/openstack")
	//	config.RegistryConfigs[0].CreateFuncPre = "New"
	//	config.RegistryConfigs[0].FuncAction = "Create"
	//	config.RegistryConfigs[0].CodeType = "Request"
	//	//Response registry config
	//	config.RegistryConfigs = append(config.RegistryConfigs, &RegistryConfig{})
	//	config.RegistryConfigs[1].CodeGenConfig = config.RegistryConfigs[0].CodeGenConfig
	//	config.RegistryConfigs[0].RegistryImportPaths = append(config.RegistryConfigs[0].RegistryImportPaths, "github.com/kube-stack/multicloud_service/src/codegen/openstack")
	//	config.RegistryConfigs[0].CreateFuncPre = "New"
	//	config.RegistryConfigs[0].FuncAction = "Extract"
	//	config.RegistryConfigs[0].CodeType = "Response"
	//	//api code config
	//
	//}
	return config
}
