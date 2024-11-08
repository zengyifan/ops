package kafka

import (
	"github.com/spf13/viper"
)

type Options struct {
	Host          string `json:"host" mapstructure:"host"`
	Port          int    `json:"port" mapstructure:"port"`
	Username      string `json:"user" mapstructure:"user"`
	Password      string `json:"password" mapstructure:"password"`
	Topic         string `json:"channel" mapstructure:"topic"`
	ConsumerGroup string `json:"consumer_group" mapstructure:"consumer_group"`
}

func NewOptions() *Options {
	return &Options{
		Host:          viper.GetString("kafka.host"),
		Port:          viper.GetInt("kafka.port"),
		Username:      viper.GetString("kafka.user"),
		Password:      viper.GetString("kafka.password"),
		Topic:         viper.GetString("kafka.topic"),
		ConsumerGroup: viper.GetString("kafka.consumer_group"),
	}
}

func (o *Options) ApplyTo(c *Config) error {
	c.Host = o.Host
	c.Port = o.Port
	c.Username = o.Username
	c.Password = o.Password
	c.Topic = o.Topic
	c.ConsumerGroup = o.ConsumerGroup
	return nil
}
