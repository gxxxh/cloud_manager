package code_generator

import (
	"cloud_manager/src/code_generator/openstack"
	"testing"
)

func TestModuleAnalyzer(t *testing.T) {
	//dir := "D:\\gh\\cloud\\gophercloud-master\\openstack"
	dir := "D:\\gh\\cloud\\gophercloud-master\\openstack\\compute\\v2\\servers"
	ma := openstack.NewModuleAnalyzer(dir)
	err := ma.DoAnalyze()
	if err != nil {
		t.Error(err)
	}
}
