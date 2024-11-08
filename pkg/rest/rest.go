package rest

import (
	"github.com/rebirthmonkey/ops/pkg/log"
)

type DB struct {
	*Config
	URL string
}

var dbInstance *DB

func Init() error {
	opts := NewOptions()
	config := NewConfig()

	if err := opts.ApplyTo(config); err != nil {
		log.Errorln("REST Init ApplyTo Config Error: ", err)
		return err
	}

	db, err2 := config.New()
	if err2 != nil {
		log.Errorln("REST Init Config Error: ", err2)
		return err2
	}

	dbInstance = db
	return nil
}

func GetUniqueDBInstance() *DB {
	if dbInstance == nil {
		log.Errorln("REST GetUniqueDBInstance Error: dbInstance is nil")
		panic("REST GetUniqueDBInstance Error: dbInstance is nil")
	}
	return dbInstance
}
