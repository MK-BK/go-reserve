package server

import (
	"go-reserve/models"
)

type JobManager struct{}

func NewJobManager() *JobManager { return &JobManager{} }

var _ models.JobManage = (*JobManager)(nil)

func (m *JobManager) List() ([]*models.Job, error) {
	var jobs []*models.Job
	if err := _db.Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (m *JobManager) Create(job *models.Job) error {
	return _db.Create(job).Error
}

func (m *JobManager) Get(id string) (*models.Job, error) {
	var job *models.Job
	if err := _db.Find(&job, id).Error; err != nil {
		return nil, err
	}

	return job, nil
}

func (m *JobManager) Update(id string, job *models.Job) error {
	return nil
}

func (m *JobManager) Delete(id string) error {
	return _db.Delete(&models.Job{}, id).Error
}

func (m *JobManager) Cancel(id string) error {
	return nil
}
