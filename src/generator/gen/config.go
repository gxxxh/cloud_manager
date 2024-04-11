package gen

import (
	"encoding/json"
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
		log.Fatal("json unmarshal error, ", err)
	}
	InitCloudConfig(config)
	return config
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
