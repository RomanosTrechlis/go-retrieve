package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/retemp/config"
	"github.com/spf13/cobra"
	"os"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Execute wizard to initialize the retemp",
	Long: `Execute wizard to initialize the retemp.

Follow the steps on the terminal to create profiles and
sources. You can reinstate an existing configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("filename")
		if filename != "" {
			overwrite, _ := cmd.Flags().GetBool("overwrite")
			err := config.ReInit(filename, overwrite)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to re-initialize: %v\n", err)
				os.Exit(1)
			}
		}

		err := config.Init()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to initialize: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("filename", "f", "", "Initialize from existing file")
	initCmd.Flags().BoolP("overwrite", "o", false, "Overwrite existing config when initializing from existing file")
}
