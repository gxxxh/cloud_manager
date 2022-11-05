package cloud_api_analyzer

import (
	"cloud_manager/src/analyzer"
	"log"
	"testing"
)

func TestModuleAnalyzer(t *testing.T) {
	dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\compute\\v2\\servers"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\identity\\v3\\extensions\\oauth1"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\blockstorage\\apiversions"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\blockstorage\\v3\\qos"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\clustering\\v1\\clusters"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\clustering\\v1\\profiletypes"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\objectstorage\\v1\\containers"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack"
	ma := analyzer.NewModuleAnalyzer()
	requestInfos, err := ma.DoAnalyze(dir)
	if err != nil {
		t.Error(err)
	}
	log.Println(len(requestInfos))
}
