package env

import (
	"io"
	"os"
	"path/filepath"
)

type ConfigEnv struct {
	HomeDir    string
	ConfigDir  string
	ConfigName string
	w          io.Writer
}

func DefaultConfigEnv() *ConfigEnv {
	home, _ := os.UserHomeDir()
	return &ConfigEnv{home, ".go-retrieve", "config.json", os.Stdout}
}

func New(homeDir, configDir, configName string, w io.Writer) *ConfigEnv {
	return &ConfigEnv{homeDir, configDir, configName, w}
}

func (c *ConfigEnv) ConfigPath() string {
	if c.HomeDir == "" {
		return c.ConfigDir
	}
	return filepath.Join(c.HomeDir, c.ConfigDir)
}

func (c *ConfigEnv) ConfigFilePath() string {
	return filepath.Join(c.HomeDir, c.ConfigDir, c.ConfigName)
}

func (c *ConfigEnv) Writer() io.Writer {
	return c.w
}
