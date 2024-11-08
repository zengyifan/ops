package app

import (
	kafkaDriver "github.com/zengyifan/ops/pkg/kafka"
	"github.com/rebirthmonkey/ops/pkg/log"
	mysqlDriver "github.com/rebirthmonkey/ops/pkg/mysql"
	redisDriver "github.com/rebirthmonkey/ops/pkg/redis"
	"github.com/rebirthmonkey/ops/pkg/utils"
)

type App struct {
	name        string
	description string

	ginServer *Server
	worker    *Worker
}

func New(name string) *App {
	utils.InitConfig()

	if err := mysqlDriver.Init(); err != nil {
		log.Errorln("Mysql.Init error: ", err)
	}

	if err := redisDriver.Init(); err != nil {
		log.Errorln("Redis.Init error: ", err)
	}

	if err := kafkaDriver.Init(); err != nil {
		log.Errorln("MQ.Init error: ", err)
	}

	//if err := restDriver.Init(); err != nil {
	//	log.Errorln("REST.Init error: ", err)
	//}

	ginServer, err := NewServer()
	if err != nil {
		log.Errorln("[App.GinServer] Init: error ", err)
	}

	worker, err := NewWorker()
	if err != nil {
		log.Errorln("[App.Worker] Init: error ", err)
	}

	app := &App{
		name:      name,
		ginServer: ginServer,
		worker:    worker,
	}

	return app
}

func (app *App) Run() {
	log.Infoln("[App] Run")
	app.worker.Run()
	app.ginServer.Run()
}
