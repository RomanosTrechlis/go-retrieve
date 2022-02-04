package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"

	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Copies the configuration file to current directory",
	Long:  `Copies the configuration file to current directory`,
	Run: func(cmd *cobra.Command, args []string) {
		f, _ := cmd.Flags().GetString("filename")
		executeBackup(env.DefaultConfigEnv(), f)
	},
}

func executeBackup(e *env.ConfigEnv, f string) {
	c, err := config.LoadConfig(e)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		nonZeroExit(1)
	}

	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to marshal config: %v\n", err)
		nonZeroExit(1)
	}

	filename := e.ConfigName
	if f != "" {
		filename = f
	}
	err = ioutil.WriteFile(filename, b, 0755)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to write config file: %v\n", err)
		nonZeroExit(1)
	}
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().StringP("filename", "f", "", "Backup to specific file")
}
