package gen

import (
	"encoding/json"
	"log"
	"os"
)

const (
	JavaResourceClassTemplate           = "java_resource_class.tmpl"
	JavaResourceDomainHeaderTemplate    = "java_resource_domain_header.tmpl"
	JavaResourceLifecycleHeaderTemplate = "java_resource_lifecycle_header.tmpl"
	JavaResourceSpecTemplate            = "java_resource_spec.tmpl"
	JavaResourceImalTemplate            = "java_resource_impl.tmpl"
)

// 递归生成Class
const JavaClassTemplate = "java_class.tmpl"

type CloudConfig struct {
	CloudType       string            `json:"CloudType"`
	RegistryConfigs []*RegistryConfig `json:"RegistryConfigs"`
	APICodeConfigs  []*APICodeConfig  `json:"APICodeConfigs"`
	JavaCodeConfig  CodeGenConfig     `json:"JavaCodeConfig"`
}

type RegistryConfig struct {
	CodeGenConfig       CodeGenConfig `json:"CodeGenConfig"`
	RegistryImportPaths []string      `json:"RegistryImportPaths"`
	CreateFuncPre       string        `json:"CreateFuncPre"`
	FuncAction          string        `json:"FuncAction"`
	CodeType            string        `json:"CodeType"` //Request or Response
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
	return config
}
