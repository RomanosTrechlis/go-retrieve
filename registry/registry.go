package registry

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

//go:generate stringer -type=Protocol
type Protocol int

const (
	FTP Protocol = iota
	HTTP
	LOCAL
)

//var Protocols = []string{"ftp", "http", "local"}

type Registry struct {
	Name      string               `json:"name" yaml:"name"`
	Protocol  Protocol             `json:"protocol" yaml:"protocol"`
	Templates []RegisteredTemplate `json:"templates" yaml:"templates"`
}

type RegisteredTemplate struct {
	Name         string   `json:"name" yaml:"name"`
	Description  string   `json:"description,omitempty" yaml:"description,omitempty"`
	Url          string   `json:"url" yaml:"url"`
	Dependencies []string `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	Commands     string   `json:"cmd,omitempty" yaml:"cmd,omitempty"`
	Destination  string   `json:"destination,omitempty" yaml:"destination,omitempty"`
	Files        []string `json:"files" yaml:"files"`
}

func (r *RegisteredTemplate) MarshalIdent(isJson bool) ([]byte, error) {
	if isJson {
		return json.MarshalIndent(r, "", "  ")
	}
	return yaml.Marshal(r)
}
