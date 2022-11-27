package main

import (
	"github.com/RomanosTrechlis/go-retrieve/cli"
)

// registryCmd represents the registry command
var registryCmd = cli.RegisterRegistryCmd()

// createCmd represents the create command
var createCmd = cli.RegisterRegistryAddCmd()

// defineCmd represents the define command
var defineCmd = cli.RegisterRegistryDefineCmd()

func init() {
	rootCmd.AddCommand(registryCmd)

	registryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("exclude", "e", "", "List of directories to exclude (CSV)")

	registryCmd.AddCommand(defineCmd)
}
