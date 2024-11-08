package worker

import model "github.com/rebirthmonkey/ops/app1/internal/user/model/v1"

type UserController interface {
	Run()
	Create(body *model.MQBody)
	Delete(body *model.MQBody)
	Update(body *model.MQBody)
	Get(body *model.MQBody)
	List(body *model.MQBody)
}
