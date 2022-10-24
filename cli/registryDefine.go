/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/registry"
)

// defineCmd represents the define command
var defineCmd = &cobra.Command{
	Use:   "define",
	Short: "Create a new empty registry definition",
	Long:  `Create a new empty registry definition`,
	Run: func(cmd *cobra.Command, args []string) {
		isJson, _ := strconv.ParseBool(rootCmd.Flag("json").Value.String())
		registryDefine(env.DefaultConfigEnv(isJson))
	},
}

func registryDefine(e *env.ConfigEnv) {
	s, err := registry.Define(e)
	if err != nil {
		_, _ = fmt.Fprintf(e.Writer(), "failed to define registry: %v\n", err)
		nonZeroExit(1)
	}
	_, _ = fmt.Fprintf(e.Writer(), s)
}

func init() {
	registryCmd.AddCommand(defineCmd)
}
