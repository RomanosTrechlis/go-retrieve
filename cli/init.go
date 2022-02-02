package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/spf13/cobra"
	"os"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Execute wizard to initialize the go-retrieve",
	Long: `Execute wizard to initialize the go-retrieve.

Follow the steps on the terminal to create profiles and
sources. You can reinstate an existing configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		e := env.DefaultConfigEnv()
		f, _ := cmd.Flags().GetString("filename")
		o, _ := cmd.Flags().GetBool("overwrite")
		executeInit(e, f, o)
	},
}

func executeInit(e *env.ConfigEnv, filename string, overwrite bool) {
	if filename != "" {
		err := config.ReInit(e, filename, overwrite)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to re-initialize: %v\n", err)
			nonZeroExit(1)
		}
	}

	err := config.Init(e)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to initialize: %v\n", err)
		nonZeroExit(1)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("filename", "f", "", "Initialize from existing file")
	initCmd.Flags().BoolP("overwrite", "o", false, "Overwrite existing config when initializing from existing file")
}
