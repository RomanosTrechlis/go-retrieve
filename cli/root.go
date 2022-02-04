package cli

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-retrieve",
	Short: "Retrieves templates from remote or local locations",
	Long: `Retrieves templates from remote or local locations and
copies them to current directory.

Used mainly for bootstraping software projects,
adding modules or components to an existing project,
or just copying resources from remote locations.

First, run the 'init' command to create the required
configuration files. Follow the steps of the wizard.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: rootFunc,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		nonZeroExit(1)
	}
}

func init() {
}
