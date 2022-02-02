package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/template"
	"github.com/spf13/cobra"
	"os"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a template definition given a path",
	Long:  `Creates a template definition given a path`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_, _ = fmt.Fprintf(os.Stderr, "please provide at least one directory or file\n")
			os.Exit(1)
		}

		ex, _ := cmd.Flags().GetString("exclude")
		executeCreate(args[0], ex)
	},
}

func executeCreate(templateName, ex string) {
	// todo: upload functionality?
	err := template.CreateTemplateConfig(templateName, ex)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to create template: %v", err)
		os.Exit(1)
	}
}

func init() {
	templateCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("exclude", "e", "", "List of directories to exclude (CSV)")
}
