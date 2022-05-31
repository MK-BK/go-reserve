package server

import (
	"go-reserve/db"
	"go-reserve/models"
)

type EnvManage struct {
	UserManger      models.UserManaer
	JobManage       models.JobManage
	ProductManage   models.ShopManager
	SessionManage   models.SessionManage
	CommodityManage models.CommodityManager
	RequestManager  models.RequestManager
}

var (
	Env EnvManage
	_db = db.GetDB()
)

func init() {
	Env = EnvManage{
		UserManger:      NewUserManager(),
		JobManage:       NewJobManager(),
		ProductManage:   NewShopManager(),
		SessionManage:   NewSessionManager(),
		CommodityManage: NewCommodityManager(),
		RequestManager:  NewRequestManager(),
	}
}
