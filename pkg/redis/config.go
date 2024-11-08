package redis

import (
	"context"

	"github.com/go-redis/redis/v8"

	"github.com/rebirthmonkey/ops/pkg/log"
)

type Config struct {
	Addr     string
	Password string
	Database int
}

func NewConfig() *Config {
	return &Config{
		Addr:     "",
		Password: "",
		Database: 0,
	}
}

func (c *Config) New() (*DB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.Database,
	})

	_, err := client.SMembers(context.Background(), "groupset").Result()
	if err != nil {
		log.Errorln("ConnectRedis executing Redis query Error: ", err)
		panic(err)
	}

	db := &DB{
		Config: c,
		DB:     client,
	}

	return db, nil
}
