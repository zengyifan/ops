package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	model "github.com/rebirthmonkey/ops/app1/internal/user/model/v1"
	userRepoInterface "github.com/rebirthmonkey/ops/app1/internal/user/repo"
	"github.com/rebirthmonkey/ops/pkg/log"
	restDriver "github.com/rebirthmonkey/ops/pkg/rest"
	"io/ioutil"
	"net/http"
)

var _ userRepoInterface.UserRepo = (*repo)(nil)

type repo struct {
	DB *restDriver.DB
}

func New() userRepoInterface.UserRepo {
	db := restDriver.GetUniqueDBInstance()
	return &repo{
		DB: db,
	}
}

func (u *repo) Create(user *model.User) error {
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Errorln("UserRepo.Create.json.Marshal error: ", err.Error())
		return err
	}

	resp, err := http.Post(u.DB.URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Errorln("UserRepo.Create.Post error: ", err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("UserRepo.Create error: failed to create user")
	}

	return nil
}

func (u *repo) Delete(username string) error {
	url := fmt.Sprintf("%s%s", u.DB.URL, username)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Errorln("UserRepo.Delete error: ", err.Error())
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorln("UserRepo.Delete error: ", err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Errorln("UserRepo.Delete error: ", string(body))
		return errors.New("UserRepo.Delete error: failed to delete user")
	}

	return nil
}

func (u *repo) Update(user *model.User) error {
	url := fmt.Sprintf("%s%s", u.DB.URL, user.Name)
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Errorln("UserRepo.Update.json.Marshal error: ", err.Error())
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Errorln("UserRepo.Update error: ", err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorln("UserRepo.Update error: ", err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Errorln("UserRepo.Update error: ", string(body))
		return errors.New("UserRepo.Update error: failed to update user")
	}

	return nil
}

func (u *repo) Get(username string) (*model.User, error) {
	url := fmt.Sprintf("%s%s", u.DB.URL, username)
	resp, err := http.Get(url)
	if err != nil {
		log.Errorln("UserRepo.Get error: ", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Errorln("UserRepo.Get error: ", string(body))
		return nil, errors.New("UserRepo.Get error: user not found")
	}

	var user model.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		log.Errorln("UserRepo.Get.json.Unmarshal error: ", err.Error())
		return nil, err
	}

	return &user, nil
}

func (u *repo) List() (*model.UserList, error) {
	url := fmt.Sprintf("%s", u.DB.URL)
	resp, err := http.Get(url)
	if err != nil {
		log.Errorln("UserRepo.List error: ", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	//if resp.StatusCode != http.StatusOK {
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Errorln("UserRepo.List error: ", string(body))
		return nil, errors.New("UserRepo.List error: failed to list users")
	}
	//}

	//var users []*model.User
	var users model.UserList
	//err = json.NewDecoder(resp.Body).Decode(&users)
	err = json.Unmarshal(body, &users)
	if err != nil {
		log.Errorln("UserRepo.List.json.Unmarshal error: ", err.Error())
		return nil, err
	}

	//ret := &model.UserList{
	//	Items: users,
	//}

	return &users, nil
}
