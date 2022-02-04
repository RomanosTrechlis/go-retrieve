package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"
)

func TestUpdateConfig(t *testing.T) {
	pwd, _ := os.Getwd()
	path := filepath.Join("..", "cli", "data")
	e := env.New(pwd, path, "config.json", os.Stdout)
	c, err := config.LoadConfig(e)
	if err != nil {
		t.Errorf("failed to execute test, didn't find the config file at %s: %v", e.ConfigFilePath(), err)
	}
	err = config.UpdateConfig(e, c)
	if err != nil {
		t.Errorf("failed to update config file: %v", err)
	}
}
