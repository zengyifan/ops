package rest

import (
	"github.com/spf13/viper"
)

type Options struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
}

func NewOptions() *Options {
	return &Options{
		Host: viper.GetString("rest.host"),
		Port: viper.GetInt("rest.port"),
	}
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *Config) error {
	c.Host = o.Host
	c.Port = o.Port
	return nil
}
