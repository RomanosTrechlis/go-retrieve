package cli

import (
	"encoding/json"
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/template"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"os"
)

// templateListCmd represents the list command
var templateListCmd = &cobra.Command{
	Use:   "list",
	Short: "Displays information on template definitions",
	Long: `Displays information on template definitions.

Given a provided template definition name, it displays 
that definition including the template files.`,
	Run: func(cmd *cobra.Command, args []string) {
		e := env.DefaultConfigEnv()
		d, _ := cmd.Flags().GetBool("dump")
		executeTemplateList(e, d, args)
	},
}

func executeTemplateList(e *env.ConfigEnv, dump bool, args []string) {
	if len(args) == 0 {
		err := template.List(e)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to list template definitions: %v", err)
			nonZeroExit(1)
		}
		return
	}

	t, _, err := template.Definition(e, args[0])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to retrieve template definition: %v\n", err)
		nonZeroExit(1)
	}

	if dump {
		spew.Fdump(e.Writer(), t)
		return
	}
	s, _ := json.MarshalIndent(t, "", "  ")
	_, _ = fmt.Fprintln(e.Writer(), string(s))
}

func init() {
	templateCmd.AddCommand(templateListCmd)
	templateListCmd.Flags().BoolP("dump", "d", false, "Dump profile configuration")
}
