package cli

import (
	"encoding/json"
	"fmt"
	"github.com/RomanosTrechlis/retemp/template"
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
		if len(args) == 0 {
			err := template.List()
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to list template definitions: %v", err)
				os.Exit(1)
			}
			return
		}

		t, _, err := template.TemplateDefinition(args[0])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to retrieve template definition: %v\n", err)
			os.Exit(1)
		}

		if d, _ := cmd.Flags().GetBool("dump"); d {
			spew.Dump(t)
			os.Exit(0)
		}
		s, _ := json.MarshalIndent(t, "", "  ")
		fmt.Println(string(s))
	},
}

func init() {
	templateCmd.AddCommand(templateListCmd)
	templateListCmd.Flags().BoolP("dump", "d", false, "Dump profile configuration")
}
