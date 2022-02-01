package cli

import (
	"encoding/json"
	"fmt"
	"github.com/RomanosTrechlis/retemp/config"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"os"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Displays the configuration file",
	Long:  `Displays the content of the configuration file if it exists`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.LoadConfig()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
			os.Exit(1)
		}

		if d, _ := cmd.Flags().GetBool("dump"); d {
			spew.Dump(c)
			os.Exit(0)
		}
		s, _ := json.MarshalIndent(c, "", "  ")
		fmt.Println(string(s))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolP("dump", "d", false, "Dump configuration file")
}
