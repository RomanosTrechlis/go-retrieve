package main

import (
	"github.com/RomanosTrechlis/go-retrieve/cli"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rta",
	Short: "Retrieve templates administration and maintenance cli",
	Long: `Retrieve templates administration and maintenance cli.

rta handles the creation of the required configuration files
that make the 'rt' work as expected.
rta creates the initial configuration file, as well as the 
accompanying registry configuration files.`,
	// show config file
	Run: func(cmd *cobra.Command, args []string) {
		e := env.DefaultConfigEnv(false)
		cli.ExecuteConfig(e, false)
	},
}

// Execute adds all child cmd to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		cli.NonZeroExit(1)
	}
}

func init() {
}
