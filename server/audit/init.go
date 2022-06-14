package audit

import (
	"go-reserve/util/db"

	"github.com/sirupsen/logrus"
)

var _db = db.GetDB()

var ch chan *AuditLog

func init() {
	ch = make(chan *AuditLog, 10)
	go worker()
}

func worker() {
	for {
		al, ok := <-ch
		if !ok {
			return
		}
		save(al)
	}
}

func save(al *AuditLog) {
	err := _db.Create(al).Error
	if err != nil {
		logrus.Error("error msg")
	}
}
