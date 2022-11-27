package cli

import (
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/dl"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/template"
	"github.com/RomanosTrechlis/go-retrieve/util"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RegisterGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Finds the first instance of a template and downloads it locally",
		Long:  `Finds the first instance of a template and downloads it locally`,
		Run:   get(),
	}
}

func RegisterGetListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Displays information on template definitions",
		Long: `Displays information on template definitions.

Given a provided template definition name, it displays
that definition including the template files.`,
		Run: list(),
	}
}

func get() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// default show template list
		if len(args) == 0 {
			e := env.DefaultConfigEnv(false)
			executeTemplateList(e, false, false, args)
			return
		}

		dir, _ := cmd.Flags().GetString("destination")
		flat, _ := cmd.Flags().GetBool("flat")
		executeTemplate(env.DefaultConfigEnv(false), dir, args[0], flat)
	}
}

func list() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		e := env.DefaultConfigEnv(false)
		d, _ := cmd.Flags().GetBool("dump")
		executeTemplateList(e, d, false, args)
	}
}

func executeTemplateList(e *env.ConfigEnv, dump, isJson bool, args []string) {
	if len(args) == 0 {
		err := template.List(e)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to list template definitions: %v", err)
			NonZeroExit(1)
		}
		return
	}

	t, _, err := template.Definition(e, args[0])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to retrieve template definition: %v\n", err)
		NonZeroExit(1)
	}

	if dump {
		spew.Fdump(e.Writer(), t)
		return
	}
	s, _ := util.MarshalIndent(t, isJson)
	_, _ = fmt.Fprintln(e.Writer(), string(s))
}

func executeTemplate(e *env.ConfigEnv, destination, templateName string, isFlat bool) {
	temp, token, err := template.Definition(e, templateName)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to retrieve template definition: %v\n", err)
		NonZeroExit(1)
	}

	// todo: check cyclical dependencies
	if len(temp.Dependencies) != 0 {
		for _, tn := range temp.Dependencies {
			executeTemplate(e, destination, tn, isFlat)
		}
	}

	destinationDir, err := createDirIfRequired(temp.Destination, destination, temp.Name, isFlat)
	if err != nil {
		_, _ = fmt.Fprintf(e.Writer(), "failed to create directory: %v\n", err)
		NonZeroExit(1)
	}

	err = writeFiles(temp.Files, temp.Url, token, destinationDir, isFlat)
	if err != nil {
		_, _ = fmt.Fprintf(e.Writer(), "failed to write file '%s': %v\n", temp.Name, err)
		NonZeroExit(1)
	}

	err = executeCommandsIfRequired(temp.Commands)
	if err != nil {
		_, _ = fmt.Fprintf(e.Writer(), "failed to execute after retrieval command '%s': %v\n",
			temp.Commands, err)
		NonZeroExit(1)
	}
}

func writeFiles(files []string, baseURL, token, destination string, isFlat bool) error {
	for _, f := range files {
		b, err := dl.Download(baseURL+f, token)
		if err != nil {
			return err
		}

		// todo: add warning when overriding
		filename := filename(destination, f, isFlat)
		err = os.WriteFile(filename, b, 0755)
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

// createDirIfRequired creates the destination directory,
// if there is not a flat flag, using the destination registered
// in the configuration, the destination passed as a flag
// and the name of the template
func createDirIfRequired(templateDir, flagDir, tempName string, isFlat bool) (string, error) {
	if isFlat {
		return "", nil
	}
	d := destinationDir(templateDir, flagDir, tempName)
	return d, os.MkdirAll(d, 0755)
}

func hasDir(destination string, isFlat bool) bool {
	return destination != "" || !isFlat
}

func filename(destination, filename string, isFlat bool) string {
	if isFlat {
		return filename
	}
	return filepath.Join(destination, filename)
}

func dir(destination, templateName string) string {
	if destination == "" {
		return templateName
	}
	return filepath.Join(destination, templateName)
}

func destinationDir(templateDir, flagDir, templateName string) string {
	d := templateDir
	if templateDir == "" {
		d = templateName
	}
	if flagDir != "" {
		d = filepath.Join(flagDir, d)
	}
	return d
}
