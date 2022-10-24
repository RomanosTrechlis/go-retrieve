package config

import (
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/util"
	"os"
)

func LoadConfig(e *env.ConfigEnv) (*Configuration, error) {
	configFile := e.ConfigFilePath()
	if !util.IsExists(configFile) {
		return nil, fmt.Errorf("couldn't find configuration file, run 'init' command first")
	}

	b, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load config file: %v", err)
	}

	return Unmarshal(b, e.IsJson())
}
