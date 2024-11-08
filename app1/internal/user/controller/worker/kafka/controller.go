package kafka

import (
	"encoding/json"

	controllerInterface "github.com/rebirthmonkey/ops/app1/internal/user/controller/worker"
	model "github.com/rebirthmonkey/ops/app1/internal/user/model/v1"
	repoInterface "github.com/rebirthmonkey/ops/app1/internal/user/repo"
	serviceInterface "github.com/rebirthmonkey/ops/app1/internal/user/service"
	serviceImpl "github.com/rebirthmonkey/ops/app1/internal/user/service/v1"
	kafkaDriver "github.com/zengyifan/ops/pkg/kafka"
	"github.com/rebirthmonkey/ops/pkg/log"
)

var _ controllerInterface.UserController = (*controller)(nil)

type controller struct {
	svc serviceInterface.UserService
}

func New(repo repoInterface.UserRepo) controllerInterface.UserController {
	return &controller{
		svc: serviceImpl.New(repo),
	}
}

func (c *controller) Run() {
	kafkaInstance := kafkaDriver.GetUniqueDBInstance()
	consumer := kafkaInstance.Consumer
	consumer.SubscribeTopics([]string{kafkaInstance.Topic}, nil)
	go func() {
		for {
			msg, err := consumer.ReadMessage(-1)
			if err == nil {
				log.Infoln("Message on", msg.TopicPartition, string(msg.Value))
				body := &model.MQBody{}
				err := json.Unmarshal(msg.Value, body)
				if err != nil {
					log.Errorln("UserController.Kafka.Run error: ", err.Error())
					continue
				}
				switch body.Category {
				case "create":
					c.Create(body)
					continue
				case "delete":
					c.Delete(body)
					continue
				case "update":
					c.Update(body)
					continue
				case "get":
					c.Get(body)
					continue
				case "list":
					c.List(body)
					continue
				default:
					log.Errorln("UserController.Kafka.Run error: ", "action not found")
				}
			}
		}
	}()

	forever := make(chan bool)
	<-forever
}

func (c *controller) Create(body *model.MQBody) {
	log.Infoln("UserController.Kafka.Create")

	if error := c.svc.Create(body.User); error != nil {
		log.Errorln("UserController.Kafka.Create error: ", error.Error())
		return
	}

	log.Infoln("UserController.Kafka.Create success")
}

func (c *controller) Delete(body *model.MQBody) {
	log.Infoln("UserController.Kafka.Delete")

	if error := c.svc.Delete(body.User.Name); error != nil {
		log.Errorln("UserController.Kafka.Delete error: ", error.Error())
		return
	}

	log.Infoln("UserController.Kafka.Delete success")
}

func (c *controller) Update(body *model.MQBody) {
	log.Infoln("UserController.Kafka.Update")

	if error := c.svc.Update(body.User); error != nil {
		log.Errorln("UserController.Kafka.Update error: ", error.Error())
		return
	}

	log.Infoln("UserController.Kafka.Update success")
}

func (c *controller) Get(body *model.MQBody) {
	log.Infoln("UserController.Kafka.Get")
	user, error := c.svc.Get(body.User.Name)
	if error != nil {
		log.Errorln("UserController.Kafka.Get error: ", error.Error())
		return
	}
	log.Infoln("UserController.Kafka.Get success", user.Name)
}

func (c *controller) List(body *model.MQBody) {
	log.Infoln("UserController.Kafka.List")
	usersList, error := c.svc.List()
	if error != nil {
		log.Errorln("UserController.Kafka.List error: ", error.Error())
		return
	}
	jsonData, _ := json.Marshal(usersList)
	log.Infoln("UserController.Kafka.List success", string(jsonData))
}
