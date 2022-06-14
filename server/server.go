package server

import (
	"go-reserve/models"
	DB "go-reserve/util/db"

	"github.com/sirupsen/logrus"
)

var (
	db  = DB.GetDB()
	log = logrus.New()
)

type EnvManager struct {
	AccountManager   models.AccountManager
	JobManager       models.JobManager
	ShoptManager     models.ShopManager
	SessionManager   models.SessionManager
	CommodityManager models.CommodityManager
	RequestManager   models.RequestManager
}

var Env = EnvManager{
	AccountManager:   NewAccountManager(),
	JobManager:       NewJobManager(),
	ShoptManager:     NewShopManager(),
	SessionManager:   NewSessionManager(),
	CommodityManager: NewCommodityManager(),
	RequestManager:   NewRequestManager(),
}
