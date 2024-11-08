// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	model "github.com/rebirthmonkey/ops/app1/internal/user/model/v1"
	"github.com/rebirthmonkey/ops/app1/internal/user/repo"
	userServiceInterface "github.com/rebirthmonkey/ops/app1/internal/user/service"
	"golang.org/x/crypto/bcrypt"
)

var _ userServiceInterface.UserService = (*service)(nil)

type service struct {
	repo repo.UserRepo
}

func New(repo repo.UserRepo) userServiceInterface.UserService {
	return &service{
		repo: repo,
	}
}

func (u *service) Create(user *model.User) error {
	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedBytes)
	user.Status = 1

	return u.repo.Create(user)
}

func (u *service) Delete(username string) error {
	return u.repo.Delete(username)
}

func (u *service) Update(user *model.User) error {
	//updateUser, err := u.repo.Get(user.Name)
	//if err != nil {
	//	log.Errorln("[UserService] Update error: ", err)
	//	return err
	//}

	updateUser := &model.User{}
	updateUser.Name = user.Name
	updateUser.Nickname = user.Nickname
	updateUser.Email = user.Email
	updateUser.Phone = user.Phone

	return u.repo.Update(updateUser)
}

func (u *service) Get(username string) (*model.User, error) {
	return u.repo.Get(username)
}

func (u *service) List() (*model.UserList, error) {
	users, err := u.repo.List()
	if err != nil {
		return nil, err
	}

	infos := make([]*model.User, 0)
	for _, user := range users.Items {
		infos = append(infos, &model.User{

			Name:     user.Name,
			Nickname: user.Nickname,
			Email:    user.Email,
			Phone:    user.Phone,
			ID:       user.ID,
		})
	}

	return &model.UserList{ListMeta: users.ListMeta, Items: infos}, nil
}
