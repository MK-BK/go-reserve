package audit

import (
	"fmt"
	"go-reserve/db"
)

var (
	ch  = make(chan *AuditLog, 10)
	_db = db.GetDB()
)

func init() {
	go notify()
}

func notify() {
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
		fmt.Println(err)
	}
}
