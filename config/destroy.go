package config

import (
	"fmt"
	"os"
	"path"
)

func Destroy() error {
	dirname, _ := os.UserHomeDir()
	configPath := path.Join(dirname, ".retemp")
	configFile := path.Join(configPath, "config.json")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return fmt.Errorf("configuration file doesn't exist, 'destroy' option is not valid")
	}

	return os.RemoveAll(configFile)
}
