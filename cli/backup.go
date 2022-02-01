package cli

import (
	"encoding/json"
	"fmt"
	"github.com/RomanosTrechlis/retemp/config"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Copies the configuration file to current directory",
	Long:  `Copies the configuration file to current directory`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.LoadConfig()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
			os.Exit(1)
		}

		b, err := json.MarshalIndent(c, "", "  ")
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to marshal config: %v\n", err)
			os.Exit(1)
		}

		err = ioutil.WriteFile("config.json", b, 0755)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to write config file: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
