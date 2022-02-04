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
		flat, _ := cmd.Flags().GetBool("flat")

		executeTemplate(env.DefaultConfigEnv(), dir, args[0], flat)
	},
}

func executeTemplate(e *env.ConfigEnv, destination, templateName string, isFlat bool) {
	temp, token, err := template.Definition(e, templateName)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to retrieve template definition: %v\n", err)
		nonZeroExit(1)
	}

	// todo: check cyclical dependencies
	if len(temp.Dependencies) != 0 {
		for _, tn := range temp.Dependencies {
			executeTemplate(e, destination, tn, isFlat)
		}
	}

	err = createDirIfRequired(destination, temp.Name, isFlat)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to create directory: %v\n", err)
		nonZeroExit(1)
	}

	err = writeFiles(temp.Files, temp.Url, token, destination, temp.Name, isFlat)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to write file '%s': %v\n", filename, err)
		nonZeroExit(1)
	}

	err = executeCommandsIfRequired(temp.Commands)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to execute after retrieval command '%s': %v\n",
			temp.Commands, err)
		nonZeroExit(1)
	}
}

func writeFiles(files []string, baseURL, token, destination, tempName string, isFlat bool) error {
	for _, f := range files {
		b, err := dl.Download(baseURL+f, token)
		if err != nil {
			return err
		}

		// todo: add warning when overriding
		filename := filename(destination, tempName, f, isFlat)
		err = ioutil.WriteFile(filename, b, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func executeCommandsIfRequired(command string) error {
	if command == "" {
		return nil
	}

	cmd := exec.Command(command)
	cc := strings.Split(command, " ")
	if len(cc) > 1 {
		cmd = exec.Command(cc[0], cc[1:]...)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func createDirIfRequired(destination, tempName string, isFlat bool) error {
	if !hasDir(destination, isFlat) {
		return nil
	}
	dir := dir(destination, tempName)
	return os.MkdirAll(dir, 0755)
}

func hasDir(destination string, isFlat bool) bool {
	return destination != "" || !isFlat
}

func filename(destination, tempName, filename string, isFlat bool) string {
	if !hasDir(destination, isFlat) {
		return filename
	}
	return filepath.Join(dir(destination, tempName), filename)
}

func dir(destination, templateName string) string {
	if destination == "" {
		return templateName
	}
	return filepath.Join(destination, templateName)
}

func init() {
	rootCmd.AddCommand(templateCmd)
	templateCmd.Flags().StringP("destination", "d", "",
		"Download template to specific destination")
	templateCmd.Flags().BoolP("flat", "f", false,
		`Download files in the destination instead of the template directory. 
Note: this applies to dependencies also.`)
}
