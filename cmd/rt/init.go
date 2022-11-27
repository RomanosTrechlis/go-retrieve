package main

import (
	i "github.com/RomanosTrechlis/go-retrieve/cli"
)

// initCmd represents the init command
var initCmd = i.RegisterSimpleInitCmd(rootCmd)

func init() {
	rootCmd.AddCommand(initCmd)
}
