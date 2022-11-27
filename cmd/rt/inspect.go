package main

import i "github.com/RomanosTrechlis/go-retrieve/cli"

// initCmd represents the init command
var inspectTemplateCmd = i.RegisterInspectTemplateCmd()

func init() {
	rootCmd.AddCommand(inspectTemplateCmd)
}
