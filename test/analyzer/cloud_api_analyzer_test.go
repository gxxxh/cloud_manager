package cloud_api_analyzer_test

import (
	multicloud_service "multicloud_service/src/analyzer"
	"multicloud_service/src/codegen/openstack"
	"testing"
)

func TestExtractCloudAPIs(t *testing.T) {
	//client := ecs.Client{}
	client := openstack.OpenstackClient{}
	analyzer := multicloud_service.NewCloudAPIAnalyzer()
	analyzer.ExtractCloudAPIs(client)
	analyzer.SaveToJson("E:\\gopath\\src\\multicloud_service\\doc\\openstack.json")
}
