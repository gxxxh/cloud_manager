package utils

import (
	"github.com/kube-stack/multicloud_service/src/utils"
	"testing"
)

func TestCreateMultiDir(t *testing.T) {
	dirPath := "E:\\gopath\\src\\multicloud_service\\out\\api\\models\\openstack"
	err := utils.CreateMultiDir(dirPath)
	if err != nil {
		t.Error(err)
	}
}
