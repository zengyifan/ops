package rest

import (
	"fmt"
	"github.com/rebirthmonkey/ops/pkg/log"
	"net/http"
)

type Config struct {
	Host string
	Port int
}

func NewConfig() *Config {
	return &Config{
		Host: "",
		Port: 0,
	}
}

func (c *Config) New() (*DB, error) {
	url := fmt.Sprintf("http://%s:%d/v1/users/", c.Host, c.Port)

	resp, err := http.Get(url)
	if err != nil {
		log.Errorln("Rest.New error: ", err.Error())
		panic(err)
	}
	defer resp.Body.Close()

	db := &DB{
		Config: c,
		URL:    url,
	}

	return db, nil
}
