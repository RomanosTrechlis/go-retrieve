package template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/registry"
	"github.com/RomanosTrechlis/go-retrieve/util"
)

// Definition finds a registered template with the required token if exists on the local registry
func Definition(e *env.ConfigEnv, templateName string) (*registry.RegisteredTemplate, string, error) {
	// todo: if template name is not found locally run update to include now template definitions
	c, err := config.LoadConfig(e)
	if err != nil {
		return nil, "", err
	}

	for _, s := range c.Active.Sources {
		templateFile := filepath.Join(e.ConfigPath(), fmt.Sprintf("%s.%s", s.Name, e.Suffix()))
		r, err := LoadRegistryFile(templateFile, e.IsJson())
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

// List prints active or all the available templates defined in the configuration directory
func List(e *env.ConfigEnv, all bool) error {
	if all {
		return filepath.Walk(e.ConfigPath(), templateWalk(e.ConfigName))
	}

	c, err := config.LoadConfig(e)
	if err != nil {
		return err
	}

	for _, s := range c.Active.Sources {
		p := filepath.Join(e.ConfigPath(), fmt.Sprintf("%s.%s", s.Name, e.Suffix()))
		r, err := LoadRegistryFile(p, strings.HasSuffix(p, "json"))
		if err != nil {
			return err
		}
		templateSummaryPrint(r)
	}
	return nil
}

// LoadRegistryFile finds a specific registry and returns it
func LoadRegistryFile(file string, isJson bool) (*registry.Registry, error) {
	b, err := util.LoadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to load config file: %v", err)
	}

	reg, err := registry.Unmarshal(b, isJson)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %v", err)
	}
	return reg, nil
}

func templateWalk(configName string) func(innerPath string, info os.FileInfo, err error) error {
	return func(innerPath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if info.Name() == configName {
			return nil
		}

		r, err := LoadRegistryFile(innerPath, strings.HasSuffix(innerPath, "json"))
		if err != nil {
			return fmt.Errorf("failed to load registry file '%s': %v", innerPath, err)
		}

		fmt.Printf("---------------- %s ----------------\n", r.Name)
		templateSummaryPrint(r)
		return nil
	}
}

func templateSummaryPrint(r *registry.Registry) {
	for i, t := range r.Templates {
		fmt.Printf("%d/ %s", i+1, t.Name)
		if t.Description != "" {
			fmt.Printf(": %s", t.Description)
		}
		fmt.Println()
	}
}
