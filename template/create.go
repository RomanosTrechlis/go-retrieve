package template

import (
	"encoding/json"
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/registry"
	"github.com/RomanosTrechlis/go-retrieve/util"
	"os"
	"path/filepath"
	"strings"
)

// CreateTemplateConfig prints the section of the configuration necessary to define a template
// given a specific path and, if necessary, paths to exclude
func CreateTemplateConfig(path string, excludeCSV string) error {
	name := util.Scan("What is the name of the template?")
	cmd := util.Scan("What command should be executed after retrieving the template?")
	url := util.Scan("What is the base url?")
	deps := util.Scan("Please enter the names of its dependencies, if applicable")

	exclude := util.CSVToArray(excludeCSV)
	files, err := walk(path, exclude)
	if err != nil {
		return fmt.Errorf("failed to walk path '%s': %v", path, err)
	}

	dependencies := util.CSVToArray(deps)
	template := &registry.RegisteredTemplate{
		Name:         name,
		Url:          url,
		Dependencies: dependencies,
		Commands:     cmd,
		Files:        files,
	}

	b, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to unmarshal template file: %v", err)
	}

	fmt.Println(string(b))
	fmt.Println("Add the above configuration to your registry template configuration")
	return nil
}

func walk(path string, exclude []string) ([]string, error) {
	if !util.IsExists(path) {
		return nil, fmt.Errorf("provided path '%s' doesn't exist", path)
	}

	var files []string
	err := filepath.Walk(path, func(innerPath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if innerPath != path {
			innerPath = strings.Replace(innerPath, path, "", 1)
		} else {
			_, innerPath = filepath.Split(innerPath)
		}

		for _, e := range exclude {
			if strings.HasPrefix(innerPath, e) {
				return nil
			}
		}

		files = append(files, innerPath)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
