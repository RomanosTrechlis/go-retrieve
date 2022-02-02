package config

import (
	"fmt"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"os"
)

func Destroy(e *env.ConfigEnv) error {
	configFile := e.ConfigFilePath()
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return fmt.Errorf("configuration file doesn't exist, 'destroy' option is not valid")
	}

	return os.RemoveAll(configFile)
}
