package server

import "go-reserve/models"

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
	return nil
}

func (m *RequestManager) Delete(id string) error {
	return db.Delete(&models.Request{}, id).Error
}
