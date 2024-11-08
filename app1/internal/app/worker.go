package app

import (
	"github.com/rebirthmonkey/ops/pkg/log"

	userControllerInterface "github.com/rebirthmonkey/ops/app1/internal/user/controller/worker"
	userController "github.com/rebirthmonkey/ops/app1/internal/user/controller/worker/kafka"
	usedUserRepo "github.com/rebirthmonkey/ops/app1/internal/user/repo/mysql"
)

type Worker struct {
	controller userControllerInterface.UserController
}

func NewWorker() (*Worker, error) {
	log.Infoln("[Worker] New")

	userRepo := usedUserRepo.New()
	ctl := userController.New(userRepo)

	w := &Worker{
		controller: ctl,
	}

	return w, nil
}

func (w *Worker) Run() error {
	go func() {
		w.controller.Run()
	}()

	return nil
}
