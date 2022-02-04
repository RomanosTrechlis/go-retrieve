package config

import (
	"fmt"
	"os"

	"github.com/RomanosTrechlis/go-retrieve/env"
)

func Destroy(e *env.ConfigEnv) error {
	configFile := e.ConfigFilePath()
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return fmt.Errorf("configuration file doesn't exist, 'destroy' option is not valid")
	}

	return os.RemoveAll(e.ConfigPath())
}
