package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"

	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/util"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Display the configuration of specific profile",
	Long:  `Display the configuration of specific profile`,
	Run: func(cmd *cobra.Command, args []string) {
		dump, _ := cmd.Flags().GetBool("dump")
		executeProfileInspect(env.DefaultConfigEnv(), args, dump)
	},
}

func executeProfileInspect(e *env.ConfigEnv, args []string, dump bool) {
	if len(args) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "provide at least one profile to inspect\n")
		nonZeroExit(1)
	}

	c, err := config.LoadConfig(e)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		nonZeroExit(1)
	}

	isFirst := true
	for _, p := range c.Profiles {
		if !util.Contains(args, p.Name) {
			continue
		}

		if !isFirst {
			isFirst = false
			_, _ = fmt.Fprintf(e.Writer(), "------------- %s -----------\n", p.Name)
		}

		if dump {
			spew.Fdump(e.Writer(), p)
			return
		}
		s, _ := json.MarshalIndent(p, "", "  ")
		_, _ = fmt.Fprintln(e.Writer(), string(s))
	}
}

func init() {
	profileCmd.AddCommand(inspectCmd)
	inspectCmd.Flags().BoolP("dump", "d", false, "Dump profile configuration")
}
