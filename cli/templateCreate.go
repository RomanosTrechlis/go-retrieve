package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/retemp/template"
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

		// todo: upload functionality?
		e, _ := cmd.Flags().GetString("exclude")
		err := template.CreateTemplateConfig(args[0], e)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to create template: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	templateCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createCmd.Flags().StringP("exclude", "e", "", "List of directories to exclude (CSV)")
}
