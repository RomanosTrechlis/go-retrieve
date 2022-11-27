package main

import (
	"github.com/RomanosTrechlis/go-retrieve/cli"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rt",
	Short: "Retrieves templates from remote or local locations",
	Long: `Retrieves templates from remote or local locations and
copies them to current directory.

Used mainly for bootstraping software projects,
adding modules or components to an existing project,
or just copying resources from remote locations.

First, run the 'init' command to create the required
configuration files. Follow the steps of the wizard.`,
	// show config file
	Run: func(cmd *cobra.Command, args []string) {
		e := env.DefaultConfigEnv(false)
		cli.ExecuteConfigPrint(e, false)
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
