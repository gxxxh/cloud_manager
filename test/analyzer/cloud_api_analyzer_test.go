package cloud_api_analyzer_test

import (
	cloud_manager "cloud_manager/src/analyzer"
	"cloud_manager/src/codegen/openstack"
	"testing"
)

func TestExtractCloudAPIs(t *testing.T) {
	//client := ecs.Client{}
	client := openstack.OpenstackClient{}
	analyzer := cloud_manager.NewCloudAPIAnalyzer()
	analyzer.ExtractCloudAPIs(client)
	analyzer.SaveToJson("E:\\gopath\\src\\multicloud_service\\doc\\openstack.json")
}
