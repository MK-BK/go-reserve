package audit

import (
	"context"

	"gorm.io/gorm"
)

type AuditLog struct {
	gorm.Model
	UserID      string
	Action      string
	ShopID      string
	CommodityID string
}

func NewAuditLog() *AuditLog {
	return &AuditLog{}
}

func (al *AuditLog) InContext(ctx context.Context) *AuditLog {
	al.UserID = "admin"
	al.Action = "user_create"
	return al
}

func (al *AuditLog) Notify() {
	ch <- al
}
