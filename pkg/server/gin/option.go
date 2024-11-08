package gin

import (
	"github.com/spf13/viper"
	"net"
	"strconv"
)

type Options struct {
	BindAddress string   `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int      `json:"bind-port"    mapstructure:"bind-port"`
	Mode        string   `json:"mode"        mapstructure:"mode"`
	Healthz     bool     `json:"healthz"     mapstructure:"healthz"`
	Middlewares []string `json:"middlewares" mapstructure:"middlewares"`
}

func NewOptions() *Options {
	return &Options{
		BindAddress: viper.GetString("gin.bind-address"),
		BindPort:    viper.GetInt("gin.bind-port"),
		Mode:        viper.GetString("gin.mode"),
		Healthz:     viper.GetBool("gin.healthz"),
		Middlewares: viper.GetStringSlice("gin.middlewares"),
	}
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *Config) error {
	c.Address = net.JoinHostPort(o.BindAddress, strconv.Itoa(o.BindPort))
	c.Mode = o.Mode
	c.Middlewares = o.Middlewares
	c.Healthz = o.Healthz
	return nil
}
