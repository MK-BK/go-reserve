package server

import (
	"go-reserve/models"
	"go-reserve/server/audit"
)

type RequestManager struct{}

var _ models.RequestManager = (*RequestManager)(nil)

func NewRequestManager() *RequestManager {
	return &RequestManager{}
}

func (m *RequestManager) List() ([]*models.Request, error) {
	var rs []*models.Request
	if err := db.Find(&rs).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return rs, nil
}

func (m *RequestManager) Create(r *models.Request) error {
	defer func() {
		audit.NewAuditLog().
			WithAction(audit.ActionRequestCreate).
			Notify()
	}()

	if r.Status == "" {
		r.Status = models.StatusNew
	}
	return db.Create(r).Error
}

func (m *RequestManager) Get(id string) (*models.Request, error) {
	var r *models.Request
	if err := db.Find(&r, id).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return r, nil
}

func (m *RequestManager) Update(id string, r *models.Request) error {
	defer func() {
		audit.NewAuditLog().
			WithAction(audit.ActionRequestUpdate).
			Notify()
	}()

	if r.Status == models.StatusApprove {
		// NewEmail().Send()

		go func() {

		}()
	}

	return nil
}

func (m *RequestManager) Delete(id string) error {
	defer func() {
		audit.NewAuditLog().
			WithAction(audit.ActionRequestDelete).
			Notify()
	}()
	return db.Delete(&models.Request{}, id).Error
}
