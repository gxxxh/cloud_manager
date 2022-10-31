package code_generator

import (
	cloud_manager "cloud_manager/src/analyzer"
	"cloud_manager/src/code_generator"
	"cloud_manager/src/utils"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"os"
	"path/filepath"
	"testing"
)

func TestGenCreateRequestRegistry(t *testing.T) {
	//client := openstack.OpenstackClient{}
	client := ecs.Client{}
	analyzer := cloud_manager.NewCloudAPIAnalyzer()
	analyzer.ExtractCloudAPIs(client)
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
	dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\keymanager\\v1\\secrets"
	ma := cloud_manager.NewModuleAnalyzer()
	resourceInfos, err := ma.DoAnalyze(dir)
	if err != nil {
		t.Error(err)
	}
	for _, resourceInfo := range resourceInfos {
		if len(resourceInfo.Actions) == 0 {
			continue
		}
		fmt.Printf("gen code for actions in resource %s\n", resourceInfo.ResourcePackageName)
		templatePath := "E:\\gopath\\src\\cloud_manager\\src\\code_generator\\templates\\openstack_request.tmpl"
		data := utils.Struct2Map(resourceInfo)
		code, err := code_generator.GenCode(templatePath, data, "openstack")
		if err != nil {
			t.Error(err)
		}
		basePath := "E:\\gopath\\src\\cloud_manager\\src\\codegen\\openstack"
		fileName := utils.JoinName(resourceInfo.ResourcePath, "openstack", "_") + ".go"
		filePath := filepath.Join(basePath, fileName)
		filePtr, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			t.Error(err)
		}
		defer filePtr.Close()
		filePtr.Write(code)
	}
}
