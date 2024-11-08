package mq

import (
	"github.com/spf13/viper"
)

type Options struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	Username string `json:"user" mapstructure:"user"`
	Password string `json:"password" mapstructure:"password"`
	Channel  string `json:"channel" mapstructure:"channel"`
}

func NewOptions() *Options {
	return &Options{
		Host:     viper.GetString("mq.host"),
		Port:     viper.GetInt("mq.port"),
		Username: viper.GetString("mq.user"),
		Password: viper.GetString("mq.password"),
		Channel:  viper.GetString("mq.channel"),
	}
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *Config) error {
	c.Host = o.Host
	c.Port = o.Port
	c.Username = o.Username
	c.Password = o.Password
	c.Channel = o.Channel
	return nil
}
