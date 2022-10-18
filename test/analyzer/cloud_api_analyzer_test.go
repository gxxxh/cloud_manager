package cloud_api_analyzer_test

import (
	cloud_manager "cloud_manager/src/analyzer"
	"testing"
)

func TestExtractCloudAPIs(t *testing.T) {
	analyzer := cloud_manager.CloudAPIAnalyzer{Kind: "aliyun"}
	analyzer.Init()
	analyzer.ExtractCloudAPIs()
	analyzer.SaveToJson()
}
