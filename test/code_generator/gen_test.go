package code_generator

import (
	cloud_manager "cloud_manager/src/analyzer"
	"cloud_manager/src/code_generator/aliyun"
	"fmt"
	"os"
	"testing"
)

func TestGenCreateRequestRegistry(t *testing.T) {
	analyzer := cloud_manager.CloudAPIAnalyzer{Kind: "aliyun"}
	analyzer.Init()
	analyzer.ExtractCloudAPIs()
	requestInfos := analyzer.ExtractRequestInfos()
	templatePath := "E:\\gopath\\src\\cloud_manager\\src\\code_generator\\aliyun\\templates\\request_map.tmpl"
	code, err := aliyun.GenCreateRequestRegistry(templatePath, requestInfos, "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs")
	if err != nil {
		t.Error(err)
	}

	filePtr, err := os.Create("create_request_registry.go")
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	defer filePtr.Close()
	filePtr.Write(code)

}
