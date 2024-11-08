package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/ops/pkg/log"
	"github.com/rebirthmonkey/ops/pkg/server/gin/util"

	controllerInterface "github.com/rebirthmonkey/ops/app1/internal/user/controller/gin"
	model "github.com/rebirthmonkey/ops/app1/internal/user/model/v1"
	repoInterface "github.com/rebirthmonkey/ops/app1/internal/user/repo"
	serviceInterface "github.com/rebirthmonkey/ops/app1/internal/user/service"
	serviceImpl "github.com/rebirthmonkey/ops/app1/internal/user/service/v1"
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

// @Summary      Create user
// @Description  Create a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body model.User  true  "User to create"
// @Success      200   {object}  model.User
// @Failure      400   {object}  gin.H
// @Router       /users [post]
func (u *controller) Create(c *gin.Context) {
	log.Infoln("[UserController.Gin] Create: start")

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Errorln("User.Controller.Create ErrBind")
		util.WriteResponse(c, errors.New("User.Controller.Create ErrBind"), nil)

		return
	}

	if err := u.svc.Create(&user); err != nil {
		util.WriteResponse(c, err, nil)
		return
	}

	util.WriteResponse(c, nil, user)
}

// @Summary      Delete user
// @Description  Delete an existing user
// @Tags         users
// @Param        name  path      string  true  "User name"
// @Success      200   {object}  gin.H
// @Failure      400   {object}  gin.H
// @Router       /users/{name} [delete]
func (u *controller) Delete(c *gin.Context) {
	log.Infoln("[UserController.Gin] Delete: start")

	if err := u.svc.Delete(c.Param("name")); err != nil {
		util.WriteResponse(c, err, nil)
		return
	}

	var msg string = "deleted user " + c.Param("name")
	util.WriteResponse(c, nil, msg)
}

// @Summary      Update user
// @Description  Update an existing user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        name  path      string  true  "User name"
// @Param        user  body      model.User  true  "User to update"
// @Success      200   {object}  model.User
// @Failure      400   {object}  gin.H
// @Router       /users/{name} [put]
func (u *controller) Update(c *gin.Context) {
	log.Infoln("[UserController.Gin] Update: start")

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Errorln("User.Controller.Update ErrBind", err)
		util.WriteResponse(c, errors.New("User.Controller.Update ErrBind"), nil)
		return
	}

	user.Name = c.Param("name")

	if err := u.svc.Update(&user); err != nil {
		util.WriteResponse(c, err, nil)
		return
	}

	util.WriteResponse(c, nil, user)
}

// @Summary      Get user
// @Description  Get an existing user
// @Tags         users
// @Param        name  path      string  true  "User name"
// @Success      200   {object}  model.User
// @Failure      400   {object}  gin.H
// @Router       /users/{name} [get]
func (u *controller) Get(c *gin.Context) {
	log.Infoln("[UserController.Gin] Get: start")

	user, err := u.svc.Get(c.Param("name"))
	if err != nil {
		util.WriteResponse(c, err, nil)
		return
	}

	util.WriteResponse(c, nil, user)
}

// @Summary      List users
// @Description  List all users
// @Tags         users
// @Success      200   {array}   model.User
// @Failure      400   {object}  gin.H
// @Router       /users [get]
func (u *controller) List(c *gin.Context) {
	log.Infoln("[UserController.Gin] List: start")

	users, err := u.svc.List()
	if err != nil {
		util.WriteResponse(c, err, nil)
		return
	}

	util.WriteResponse(c, nil, users)
}
