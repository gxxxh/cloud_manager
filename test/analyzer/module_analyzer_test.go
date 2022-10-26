package cloud_api_analyzer

import (
	"cloud_manager/src/analyzer"
	"testing"
)

func TestModuleAnalyzer(t *testing.T) {
	//dir := "D:\\gh\\cloud\\gophercloud-master\\openstack\\compute"
	dir := "D:\\gh\\cloud\\gophercloud-master\\openstack\\compute\\v2\\servers"
	ma := analyzer.NewModuleAnalyzer()
	_, err := ma.DoAnalyze(dir)
	if err != nil {
		t.Error(err)
	}
}
