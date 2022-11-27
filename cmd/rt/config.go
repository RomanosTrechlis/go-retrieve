package main

import (
	"github.com/RomanosTrechlis/go-retrieve/cli"
)

// configCmd represents the config command
var configCmd = cli.RegisterConfig(rootCmd)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolP("dump", "d", false, "Dump configuration file")
}
