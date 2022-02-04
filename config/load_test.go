package config_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/env"
)

func TestLoadConfig(t *testing.T) {
	pwd, _ := os.Getwd()
	var output bytes.Buffer
	path := filepath.Join("..", "cli", "data")
	e := env.New(pwd, path, "config.json", &output)

	_, err := config.LoadConfig(e)
	if err != nil {
		t.Errorf("failed to load config: expected no error, got '%v'", err)
	}

	e = env.New(pwd, path, "dummy.json", &output)
	c, err := config.LoadConfig(e)
	if err == nil {
		t.Errorf("dummy config file: expected unmarshaling error, got '%v'", c)
	}
}
