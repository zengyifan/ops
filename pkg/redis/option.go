package redis

import (
	"github.com/spf13/viper"
)

type Options struct {
	Addr     string `json:"addr" mapstructure:"addr"`
	Password string `json:"password" mapstructure:"password"`
	Database int    `json:"db" mapstructure:"db"`
}

func NewOptions() *Options {
	return &Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		Database: viper.GetInt("redis.db"),
	}
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *Config) error {
	c.Addr = o.Addr
	c.Password = o.Password
	c.Database = o.Database
	return nil
}
