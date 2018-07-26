package env

import (
	"github.com/spiral/roadrunner/service"
)

// Config defines set of env values for RR workers.
type Config struct {
	// Values to set as worker _ENV.
	Values map[string]string
}

// Hydrate must populate Config values using given Config source. Must return error if Config is not valid.
func (c *Config) Hydrate(cfg service.Config) error {
	c.Values = map[string]string{"RR": "YES"}
	return cfg.Unmarshal(&c.Values)
}
