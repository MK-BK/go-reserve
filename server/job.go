package server

import (
	"go-reserve/models"
)

type JobManager struct{}

var _ models.JobManager = (*JobManager)(nil)

func NewJobManager() *JobManager { return &JobManager{} }

func (m *JobManager) List() ([]*models.Job, error) {
	var jobs []*models.Job
	if err := db.Find(&jobs).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return jobs, nil
}

func (m *JobManager) Create(job *models.Job) error {
	return db.Create(job).Error
}

func (m *JobManager) Get(id string) (*models.Job, error) {
	var job *models.Job
	if err := db.Find(&job, id).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return job, nil
}

func (m *JobManager) Update(id string, job *models.Job) error {
	return nil
}

func (m *JobManager) Delete(id string) error {
	return db.Delete(&models.Job{}, id).Error
}

func (m *JobManager) Cancel(id string) error {
	return nil
}
