package mq

import (
	"fmt"
	"github.com/rebirthmonkey/ops/pkg/log"
	"github.com/streadway/amqp"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Channel  string
}

func NewConfig() *Config {
	return &Config{
		Host:     "",
		Port:     0,
		Username: "",
		Password: "",
		Channel:  "",
	}
}

func (c *Config) New() (*DB, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", c.Username, c.Password, c.Host, c.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Errorln("MQ.New.Conn error: ", err.Error())
		panic(err)
	}
	//defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Errorln("MQ.New.Channel error: ", err.Error())
		panic(err)
	}
	//defer ch.Close()

	q, err := ch.QueueDeclare(
		c.Channel, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Errorln("MQ.New.QueueDeclare error: ", err.Error())
		panic(err)
	}

	db := &DB{
		Config:  c,
		Channel: ch,
		Queue:   &q,
	}

	return db, nil
}
