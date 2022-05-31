package models

import "gorm.io/gorm"

type JobManage interface {
	List() ([]*Job, error)
	Cancel(id string) error

	Create(job *Job) error
	Get(id string) (*Job, error)
	Update(id string, job *Job) error
	Delete(id string) error
}

type Job struct {
	gorm.Model
	UserID      string `json:"userID" binding:"required"`
	CommodityID string `json:"commodityID" binding:"required"`
	Status      string
}
