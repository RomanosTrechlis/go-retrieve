package config

type Configuration struct {
	Active   *ConfigurationProfile   `json:"active"`
	Profiles []*ConfigurationProfile `json:"profiles"`
}

type ConfigurationProfile struct {
	Name    string                 `json:"name"`
	Sources []*ConfigurationSource `json:"sources"`
}

type ConfigurationSource struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Token string `json:"token"`
}
