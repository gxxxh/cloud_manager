package cloud_api_analyzer

import (
	"github.com/kube-stack/multicloud_service/src/analyzer"
	"log"
	"testing"
)

func TestModuleAnalyzer(t *testing.T) {
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\compute\\v2\\servers"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\baremetal\\apiversions"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\identity\\v3\\extensions\\oauth1"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\blockstorage\\apiversions"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\blockstorage\\v3\\qos"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\clustering\\v1\\clusters"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\clustering\\v1\\profiletypes"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\objectstorage\\v1\\containers"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\db\\v1\\configurations"
	dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack"
	ma := analyzer.NewModuleAnalyzer()
	resourceInfos, err := ma.DoAnalyze(dir)
	for _, resourceInfo := range resourceInfos {
		for _, actionInfo := range resourceInfo.ActionInfos {
			if actionInfo.ResultExtractInfo == nil &&
				actionInfo.PageExtractInfo == nil {
				log.Printf("Action %v in package %v is not handled\n", actionInfo.ActionName, resourceInfo.ResourcePath)
			}
		}
	}
	if err != nil {
		t.Error(err)
	}
}
