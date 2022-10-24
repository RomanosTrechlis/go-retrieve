package config_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"
)

func TestLoadJSONConfig(t *testing.T) {
	e := createConfig("config.json")
	_, err := config.LoadConfig(e)
	if err != nil {
		t.Errorf("failed to load config: expected no error, got '%v'", err)
	}

	e = createConfig("dummy.json")
	c, err := config.LoadConfig(e)
	if err == nil {
		t.Errorf("dummy config file: expected unmarshaling error, got '%v'", c)
	}
}

func TestLoadYMLConfig(t *testing.T) {
	e := createConfig("config.yml")
	_, err := config.LoadConfig(e)
	if err != nil {
		t.Errorf("failed to load config: expected no error, got '%v'", err)
	}

	e = createConfig("dummy.yml")
	c, err := config.LoadConfig(e)
	if err == nil {
		t.Errorf("dummy config file: expected unmarshaling error, got '%v'", c)
	}
}

func createConfig(name string) *env.ConfigEnv {
	pwd, _ := os.Getwd()
	var output bytes.Buffer
	path := filepath.Join("..", "cli", "data")
	return env.New(pwd, path, name, &output)
}
