package utils

import (
	"fmt"
	"github.com/kube-stack/multicloud_service/src/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpper(t *testing.T) {
	fmt.Println(utils.UpperFirst("a"))
}

func TestParseRequestName(t *testing.T) {
	requestName := "CreateComputeV2ServersRequest"
	actionName, openstackResourceName, javaResourceName := utils.ParseRequestName(requestName)
	assert.Equal(t, "Create", actionName)
	assert.Equal(t, "ComputeV2Servers", openstackResourceName)
	assert.Equal(t, "server", javaResourceName)
}
