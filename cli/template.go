package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/RomanosTrechlis/go-retrieve/dl"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/template"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Finds the first instance of a template and downloads it locally",
	Long:  `Finds the first instance of a template and downloads it locally`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_, _ = fmt.Fprintf(os.Stderr, "please provide a template name\n")
			nonZeroExit(1)
		}
		dir, _ := cmd.Flags().GetString("destination")

		executeTemplate(env.DefaultConfigEnv(), dir, args[0])
	},
}

func executeTemplate(e *env.ConfigEnv, destination, templateName string) {
	temp, token, err := template.Definition(e, templateName)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to retrieve template definition: %v\n", err)
		nonZeroExit(1)
	}
	dir := ""
	if destination == "" {
		dir = temp.Name
	} else {
		dir = filepath.Join(destination, temp.Name)
	}

	// todo: check cyclical dependencies
	if len(temp.Dependencies) != 0 {
		for _, tn := range temp.Dependencies {
			executeTemplate(e, destination, tn)
		}
	}

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to create directory: %v\n", err)
		nonZeroExit(1)
	}

	for _, f := range temp.Files {
		b, err := dl.Download(temp.Url+f, token)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to download file '%s': %v\n", f, err)
			nonZeroExit(1)
		}

		filename := filepath.Join(dir, f)
		err = ioutil.WriteFile(filename, b, 0755)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to write file '%s': %v\n", filename, err)
			nonZeroExit(1)
		}
	}

	if temp.Commands != "" {
		cmd := exec.Command(temp.Commands)
		cc := strings.Split(temp.Commands, " ")
		if len(cc) > 1 {
			cmd = exec.Command(cc[0], cc[1:len(cc)]...)
		}

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to execute after retrieval command '%s': %v\n",
				temp.Commands, err)
			nonZeroExit(1)
		}
	}
}

func init() {
	rootCmd.AddCommand(templateCmd)
	templateCmd.Flags().StringP("destination", "d", "", "Download template to specific destination")
}
