package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Version should be replaced by the makefile
var Version = "unstable"

// Command-line action
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show moloon version",
	Run: func(cmd *cobra.Command, args []string) {
		versionInformation := "Version " + Version
		color.Green(versionInformation)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
