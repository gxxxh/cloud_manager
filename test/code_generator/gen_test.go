package code_generator

import (
	cloud_manager "cloud_manager/src/analyzer"
	"cloud_manager/src/code_generator"
	"cloud_manager/src/utils"
	"fmt"
	"os"
	"testing"
)

func TestGenCreateRequestRegistry(t *testing.T) {
	analyzer := cloud_manager.CloudAPIAnalyzer{Kind: "aliyun"}
	analyzer.Init()
	analyzer.ExtractCloudAPIs()
	requestRegistryInfo := analyzer.ExtractRequestInfos()
	templatePath := "E:\\gopath\\src\\cloud_manager\\src\\code_generator\\templates\\request_map.tmpl"
	requestRegistryInfo.ImportPaths = []string{"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"}
	data := utils.Struct2Map(requestRegistryInfo)
	code, err := code_generator.GenCode(templatePath, data, "aliyun")
	//code, err := code_generator.GenCreateRequestRegistry(templatePath, requestRegistryInfo, "aliyun")
	if err != nil {
		t.Error(err)
	}

	filePtr, err := os.Create("E:\\gopath\\src\\cloud_manager\\src\\codegen\\aliyun\\create_request_registry.go")
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	defer filePtr.Close()
	filePtr.Write(code)
}

func TestGenOpenstackCode(t *testing.T) {
	dir := "D:\\gh\\cloud\\gophercloud-master\\openstack\\compute"
	//dir := "D:\\gh\\cloud\\gophercloud-master\\openstack\\compute\\v2\\servers"
	ma := cloud_manager.NewModuleAnalyzer()
	resourceInfos, err := ma.DoAnalyze(dir)
	if err != nil {
		t.Error(err)
	}
	for _, resourceInfo := range resourceInfos {
		templatePath := "E:\\gopath\\src\\cloud_manager\\src\\code_generator\\templates\\openstack_request.tmpl"
		data := utils.Struct2Map(resourceInfo)
		code, err := code_generator.GenCode(templatePath, data, "openstack")
		if err != nil {
			t.Error(err)
		}
		filePath := fmt.Sprintf("E:\\gopath\\src\\cloud_manager\\src\\codegen\\openstack\\%s.go", resourceInfo.ResourcePackageName)
		filePtr, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			t.Error(err)
		}
		defer filePtr.Close()
		filePtr.Write(code)
	}
}
