package config

import (
	"encoding/json"
	"fmt"
	"github.com/RomanosTrechlis/retemp/util"
	"io/ioutil"
	"os"
	"path"
)

func ReInit(filename string, overwrite bool) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file '%s': %v", filename, err)
	}

	var c *Configuration
	err = json.Unmarshal(b, c)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config file: %v", err)
	}

	configFile := path.Join(util.ConfigPath(), "config.json")
	if _, err := os.Stat(configFile); err == nil {
		if !overwrite {
			return fmt.Errorf("configuration file already exists, 'init' option is not valid")
		}
		err := os.Remove(configFile)
		if err != nil {
			return fmt.Errorf("failed to remove existing config file: %v", err)
		}
	}

	err = util.WriteFile(util.ConfigPath(), "config.json", c)
	if err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}
	return nil
}

func Init() error {
	dirname, _ := os.UserHomeDir()
	configPath := path.Join(dirname, ".retemp")
	filename := "config.json"
	configFile := path.Join(configPath, filename)
	if _, err := os.Stat(configFile); err == nil {
		return fmt.Errorf("configuration file already exists, 'init' option is not valid")
	}

	fmt.Println("Creating configuration...")
	profiles, err := createProfiles()
	if err != nil {
		return err
	}

	configuration := &Configuration{nil, profiles}

	return util.WriteFile(configPath, filename, &configuration)
}

func createProfiles() ([]*ConfigurationProfile, error) {
	profiles := make([]*ConfigurationProfile, 0)
	for true {
		name := util.Scan("Please insert a name for the profile: ")
		if name == "" {
			break
		}

		sources, err := createSources()
		if err != nil {
			return nil, fmt.Errorf("failed to create sources for profile '%s': %v", name, err)
		}

		if len(sources) == 0 {
			fmt.Printf("Profile '%s' is empty!", name)
		}

		profiles = append(profiles, &ConfigurationProfile{name, sources})

		if !more("profile") {
			break
		}
	}
	return profiles, nil
}

func createSources() ([]*ConfigurationSource, error) {
	sources := make([]*ConfigurationSource, 0)
	for true {
		name := util.Scan("Please insert a name for the source: ")
		if name == "" {
			break
		}

		url := util.Scan("Please insert the source URL: ")
		token := util.Scan("Please insert the environment variable for the token, if required: ")

		sources = append(sources, &ConfigurationSource{name, url, token})

		if !more("source") {
			break
		}
	}
	return sources, nil
}

func more(section string) bool {
	more := util.Scan(fmt.Sprintf("Whould you like to add another %s? [Y]es, [N]o: ", section))
	if more == "YES" || more == "yes" || more == "Y" || more == "y" {
		return true
	}
	return false
}
