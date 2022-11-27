package main

import (
	i "github.com/RomanosTrechlis/go-retrieve/cli"
)

// initCmd represents the init command
var initCmd = i.RegisterInitCmd(rootCmd)

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("filename", "f", "", "Initialize from existing file")
	initCmd.Flags().BoolP("overwrite", "o", false, "Overwrite existing config when initializing from existing file")
}
