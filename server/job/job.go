package job

import (
	"go-reserve/common/db"
	"go-reserve/models"
)

var _db = db.GetDB()

type manager struct{}

func NewManager() *manager { return &manager{} }

var _ models.JobManage = (*manager)(nil)

func (m *manager) Create(job *models.Job) error {
	return _db.Create(job).Error
}

func (m *manager) Get(id string) (*models.Job, error) {
	var job *models.Job
	if err := _db.Find(&job, id).Error; err != nil {
		return nil, err
	}

	return job, nil
}

func (m *manager) List() ([]*models.Job, error) {
	var jobs []*models.Job
	if err := _db.Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (m *manager) Update(id string, job *models.Job) error { return nil }

func (m *manager) Delete(id string) error {
	return _db.Delete(&models.Job{}, id).Error
}
