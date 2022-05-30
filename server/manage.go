package server

import (
	"go-reserve/models"
	"go-reserve/server/job"
	"go-reserve/server/product"
	"go-reserve/server/session"
	"go-reserve/server/user"
)

type EnvManage struct {
	UserManger    models.UserManaer
	JobManage     models.JobManage
	ProductManage models.ProductManage
	SessionManage models.SessionManage
}

var Env EnvManage

func init() {
	Env = EnvManage{
		UserManger:    user.NewManger(),
		JobManage:     job.NewManager(),
		ProductManage: product.NewManager(),
		SessionManage: session.NewManager(),
	}
}
