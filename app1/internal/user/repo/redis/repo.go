// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	"context"
	"encoding/json"
	model "github.com/rebirthmonkey/ops/app1/internal/user/model/v1"
	userRepoInterface "github.com/rebirthmonkey/ops/app1/internal/user/repo"
	"github.com/rebirthmonkey/ops/pkg/log"
	redisDriver "github.com/rebirthmonkey/ops/pkg/redis"
)

var _ userRepoInterface.UserRepo = (*repo)(nil)

type repo struct {
	DB *redisDriver.DB
}

func New() userRepoInterface.UserRepo {
	db := redisDriver.GetUniqueDBInstance()

	return &repo{
		DB: db,
	}
}

func (u *repo) Create(user *model.User) error {
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Errorln("Repo.Redis.Create.json.Marshal error: ", err.Error())
		return err
	}

	err = u.DB.DB.Set(context.Background(), user.Name, jsonData, 0).Err()
	if err != nil {
		log.Errorln("Repo.Redis.Create error: ", err.Error())
		return err
	}

	return nil
}

func (u *repo) Delete(username string) error {
	err := u.DB.DB.Del(context.Background(), username).Err()
	if err != nil {
		log.Errorln("Repo.Redis.Delete error: ", err.Error())
		return err
	}

	return nil
}

func (u *repo) Update(user *model.User) error {
	if err := u.Delete(user.Name); err != nil {
		return err
	}

	if err := u.Create(user); err != nil {
		return err
	}

	return nil
}

func (u *repo) Get(username string) (*model.User, error) {
	user := &model.User{}

	val, err := u.DB.DB.Get(context.Background(), username).Result()
	if err != nil {
		log.Errorln("Repo.Redis.Get error: ", err.Error())
		return nil, err
	}

	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		log.Errorln("Repo.Redis.Create.json.Unmarshal error: ", err.Error())
		return nil, err
	}

	return user, nil
}

func (u *repo) List() (*model.UserList, error) {
	ret := &model.UserList{}

	keys, err := u.DB.DB.Keys(context.Background(), "*").Result()
	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		val, err := u.DB.DB.Get(context.Background(), key).Result()
		if err != nil {
			log.Errorln("Repo.Redis.List error: ", err.Error())
			return nil, err
		}

		var user model.User
		err = json.Unmarshal([]byte(val), &user)
		if err != nil {
			log.Errorln("Repo.Redis.List.json.Unmarshal error: ", err.Error())
			return nil, err
		}

		ret.Items = append(ret.Items, &user)
	}

	ret.TotalCount = int64(len(ret.Items))
	return ret, nil
}
