package main

import (
	"github.com/RomanosTrechlis/go-retrieve/cli"
)

// templateCmd represents the template command
var templateCmd = cli.RegisterGetCmd()

// templateListCmd represents the list command
var templateListCmd = cli.RegisterGetListCmd()

func init() {
	rootCmd.AddCommand(templateCmd)
	templateCmd.Flags().StringP("destination", "d", "",
		"Download template to specific destination")
	templateCmd.Flags().BoolP("flat", "f", false,
		`Download files in the destination instead of the template directory. 
Note: this applies to dependencies also.`)

	//templateCmd.AddCommand(templateListCmd)
	//templateListCmd.Flags().BoolP("dump", "d", false, "Dump profile configuration")
	rootCmd.AddCommand(templateListCmd)
	templateListCmd.Flags().BoolP("all", "a", false,
		`List the registered templates of all the sources`)
}
