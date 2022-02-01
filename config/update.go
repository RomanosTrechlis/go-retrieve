package config

import "github.com/RomanosTrechlis/retemp/util"

func UpdateConfig(c *Configuration) error {
	return util.WriteFile(util.ConfigPath(), "config.json", c)
}
