package env

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

type ConfigEnv struct {
	HomeDir    string
	ConfigDir  string
	ConfigName string
	w          io.Writer
	j          bool
}

func DefaultConfigEnv(isJson bool) *ConfigEnv {
	home, _ := os.UserHomeDir()
	configName := "config.yml"
	return &ConfigEnv{home, ".rt", configName, os.Stdout, isJson}
}

func New(homeDir, configDir, configName string, w io.Writer) *ConfigEnv {
	return &ConfigEnv{
		homeDir,
		configDir,
		configName,
		w,
		strings.HasSuffix(configName, "json")}
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

func (c *ConfigEnv) IsJson() bool {
	return c.j
}

func (c *ConfigEnv) Suffix() string {
	if c.IsJson() {
		return "json"
	}
	return "yml"
}
