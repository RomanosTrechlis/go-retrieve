package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/retemp/dl"
	"github.com/RomanosTrechlis/retemp/template"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Finds the first instance of a template and downloads it locally",
	Long:  `Finds the first instance of a template and downloads it locally`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_, _ = fmt.Fprintf(os.Stderr, "please provide a template name\n")
			os.Exit(1)
		}

		t, token, err := template.TemplateDefinition(args[0])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to retrieve template definition: %v\n", err)
			os.Exit(1)
		}

		dir, _ := cmd.Flags().GetString("destination")
		if dir == "" {
			dir = t.Name
		} else {
			dir = filepath.Join(dir, t.Name)
		}

		err = os.MkdirAll(dir, 0755)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to create directory: %v\n", err)
			os.Exit(1)
		}

		for _, f := range t.Files {
			b, err := dl.Download(t.Url+f, token)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to download file '%s': %v\n", f, err)
				os.Exit(1)
			}

			filename := filepath.Join(dir, f)
			err = ioutil.WriteFile(filename, b, 0755)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to write file '%s': %v\n", filename, err)
				os.Exit(1)
			}
		}

		if t.Commands != "" {
			cmd := exec.Command(t.Commands)
			cc := strings.Split(t.Commands, " ")
			if len(cc) > 1 {
				cmd = exec.Command(cc[0], cc[1:len(cc)]...)
			}

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to execute after retrieval command '%s': %v\n", t.Commands, err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// templateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	templateCmd.Flags().StringP("destination", "d", "", "Download template to specific destination")
}
