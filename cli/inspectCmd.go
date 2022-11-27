package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/template"
	"github.com/RomanosTrechlis/go-retrieve/util"
	"github.com/spf13/cobra"
	"os"
)

func RegisterInspectTemplateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "inspect",
		Short: "Inspects the configuration of the given template",
		Long:  `Inspects the configuration of the given template`,
		Run:   inspect(),
	}
}

func inspect() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_, _ = fmt.Fprintf(os.Stderr, "Please provide a valid template name\n")
			NonZeroExit(1)
		}

		e := env.DefaultConfigEnv(false)
		t, _, err := template.Definition(e, args[0])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to retrieve template definition: %v\n", err)
			NonZeroExit(1)
		}
		s, _ := util.MarshalIndent(t, false)
		_, _ = fmt.Fprintln(e.Writer(), string(s))
	}
}
