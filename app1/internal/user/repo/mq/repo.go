package mq

import (
	"encoding/json"

	"github.com/rebirthmonkey/ops/pkg/log"
	mqDriver "github.com/rebirthmonkey/ops/pkg/mq"
	"github.com/streadway/amqp"

	model "github.com/rebirthmonkey/ops/app1/internal/user/model/v1"
	userRepoInterface "github.com/rebirthmonkey/ops/app1/internal/user/repo"
)

var _ userRepoInterface.UserRepo = (*repo)(nil)

type repo struct {
	DB *mqDriver.DB
}

func New() userRepoInterface.UserRepo {
	db := mqDriver.GetUniqueDBInstance()

	return &repo{
		DB: db,
	}
}

func (u *repo) Create(user *model.User) error {
	log.Infoln("[UserRepo.MQ] Create: start")

	mb := model.MQBody{
		Category: "create",
		User:     user,
	}

	jsonMQBody, err := json.Marshal(mb)
	if err != nil {
		log.Errorln("[UserRepo.MQ] Create: error ", err.Error())
		return err
	}

	if err := u.DB.Channel.Publish(
		"",
		u.DB.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        jsonMQBody,
		}); err != nil {
		log.Errorln("UserRepo.Create.Publish error: ", err.Error())
	}

	log.Infoln("[UserRepo.MQ] Create: success with username", user.Name)
	return nil
}

func (u *repo) Delete(username string) error {
	log.Infoln("[UserRepo.MQ] Delete: start")

	user := &model.User{
		Name: username,
	}

	mb := model.MQBody{
		Category: "delete",
		User:     user,
	}

	jsonMQBody, err := json.Marshal(mb)
	if err != nil {
		log.Errorln("UserRepo.Delete.json.Marshal error: ", err.Error())
		return err
	}

	if err := u.DB.Channel.Publish(
		"",
		u.DB.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        jsonMQBody,
		}); err != nil {
		log.Errorln("UserRepo.Delete.Publish error: ", err.Error())
	}

	log.Infoln("[UserRepo.MQ] Delete: success")

	return nil
}

func (u *repo) Update(user *model.User) error {
	log.Infoln("[UserRepo.MQ] Update: start")

	mb := model.MQBody{
		Category: "update",
		User:     user,
	}

	jsonMQBody, err := json.Marshal(mb)
	if err != nil {
		log.Errorln("[UserRepo.MQ] Update: error ", err.Error())
		return err
	}

	if err := u.DB.Channel.Publish(
		"",
		u.DB.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        jsonMQBody,
		}); err != nil {
		log.Errorln("[UserRepo.MQ] Update: error: ", err.Error())
	}

	log.Infoln("[UserRepo.MQ] Update: success with ", string(jsonMQBody))
	return nil
}

func (u *repo) Get(username string) (*model.User, error) {
	log.Infoln("[UserRepo.MQ] Get: start")

	user := &model.User{
		Name: username,
	}

	mb := model.MQBody{
		Category: "get",
		User:     user,
	}

	jsonMQBody, err := json.Marshal(mb)
	if err != nil {
		log.Errorln("[UserRepo.MQ] Get: error ", err.Error())
		return nil, err
	}

	if err := u.DB.Channel.Publish(
		"",
		u.DB.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        jsonMQBody,
		}); err != nil {
		log.Errorln("[UserRepo.MQ] Get: error ", err.Error())
		return nil, err
	}

	log.Infoln("[UserRepo.MQ] Get: success with ", mb)
	return nil, nil
}

func (u *repo) List() (*model.UserList, error) {
	log.Infoln("[UserRepo.MQ] List: start")

	users := &model.UserList{}

	mb := model.MQBody{
		Category: "list",
		UserList: users,
	}

	jsonMQBody, err := json.Marshal(mb)
	if err != nil {
		log.Errorln("[UserRepo.MQ] List: error ", err.Error())
		return nil, err
	}

	if err := u.DB.Channel.Publish(
		"",
		u.DB.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        jsonMQBody,
		}); err != nil {
		log.Errorln("[UserRepo.MQ] List: error ", err.Error())
	}

	log.Infoln("[UserRepo.MQ] List: success")
	return users, nil
}
