package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/util"
	"github.com/spf13/cobra"
	"os"
)

func RegisterBackupCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "backup",
		Short: "Copies the configuration file to current directory",
		Long:  `Copies the configuration file to current directory`,
		Run:   backup(),
	}
}

func backup() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		f, _ := cmd.Flags().GetString("filename")
		executeBackup(env.DefaultConfigEnv(false), f)
	}
}

func executeBackup(e *env.ConfigEnv, f string) {
	c, err := config.LoadConfig(e)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		NonZeroExit(1)
	}

	b, err := util.MarshalIndent(c, false)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to marshal config: %v\n", err)
		NonZeroExit(1)
	}

	filename := e.ConfigName
	if f != "" {
		filename = f
	}
	err = os.WriteFile(filename, b, 0755)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to write config file: %v\n", err)
		NonZeroExit(1)
	}
}
