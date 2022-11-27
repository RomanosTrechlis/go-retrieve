package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/registry"
	"github.com/RomanosTrechlis/go-retrieve/template"
	"github.com/spf13/cobra"
	"os"
)

func RegisterRegistryCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "registry",
		Short: "Not yet implemented",
		Long:  `Not yet implemented`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}

func RegisterRegistryAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Creates a template definition given a path",
		Long:  `Creates a template definition given a path`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				_, _ = fmt.Fprintf(os.Stderr, "please provide at least one directory or file\n")
				NonZeroExit(1)
			}

			ex, _ := cmd.Flags().GetString("exclude")
			executeCreate(args[0], ex, false)
		},
	}
}

func RegisterRegistryDefineCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "define",
		Short: "Create a new empty registry definition",
		Long:  `Create a new empty registry definition`,
		Run: func(cmd *cobra.Command, args []string) {
			registryDefine(env.DefaultConfigEnv(false))
		},
	}
}

func registryDefine(e *env.ConfigEnv) {
	s, err := registry.Define(e)
	if err != nil {
		_, _ = fmt.Fprintf(e.Writer(), "failed to define registry: %v\n", err)
		NonZeroExit(1)
	}
	_, _ = fmt.Fprintf(e.Writer(), s)
}

func executeCreate(templateName, ex string, isJson bool) {
	// todo: upload functionality?
	err := template.CreateTemplateConfig(templateName, ex, isJson)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to create template: %v", err)
		NonZeroExit(1)
	}
}
