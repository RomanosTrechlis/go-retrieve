package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/util"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"

	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/template"
)

// templateListCmd represents the list command
var templateListCmd = &cobra.Command{
	Use:   "list",
	Short: "Displays information on template definitions",
	Long: `Displays information on template definitions.

Given a provided template definition name, it displays
that definition including the template files.`,
	Run: func(cmd *cobra.Command, args []string) {
		j, _ := strconv.ParseBool(rootCmd.Flag("json").Value.String())
		e := env.DefaultConfigEnv(j)
		d, _ := cmd.Flags().GetBool("dump")
		executeTemplateList(e, d, j, args)
	},
}

func executeTemplateList(e *env.ConfigEnv, dump, isJson bool, args []string) {
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
	s, _ := util.MarshalIndent(t, isJson)
	_, _ = fmt.Fprintln(e.Writer(), string(s))
}

func init() {
	templateCmd.AddCommand(templateListCmd)
	templateListCmd.Flags().BoolP("dump", "d", false, "Dump profile configuration")
}
