package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"

	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Displays the configuration file",
	Long:  `Displays the content of the configuration file if it exists`,
	Run: func(cmd *cobra.Command, args []string) {
		e := env.DefaultConfigEnv()
		d, _ := cmd.Flags().GetBool("dump")
		executeConfig(e, d)
	},
}

func executeConfig(e *env.ConfigEnv, dump bool) {
	c, err := config.LoadConfig(e)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		nonZeroExit(1)
		//os.Exit(1)
		//panic(fmt.Sprintf("failed to load config: %v\n", err))
	}

	if dump {
		spew.Fdump(e.Writer(), c)
		return
	}

	s, _ := json.MarshalIndent(c, "", "  ")
	_, _ = fmt.Fprintln(e.Writer(), string(s))
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolP("dump", "d", false, "Dump configuration file")
}
