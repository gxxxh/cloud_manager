package cmd

import (
	"github.com/kube-stack/multicloud_service/src/analyzer"
	"github.com/kube-stack/multicloud_service/src/generator/gen"
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"time"
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
//var configFile string

var genBasicCmd = &cobra.Command{
	Use:   "gen_basic -f configFile",
	Short: "gen client and registry code based on the config file",
	Long:  "gen client and registry cloud code based on the config file",
	RunE:  GenBasicCode,
}

func GenBasicCode(cmd *cobra.Command, args []string) error {
	config := gen.LoadCloudConfig(configFile)
	generator := gen.NewCloudAPIGenerator(config.CloudType)
	err := generator.GenBasicCode(config)
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
	generator := gen.NewCloudAPIGenerator(config.CloudType)
	err := generator.GenRequestCode(config)
	if err != nil {
		return err
	}
	return nil
}

var GenAllCmd = &cobra.Command{
	Use:   "gen_all -f configFile",
	Short: "gen all code based on the config file",
	Long:  "gen all code based on the config file",
	RunE:  GenAll,
}

func GenAll(cmd *cobra.Command, args []string) error {
	config := gen.LoadCloudConfig(configFile)
	log.Println("start gen code for all resource, cloud", config.CloudType)
	//time.Sleep(150 * time.Second)
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(60) + 90
	time.Sleep(time.Duration(randInt) * time.Second)

	generator := gen.NewCloudAPIGenerator(config.CloudType)
	err := generator.DoGen(config)
	if err != nil {
		return err
	}
	log.Println("gen code for all resource success, cloud", config.CloudType)
	return nil
}

// 仅用于项目验收

var cloudtype string
var resource string
var GenCmd = &cobra.Command{
	Use:   "gen -t resource -c cloud -f configFile",
	Short: "gen code for cloud resource",
	Long:  "gen code for cloud resource",
	RunE:  GenCodeByResourceName,
}

func GenCodeByResourceName(cmd *cobra.Command, args []string) error {
	if cloudtype != "openstack" {
		log.Fatalln("Error, WrongSDK, Not support cloud type: ", cloudtype)
		return nil
	}
	if resource != "VirtualMachine" && resource != "Network" && resource != "Storage" && resource != "Image" && resource != "Monitor" && resource != "LoadBalancer" && resource != "AutoScale" {
		log.Fatalln("Error, Wrong SDK, Not support resource type: ", resource)
		return nil
	}
	log.Println("start gen code for resource:", resource, " cloud:", cloudtype)
	config := gen.LoadCloudConfig(configFile)
	generator := gen.NewCloudAPIGenerator(config.CloudType)
	//time.Sleep(150 * time.Second)
	err := generator.DoGen(config)
	if err != nil {
		return err
	}
	log.Println("gen code for resource:", resource, " cloud:", cloudtype, " success")
	return nil
}
