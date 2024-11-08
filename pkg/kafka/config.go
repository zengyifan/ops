package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rebirthmonkey/ops/pkg/log"
)

type Config struct {
	Host          string
	Port          int
	Username      string
	Password      string
	Topic         string
	ConsumerGroup string
}

func NewConfig() *Config {
	return &Config{
		Host:          "",
		Port:          0,
		Username:      "",
		Password:      "",
		Topic:         "",
		ConsumerGroup: "",
	}
}

func (c *Config) New() (*DB, error) {
	url := fmt.Sprintf("%s:%d", c.Host, c.Port)
	log.Infoln("Kafka.New.URL: ", url)
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": url,
	})
	if err != nil {
		log.Errorln("Kafka.New.Producer error: ", err.Error())
		panic(err)
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": url,
		"group.id":          c.ConsumerGroup,
	})
	if err != nil {
		log.Errorln("Kafka.New.Consumer error: ", err.Error())
		panic(err)
	}

	db := &DB{
		Config:   c,
		Producer: producer,
		Consumer: consumer,
	}

	return db, nil
}
