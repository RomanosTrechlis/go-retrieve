package registry

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

func Unmarshal(b []byte, isJson bool) (*Registry, error) {
	var c *Registry
	var err error
	if isJson {
		err = json.Unmarshal(b, &c)
	} else {
		err = yaml.Unmarshal(b, &c)
	}

	return c, err
}
