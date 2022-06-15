package models

import "gorm.io/gorm"

type JobManager interface {
	List() ([]*Job, error)
	Create(job *Job) error
	Get(id string) (*Job, error)
	Update(id string, job *Job) error
	Delete(id string) error
	Cancel(id string) error
}

type Job struct {
	gorm.Model
	UserID      string `json:"userID" binding:"required"`
	CommodityID string `json:"commodityID" binding:"required"`
	Number      int
	Price       int
	Status      string
}
