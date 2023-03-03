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
		fmt.Fprintf(os.Stdout, "generate code success!")
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)

	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringVarP(&configFile, "configFile", "f", "", "config file")
	genCmd.MarkFlagRequired("configFile")
}
