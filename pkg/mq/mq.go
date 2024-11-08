package mq

import (
	"github.com/rebirthmonkey/ops/pkg/log"
	"github.com/streadway/amqp"
)

type DB struct {
	*Config
	Channel *amqp.Channel
	Queue   *amqp.Queue
}

var dbInstance *DB

func Init() error {
	opts := NewOptions()
	config := NewConfig()

	if err := opts.ApplyTo(config); err != nil {
		log.Errorln("MQ Init ApplyTo Config Error: ", err)
		return err
	}

	db, err := config.New()
	if err != nil {
		log.Errorln("MQ Init Config Error: ", err)
		return err
	}

	dbInstance = db
	return nil
}

//func New() *DB {
//	opts := NewOptions()
//	config := NewConfig()
//
//	if err := opts.ApplyTo(config); err != nil {
//		log.Errorln("MQ Init ApplyTo Config Error: ", err)
//		return nil
//	}
//
//	db, err := config.New()
//	if err != nil {
//		log.Errorln("MQ Init Config Error: ", err)
//		return nil
//	}
//
//	return db
//}

func GetUniqueDBInstance() *DB {
	if dbInstance == nil {
		log.Errorln("MQ GetUniqueDBInstance Error: dbInstance is nil")
		panic("MQ GetUniqueDBInstance Error: dbInstance is nil")
	}

	return dbInstance
}
