package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Execute wizard to initialize the rt configuration",
	Long: `Execute wizard to initialize the rt configuration.

Follow the steps on the terminal to create profiles and
sources. You can reinstate an existing configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		isJson, _ := strconv.ParseBool(rootCmd.Flag("json").Value.String())
		e := env.DefaultConfigEnv(isJson)
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
