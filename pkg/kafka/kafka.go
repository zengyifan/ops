package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rebirthmonkey/ops/pkg/log"
)

type DB struct {
	*Config
	Producer *kafka.Producer
	Consumer *kafka.Consumer
}

var dbInstance *DB

func Init() error {
	opts := NewOptions()
	config := NewConfig()

	if err := opts.ApplyTo(config); err != nil {
		log.Errorln("Kafka Init ApplyTo Config Error: ", err)
		return err
	}

	db, err := config.New()
	if err != nil {
		log.Errorln("Kafka Init Config Error: ", err)
		return err
	}

	dbInstance = db
	return nil
}

func GetUniqueDBInstance() *DB {
	if dbInstance == nil {
		log.Errorln("Kafka GetUniqueDBInstance Error: ", "dbInstance is nil")
		panic("dbInstance is nil")
	}
	return dbInstance
}
