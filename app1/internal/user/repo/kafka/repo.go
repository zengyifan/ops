package kafka

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	kafkaDriver "github.com/zengyifan/ops/pkg/kafka"
	"github.com/rebirthmonkey/ops/pkg/log"

	model "github.com/rebirthmonkey/ops/app1/internal/user/model/v1"
	userRepoInterface "github.com/rebirthmonkey/ops/app1/internal/user/repo"
)

var _ userRepoInterface.UserRepo = (*repo)(nil)

type repo struct {
	DB *kafkaDriver.DB
}

func New() userRepoInterface.UserRepo {
	db := kafkaDriver.GetUniqueDBInstance()

	return &repo{
		DB: db,
	}
}

func (u *repo) Create(user *model.User) error {
	log.Infoln("[UserRepo.Kafka] Create: start")

	mb := model.MQBody{
		Category: "create",
		User:     user,
	}

	jsonMQBody, err := json.Marshal(mb)
	if err != nil {
		log.Errorln("[UserRepo.Kafka] Create: error ", err.Error())
		return err
	}

	if err := u.DB.Producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &u.DB.Config.Topic, Partition: kafka.PartitionAny},
			Value:          jsonMQBody,
		},
		nil,
	); err != nil {
		log.Errorln("[UserRepo.Kafka] Create.Publish error: ", err.Error())
		return err
	}

	log.Infoln("[UserRepo.Kafka] Create: success with ", mb)
	return nil
}

func (u *repo) Delete(username string) error {
	log.Infoln("[UserRepo.Kafka] Delete: start")

	user := &model.User{
		Name: username,
	}

	mb := model.MQBody{
		Category: "delete",
		User:     user,
	}

	jsonMQBody, err := json.Marshal(mb)
	if err != nil {
		log.Errorln("[UserRepo.Kafka] Delete: error ", err.Error())
		return err
	}

	if err := u.DB.Producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &u.DB.Config.Topic, Partition: kafka.PartitionAny},
			Value:          jsonMQBody,
		},
		nil,
	); err != nil {
		log.Errorln("[UserRepo.Kafka] Delete.Publish error: ", err.Error())
		return err
	}

	log.Infoln("[UserRepo.Kafka] Delete: success")
	return nil
}

func (u *repo) Update(user *model.User) error {
	log.Infoln("[UserRepo.Kafka] Update: start")

	mb := model.MQBody{
		Category: "update",
		User:     user,
	}

	jsonMQBody, err := json.Marshal(mb)
	if err != nil {
		log.Errorln("[UserRepo.Kafka] Update: error ", err.Error())
		return err
	}

	if err := u.DB.Producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &u.DB.Config.Topic, Partition: kafka.PartitionAny},
			Value:          jsonMQBody,
		},
		nil,
	); err != nil {
		log.Errorln("[UserRepo.Kafka] Update.Publish error: ", err.Error())
		return err
	}
	return nil
}

func (u *repo) Get(username string) (*model.User, error) {
	log.Infoln("[UserRepo.Kafka] Get: start")

	user := &model.User{
		Name: username,
	}

	mb := model.MQBody{
		Category: "get",
		User:     user,
	}

	jsonMQBody, err := json.Marshal(mb)
	if err != nil {
		log.Errorln("[UserRepo.Kafka] Get: error ", err.Error())
		return nil, err
	}

	if err := u.DB.Producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &u.DB.Config.Topic, Partition: kafka.PartitionAny},
			Value:          jsonMQBody,
		},
		nil,
	); err != nil {
		log.Errorln("[UserRepo.Kafka] Get.Publish error: ", err.Error())
		return nil, err
	}

	log.Infoln("[UserRepo.Kafka] Get: success")
	return user, nil
}

func (u *repo) List() (*model.UserList, error) {
	log.Infoln("[UserRepo.Kafka] List: start")

	users := &model.UserList{}

	mb := model.MQBody{
		Category: "list",
		UserList: users,
	}

	jsonMQBody, err := json.Marshal(mb)
	if err != nil {
		log.Errorln("[UserRepo.Kafka] List: error ", err.Error())
		return nil, err
	}

	if err := u.DB.Producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &u.DB.Config.Topic, Partition: kafka.PartitionAny},
			Value:          jsonMQBody,
		},
		nil,
	); err != nil {
		log.Errorln("[UserRepo.Kafka] List.Publish error: ", err.Error())
		return nil, err
	}

	log.Infoln("[UserRepo.Kafka] List: success")
	return users, nil
}
