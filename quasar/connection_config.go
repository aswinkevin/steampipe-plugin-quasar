package quasar

type quasarConfig struct {
	Domains []string `hcl:"regions"`
}

func (c *quasarConfig) GetDomains() []string {
	return c.Domains
}

func ConfigInstance() interface{} {
	return &quasarConfig{}
}
