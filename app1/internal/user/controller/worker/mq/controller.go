package mq

import (
	"encoding/json"
	controllerInterface "github.com/rebirthmonkey/ops/app1/internal/user/controller/worker"
	model "github.com/rebirthmonkey/ops/app1/internal/user/model/v1"
	repoInterface "github.com/rebirthmonkey/ops/app1/internal/user/repo"
	serviceInterface "github.com/rebirthmonkey/ops/app1/internal/user/service"
	serviceImpl "github.com/rebirthmonkey/ops/app1/internal/user/service/v1"
	"github.com/rebirthmonkey/ops/pkg/log"
	"github.com/rebirthmonkey/ops/pkg/mq"
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

func (u *controller) Run() {
	mqInstnance := mq.GetUniqueDBInstance()

	msgs, err := mqInstnance.Channel.Consume(
		mqInstnance.Queue.Name, // queue
		"",                     // consumer
		true,                   // auto-ack
		false,                  // exclusive
		false,                  // no-local
		false,                  // no-wait
		nil,                    // args
	)
	if err != nil {
		log.Errorln("UserController.MQ.Run error: ", err.Error())
		return
	}

	go func() {
		for d := range msgs {
			log.Infoln("[UserController.MQ] Run: Received message: ", string(d.Body))
			mb := model.MQBody{}
			if err := json.Unmarshal(d.Body, &mb); err != nil {
				log.Errorln("[UserController.MQ] Run.json.Unmarshal error: ", err.Error())
				continue
			}
			switch mb.Category {
			case "create":
				u.Create(&mb)
				continue
			case "delete":
				u.Delete(&mb)
				continue
			case "update":
				u.Update(&mb)
				continue
			case "get":
				u.Get(&mb)
				continue
			case "list":
				u.List(&mb)
				continue
			default:
				log.Errorln("[UserController.MQ] Run error: unknown category")
			}
		}
	}()

	forever := make(chan bool)
	<-forever
}

func (u *controller) Create(body *model.MQBody) {
	log.Infoln("[UserController.MQ] create: start")

	if err := u.svc.Create(body.User); err != nil {
		log.Errorln("[UserController.MQ] error: ", err.Error())
		return
	}

	log.Infoln("[UserController.MQ] Create: success")
	return
}

func (u *controller) Delete(body *model.MQBody) {
	log.Infoln("[UserController.MQ] Delete: start")

	if err := u.svc.Delete(body.User.Name); err != nil {
		log.Errorln("[UserController.MQ] Delete: error: ", err.Error())
		return
	}

	log.Infoln("[UserController.MQ] Delete: success")
	return
}

func (u *controller) Update(body *model.MQBody) {
	log.Infoln("[UserController.MQ] Update: start")

	if err := u.svc.Update(body.User); err != nil {
		log.Errorln("[UserController.MQ] Update: error: ", err.Error())
		return
	}

	log.Infoln("[UserController.MQ] Update: success")
	return
}

func (u *controller) Get(body *model.MQBody) {
	log.Infoln("[UserController.MQ] Get: start")

	user, err := u.svc.Get(body.User.Name)
	if err != nil {
		log.Errorln("[UserController.MQ] Get: error: ", err.Error())
		return
	}

	log.Infoln("[UserController.MQ] Get: success with username: ", user.Name)
	return
}

func (u *controller) List(body *model.MQBody) {
	log.Infoln("[UserController.MQ] List: start")

	userList, err := u.svc.List()
	if err != nil {
		log.Errorln("[UserController.MQ] List: error: ", err.Error())
		return
	}

	jsonData, _ := json.Marshal(userList)

	log.Infoln("[UserController.MQ] List: success with ", string(jsonData))
	return
}
