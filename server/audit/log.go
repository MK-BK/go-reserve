package audit

import (
	"context"
	"go-reserve/models"

	"gorm.io/gorm"
)

const (
	StatusSuccess = "success"
	StatusFailure = "failed"
	StatusWarn    = "warn"
)

type AuditLog struct {
	gorm.Model
	UserID string
	Action string
	Status string
}

func NewAuditLog() *AuditLog {
	return &AuditLog{
		Status: StatusSuccess,
	}
}

func (al *AuditLog) WithAction(action string) *AuditLog {
	al.Action = action
	return al
}

func (al *AuditLog) InContext(ctx context.Context) *AuditLog {
	value := ctx.Value(models.CtxHeaderUser)

	if v, ok := value.(string); ok {
		al.UserID = v
	}

	return al
}

func (al *AuditLog) DetectError(err error) *AuditLog {
	al.Status = StatusSuccess
	if err != nil {
		al.Status = StatusFailure
	}
	return al
}

func (al *AuditLog) Notify() {
	ch <- al
}
