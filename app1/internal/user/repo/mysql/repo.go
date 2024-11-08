// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	"errors"
	"regexp"

	"gorm.io/gorm"

	model "github.com/rebirthmonkey/ops/app1/internal/user/model/v1"
	userRepoInterface "github.com/rebirthmonkey/ops/app1/internal/user/repo"
	"github.com/rebirthmonkey/ops/pkg/log"
	mysqlDriver "github.com/rebirthmonkey/ops/pkg/mysql"
)

var _ userRepoInterface.UserRepo = (*repo)(nil)

type repo struct {
	DB *mysqlDriver.DB
}

func New() userRepoInterface.UserRepo {
	db := mysqlDriver.GetUniqueDBInstance()

	return &repo{
		DB: db,
	}
}

func (u *repo) Create(user *model.User) error {
	tmpUser := model.User{}
	u.DB.DBEngine.Where("name = ?", user.Name).Find(&tmpUser)
	if tmpUser.Name != "" {
		log.Errorln("UserRepo.Create error: RecordAlreadyExist")
		return errors.New("UserRepo.Create error: ErrRecordAlreadyExist")
	}

	if err := u.DB.DBEngine.Create(&user).Error; err != nil {
		if match, _ := regexp.MatchString("Duplicate entry", err.Error()); match {
			return errors.New("UserRepo.Create error: ErrRecordAlreadyExist")
		}

		return err
	}

	return nil
}

func (u *repo) Delete(username string) error {
	if err := u.DB.DBEngine.Where("name = ?", username).Delete(&model.User{}).Error; err != nil {
		return err
	}

	return nil
}

func (u *repo) Update(user *model.User) error {
	tmpUser := model.User{}
	u.DB.DBEngine.Where("name = ?", user.Name).Find(&tmpUser)
	if tmpUser.Name == "" {
		err := errors.New("UserRepo.Update error: ErrRecordNotFound")
		log.Errorln("UserRepo.Update error: ErrRecordNotFound")
		return err
	}

	//if err := u.DB.DBEngine.Save(user).Error; err != nil {
	//	log.Errorln("[UserRepo.MySQL] Update: error ", err.Error())
	//	return err
	//}

	if err := u.DB.DBEngine.Model(&user).Where("name = ?", user.Name).Updates(user).Error; err != nil {
		log.Errorln("[UserRepo.MySQL] Update: error ", err.Error())
		return err
	}

	return nil
}

func (u *repo) Get(username string) (*model.User, error) {
	user := &model.User{}
	err := u.DB.DBEngine.Where("name = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("UserRepo.Get error: ErrRecordNotFound")
		}

		return nil, errors.New("UserRepo.Create error: ErrRecordNotFound")
	}

	return user, nil
}

func (u *repo) List() (*model.UserList, error) {
	ret := &model.UserList{}

	d := u.DB.DBEngine.
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}
