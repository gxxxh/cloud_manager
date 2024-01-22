package cmd

import (
	"github.com/kube-stack/multicloud_service/src/analyzer"
	"github.com/kube-stack/multicloud_service/src/code_generator/gen"
	"github.com/spf13/cobra"
	"log"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze -f sdkFilePath",
	Short: "analyze cloud code based on the sdk ",
	Long:  "analyze cloud code based on the sdk",
	RunE:  AnalyzeCloudSDK,
}
var sdkFilePath string

func AnalyzeCloudSDK(cmd *cobra.Command, args []string) error {

	ma := analyzer.NewModuleAnalyzer()
	resourceInfos, err := ma.DoAnalyze(sdkFilePath)
	if err != nil {
		return err
	}
	log.Printf("find %d resources", len(resourceInfos))
	return nil
}

// flags
var configFile string

var genBasicCmd = &cobra.Command{
	Use:   "gen_basic -f configFile",
	Short: "gen client and registry code based on the config file",
	Long:  "gen client and registry cloud code based on the config file",
	RunE:  GenBasicCode,
}

func GenBasicCode(cmd *cobra.Command, args []string) error {
	config := gen.LoadCloudConfig(configFile)
	err := gen.GenBasicCode(config)
	if err != nil {
		return err
	}
	return nil
}

var GenRequestCmd = &cobra.Command{
	Use:   "gen_request -f configFile",
	Short: "gen request and result code based on the config file",
	Long:  "gen request and result code based on the config file",
	RunE:  GenRequestCode,
}

func GenRequestCode(cmd *cobra.Command, args []string) error {
	config := gen.LoadCloudConfig(configFile)
	err := gen.GenRequestCode(config)
	if err != nil {
		return err
	}
	return nil
}

var GenAllCmd = &cobra.Command{
	Use:   "gen_all -f configFile",
	Short: "gen all code based on the config file",
	Long:  "gen all code based on the config file",
	RunE:  GenAllCode,
}

// 仅用于项目验收

var cloudtype string
var resource string
var GenCmd = &cobra.Command{
	Use:   "gen -t resource -c cloud",
	Short: "gen code for cloud resource",
	Long:  "gen code for cloud resource",
	RunE:  GenAllCode,
}

func GenAllCode(cmd *cobra.Command, args []string) error {
	if cloudtype == "openstack" {
		configFile = "/mnt/e/gopath/src/multicloud_service/src/code_generator/configs/openstack_linux.json"
	}
	log.Println("start gen code for resource:", resource, " cloud:", cloudtype)
	config := gen.LoadCloudConfig(configFile)
	err := gen.GenCloudCode(config)
	if err != nil {
		return err
	}
	log.Println("gen code for resource:", resource, " cloud:", cloudtype, " success")
	return nil
}
