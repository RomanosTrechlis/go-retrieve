package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"os"

	"github.com/spf13/cobra"
)

// profileListCmd represents the list command
var profileListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the available profiles",
	Long:  `Lists all the available profiles`,
	Run: func(cmd *cobra.Command, args []string) {
		executeProfileList(env.DefaultConfigEnv())
	},
}

func executeProfileList(e *env.ConfigEnv) {
	c, err := config.LoadConfig(e)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		nonZeroExit(1)
	}

	for _, p := range c.Profiles {
		_, _ = fmt.Fprintln(e.Writer(), p.Name)
	}
}

func init() {
	profileCmd.AddCommand(profileListCmd)
}
