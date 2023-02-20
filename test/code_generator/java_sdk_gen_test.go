package code_generator

import (
	"github.com/kube-stack/multicloud_service/src/code_generator/gen"
	"testing"
)

func TestGenObjectClass(t *testing.T) {
	configPath := "E:\\gopath\\src\\multicloud_service\\test\\code_generator\\test_config.json"
	config := gen.LoadCloudConfig(configPath)
	java_generator := gen.NewJavaSDKGenerator(config)
	java_generator.GenResourceClass()
}

func TestGenResourceSpec(t *testing.T) {
	configPath := "E:\\gopath\\src\\multicloud_service\\test\\code_generator\\test_config.json"
	config := gen.LoadCloudConfig(configPath)
	java_generator := gen.NewJavaSDKGenerator(config)
	java_generator.GenResourceSpec()
}
