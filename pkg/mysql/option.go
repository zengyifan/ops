package mysql

import (
	"github.com/spf13/viper"
)

type Options struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	Username string `json:"user" mapstructure:"user"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"dbname" mapstructure:"dbname"`
}

func NewOptions() *Options {
	return &Options{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetInt("mysql.port"),
		Username: viper.GetString("mysql.user"),
		Password: viper.GetString("mysql.password"),
		Database: viper.GetString("mysql.dbname"),
	}
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *Config) error {
	c.Host = o.Host
	c.Port = o.Port
	c.Username = o.Username
	c.Password = o.Password
	c.Database = o.Database
	return nil
}
