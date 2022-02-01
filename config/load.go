package config

import (
	"encoding/json"
	"fmt"
	"github.com/RomanosTrechlis/retemp/util"
	"io/ioutil"
)

func LoadConfig() (*Configuration, error) {
	configFile := util.ConfigFilePath()
	if !util.IsExists(configFile) {
		return nil, fmt.Errorf("coundn't find configuration file, run 'init' command first")
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
