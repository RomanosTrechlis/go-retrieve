package config

import (
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/util"
)

func UpdateConfig(e *env.ConfigEnv, c *Configuration) error {
	return util.WriteFile(e.ConfigPath(), e.ConfigName, c)
}
