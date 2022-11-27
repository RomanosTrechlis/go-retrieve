package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/spf13/cobra"
	"os"
)

func RegisterNukeCmd(rootCmd *cobra.Command) *cobra.Command {
	return &cobra.Command{
		Use:   "nuke",
		Short: "Delete configuration folder with the corresponding files",
		Long:  `Delete configuration folder with the corresponding files`,
		Run:   nuke(rootCmd),
	}
}

func nuke(rootCmd *cobra.Command) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		executeNuke(env.DefaultConfigEnv(false))
	}
}

func executeNuke(e *env.ConfigEnv) {
	err := config.Destroy(e)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to destroy environment: %v\n", err)
		NonZeroExit(1)
	}
}
