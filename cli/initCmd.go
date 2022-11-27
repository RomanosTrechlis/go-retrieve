package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func RegisterInitCmd(rootCmd *cobra.Command) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Execute wizard to initialize the rt configuration",
		Long: `Execute wizard to initialize the rt configuration.

Follow the steps on the terminal to create profiles and
sources. You can reinstate an existing configuration file.`,
		Run: initialize(rootCmd),
	}
}

func RegisterSimpleInitCmd(rootCmd *cobra.Command) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Copy configuration file to configuration path",
		Long:  `Copy configuration file to configuration path`,
		Run:   initializeSimple(),
	}
}

func initializeSimple() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_, _ = fmt.Fprintf(os.Stderr,
				"Please provide the configuration 'config.yml' file as well as the registy files\n")
			NonZeroExit(1)
		}

		e := env.DefaultConfigEnv(false)
		err := os.MkdirAll(filepath.Join(e.HomeDir, e.ConfigDir), os.ModePerm)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to create config path: %v\n", err)
			NonZeroExit(1)
		}

		for i := 0; i < len(args); i++ {
			err = copyFile(e, args[i])
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to copy file '%s': %v\n", args[i], err)
				_ = os.RemoveAll(e.ConfigPath())
				NonZeroExit(1)
			}
		}
	}
}

func copyFile(e *env.ConfigEnv, filename string) error {
	b, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file '%s': %v\n", filename, err)
	}
	err = os.WriteFile(filepath.Join(e.ConfigPath(), filename), b, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write file '%s' to configuration folder: %v\n", filename, err)
	}
	return nil
}

func initialize(rootCmd *cobra.Command) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		e := env.DefaultConfigEnv(false)
		f, _ := cmd.Flags().GetString("filename")
		o, _ := cmd.Flags().GetBool("overwrite")
		executeInit(e, f, o)
	}
}

func executeInit(e *env.ConfigEnv, filename string, overwrite bool) {
	if filename != "" {
		err := config.ReInit(e, filename, overwrite)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to re-initialize: %v\n", err)
			NonZeroExit(1)
		}
	}

	err := config.Init(e)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to initialize: %v\n", err)
		NonZeroExit(1)
	}
}
