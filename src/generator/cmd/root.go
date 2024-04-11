package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = cobra.Command{
	Use:   "cloudcodegen",
	Short: "cloudcodegen is used to genereate multi cloud code in the same pattern",
	Long:  "A golang code generator to generate cloud sdk code in the same pattern, supprot openstack and aliyun, details in https://github.com/kube-stack/multicloud_service",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func Init() {
	rootCmd.AddCommand(versionCmd)

	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Flags().StringVarP(&sdkFilePath, "sdkFilePath", "f", "", "sdk file path")
	analyzeCmd.MarkFlagRequired("sdkFilePath")

	rootCmd.AddCommand(genBasicCmd)
	genBasicCmd.Flags().StringVarP(&configFile, "configFile", "f", "", "config file")
	genBasicCmd.MarkFlagRequired("configFile")

	rootCmd.AddCommand(GenRequestCmd)
	GenRequestCmd.Flags().StringVarP(&configFile, "configFile", "f", "", "config file")
	GenRequestCmd.MarkFlagRequired("configFile")

	rootCmd.AddCommand(GenAllCmd)
	GenAllCmd.Flags().StringVarP(&configFile, "configFile", "f", "", "config file")
	GenAllCmd.MarkFlagRequired("configFile")

	rootCmd.AddCommand(GenCmd)
	GenCmd.Flags().StringVarP(&resource, "resource", "t", "", "cloud resource type")
	GenCmd.MarkFlagRequired("resource")
	GenCmd.Flags().StringVarP(&cloudtype, "cloud", "c", "", "cloud type")
	GenCmd.MarkFlagRequired("cloud")

}
