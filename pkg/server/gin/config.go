package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/ops/pkg/log"
)

type Config struct {
	Address     string
	Mode        string
	Middlewares []string
	Healthz     bool
}

func NewConfig() *Config {
	return &Config{
		Address:     ":8080",
		Healthz:     true,
		Mode:        gin.ReleaseMode,
		Middlewares: []string{},
	}
}

func (c Config) New() (*Server, error) {
	gin.SetMode(c.Mode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Infof("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	s := &Server{
		Config: &c,
		Engine: gin.New(),
	}

	return s, nil
}
