package cloud_api_analyzer

import (
	"cloud_manager/src/analyzer"
	"log"
	"testing"
)

func TestModuleAnalyzer(t *testing.T) {
	dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\compute\\v2\\servers"
	ma := analyzer.NewModuleAnalyzer()
	requestInfos, err := ma.DoAnalyze(dir)
	if err != nil {
		t.Error(err)
	}
	log.Println(len(requestInfos))
}
