package main

import (
	"github.com/RomanosTrechlis/go-retrieve/cli"
)

// profileCmd represents the profile command
var profileCmd = cli.RegisterProfileCmd(rootCmd)

// inspectCmd represents the inspect command
var inspectCmd = cli.RegisterProfileInspectCmd()

// profileListCmd represents the list command
var profileListCmd = cli.RegisterProfileListCmd()

func init() {
	rootCmd.AddCommand(profileCmd)

	profileCmd.AddCommand(inspectCmd)
	inspectCmd.Flags().BoolP("dump", "d", false, "Dump profile configuration")

	profileCmd.AddCommand(profileListCmd)
}
