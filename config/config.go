package config

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Active   *ConfigurationProfile   `json:"active" yaml:"active"`
	Profiles []*ConfigurationProfile `json:"profiles" yaml:"profiles"`
}

type ConfigurationProfile struct {
	Name    string                 `json:"name" yaml:"name"`
	Sources []*ConfigurationSource `json:"sources" yaml:"sources"`
}

type ConfigurationSource struct {
	Name  string `json:"name" yaml:"name"`
	Url   string `json:"url" yaml:"url"`
	Token string `json:"token" yaml:"token"`
}

func Unmarshal(b []byte, isJson bool) (*Configuration, error) {
	var c *Configuration
	var err error
	if isJson {
		err = json.Unmarshal(b, &c)
	} else {
		err = yaml.Unmarshal(b, &c)
	}

	return c, err
}
