package registry

import (
	"encoding/json"

	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/dl"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/util"
)

func Upgrade(e *env.ConfigEnv) error {
	c, err := config.LoadConfig(e)
	if err != nil {
		return err
	}

	sources := c.Active.Sources
	for _, s := range sources {
		err := upgradeSource(e, s)
		if err != nil {
			return err
		}
	}

	return nil
}

func upgradeSource(e *env.ConfigEnv, s *config.ConfigurationSource) error {
	b, err := dl.Download(s.Url, s.Token)
	if err != nil {
		return err
	}

	var r *Registry
	err = json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	return util.WriteFile(e.ConfigPath(), s.Name+".json", r)
}
