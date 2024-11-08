package redis

import (
	"github.com/go-redis/redis/v8"

	"github.com/rebirthmonkey/ops/pkg/log"
)

type DB struct {
	*Config
	DB *redis.Client
}

var dbInstance *DB

func Init() error {
	opts := NewOptions()
	config := NewConfig()

	if err := opts.ApplyTo(config); err != nil {
		log.Errorln("Redis New ApplyTo Config Error: ", err)
		return err
	}

	db, err := config.New()
	if err != nil {
		log.Errorln("Redis New Config Error: ", err)
		return err
	}

	dbInstance = db

	return nil
}

func GetUniqueDBInstance() *DB {
	if dbInstance == nil {
		log.Errorln("Redis GetUniqueDBInstance Error: dbInstance is nil")
		panic("Redis GetUniqueDBInstance Error: dbInstance is nil")
	}
	return dbInstance
}
