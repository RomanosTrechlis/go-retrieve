package main

import (
	"github.com/RomanosTrechlis/go-retrieve/cli"
)

// destroyCmd represents the destroy command
var destroyCmd = cli.RegisterNukeCmd(rootCmd)

func init() {
	rootCmd.AddCommand(destroyCmd)
}
