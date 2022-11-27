package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/util"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"os"
)

func RegisterConfig(rootCmd *cobra.Command) *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Displays the configuration file",
		Long:  `Displays the content of the configuration file if it exists`,
		Run:   cfg(rootCmd),
	}
}

func cfg(rootCmd *cobra.Command) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		e := env.DefaultConfigEnv(false)
		d, _ := cmd.Flags().GetBool("dump")
		ExecuteConfig(e, d)
	}
}

func ExecuteConfig(e *env.ConfigEnv, dump bool) {
	c, err := config.LoadConfig(e)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		NonZeroExit(1)
		//os.Exit(1)
		//panic(fmt.Sprintf("failed to load config: %v\n", err))
	}

	if dump {
		spew.Fdump(e.Writer(), c)
		return
	}

	s, _ := util.MarshalIndent(c, false)
	_, _ = fmt.Fprintln(e.Writer(), string(s))
}
