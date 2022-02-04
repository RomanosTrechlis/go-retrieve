package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/util"
)

func LoadConfig(e *env.ConfigEnv) (*Configuration, error) {
	configFile := e.ConfigFilePath()
	if !util.IsExists(configFile) {
		return nil, fmt.Errorf("couldn't find configuration file, run 'init' command first")
	}

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load config file: %v", err)
	}

	var configuration *Configuration
	err = json.Unmarshal(b, &configuration)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %v", err)
	}
	return configuration, nil
}
