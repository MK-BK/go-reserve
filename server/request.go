package server

import "go-reserve/models"

type RequestManager struct{}

var _ models.RequestManager = (*RequestManager)(nil)

func NewRequestManager() *RequestManager {
	return &RequestManager{}
}

func (m *RequestManager) List() ([]*models.Request, error) {
	return nil, nil
}

func (m *RequestManager) Create(r *models.Request) error {
	return _db.Create(r).Error
}

func (m *RequestManager) Get(id string) (*models.Request, error) {
	return nil, nil
}

func (m *RequestManager) Update(id string, r *models.Request) error {
	return nil
}
