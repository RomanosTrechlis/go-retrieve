package template

import (
	"encoding/json"
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/registry"
	"github.com/RomanosTrechlis/go-retrieve/util"
	"io/ioutil"
	"os"
	"path/filepath"
)

func TemplateDefinition(e *env.ConfigEnv, templateName string) (*registry.RegisteredTemplate, string, error) {
	// todo: if template name is not found locally run update to include now template definitions
	c, err := config.LoadConfig(e)
	if err != nil {
		return nil, "", err
	}

	for _, s := range c.Active.Sources {
		templateFile := filepath.Join(e.ConfigPath(), s.Name+".json")
		r, err := LoadRegistryFile(templateFile)
		if err != nil {
			return nil, "", fmt.Errorf("failed to load registry file '%s': %v", templateFile, err)
		}

		for _, t := range r.Templates {
			if t.Name == templateName {
				return &t, s.Token, nil
			}
		}
	}
	return nil, "", fmt.Errorf("template with name '%s' doesn't exist", templateName)
}

func List(e *env.ConfigEnv) error {
	files := make([]string, 0)
	err := filepath.Walk(e.ConfigPath(), templateWalk(files, e.ConfigName))

	for _, f := range files {
		r, err := LoadRegistryFile(f)
		if err != nil {
			return fmt.Errorf("failed to load registry file '%s': %v", f, err)
		}

		for i, t := range r.Templates {
			fmt.Printf("%d/\n", i+1)
			fmt.Printf("Name: \t\t%s\nURL: \t\t%s\nCommands: \t%s\nDependencies: \t%v\n",
				t.Name, t.Url, t.Commands, t.Dependencies)
		}
	}
	return err
}

func LoadRegistryFile(file string) (*registry.Registry, error) {
	if !util.IsExists(file) {
		return nil, fmt.Errorf("couldn't find configuration file, run 'init' command first")
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to load config file: %v", err)
	}

	var reg *registry.Registry
	err = json.Unmarshal(b, &reg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %v", err)
	}
	return reg, nil
}

func templateWalk(files []string, configName string) func(innerPath string, info os.FileInfo, err error) error {
	return func(innerPath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if info.Name() == configName {
			return nil
		}

		files = append(files, innerPath)
		return nil
	}
}
