package worker

import "github.com/rebirthmonkey/ops/pkg/log"

type Worker struct {
}

func New() (*Worker, error) {
	instance := &Worker{}

	return instance, nil
}

func (w *Worker) Run() error {
	log.Infoln("[Worker] Run")

	return nil
}
