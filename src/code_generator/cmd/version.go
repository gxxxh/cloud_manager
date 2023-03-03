package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of cloudcodegen",
	Long:  "Version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cloudcodegen v1.0")
	},
}
