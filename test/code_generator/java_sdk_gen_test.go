package code_generator

import (
	"github.com/kube-stack/multicloud_service/src/code_generator/gen"
	"testing"
)

func TestGenObjectClass(t *testing.T) {
	configPath := "E:\\gopath\\src\\multicloud_service\\test\\code_generator\\test_config.json"
	config := gen.LoadCloudConfig(configPath)
	java_generator := gen.NewJavaSDKGenerator(config)
	java_generator.GenResourceClass("ComputeV2Servers")
}

func TestGenResourceSpec(t *testing.T) {
	configPath := "E:\\gopath\\src\\multicloud_service\\test\\code_generator\\test_config.json"
	config := gen.LoadCloudConfig(configPath)
	java_generator := gen.NewJavaSDKGenerator(config)
	java_generator.GenResourceSpec("ComputeV2Servers")
}

func TestGenResourceLifecycle(t *testing.T) {
	configPath := "E:\\gopath\\src\\multicloud_service\\test\\code_generator\\test_config.json"
	config := gen.LoadCloudConfig(configPath)
	java_generator := gen.NewJavaSDKGenerator(config)
	java_generator.GenResourceLifecycle("ComputeV2Servers")
}

func TestGenResourceDomain(t *testing.T) {
	configPath := "E:\\gopath\\src\\multicloud_service\\test\\code_generator\\test_config.json"
	config := gen.LoadCloudConfig(configPath)
	java_generator := gen.NewJavaSDKGenerator(config)
	java_generator.GenResourceDomain("ComputeV2Servers")
}

func TestGenResourceImpl(t *testing.T) {
	configPath := "E:\\gopath\\src\\multicloud_service\\test\\code_generator\\test_config.json"
	config := gen.LoadCloudConfig(configPath)
	java_generator := gen.NewJavaSDKGenerator(config)
	java_generator.GenResourceImpl("ComputeV2Servers")
}
func TestGenAll(t *testing.T) {
	configPath := "E:\\gopath\\src\\multicloud_service\\test\\code_generator\\test_config.json"
	config := gen.LoadCloudConfig(configPath)
	java_generator := gen.NewJavaSDKGenerator(config)
	//java_generator.GenAll("ComputeV2Servers")
	java_generator.GenAll("ImageserviceV2Images")
	java_generator.GenAll("BlockstorageV3Snapshots")
	java_generator.GenAll("NetworkingV2Networks")
	java_generator.GenAll("NetworkingV2ExtensionsLayer3Routers")
}
