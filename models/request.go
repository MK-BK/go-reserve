package models

import "gorm.io/gorm"

type RequestManager interface {
	List() ([]*Request, error)
	Create(r *Request) error
	Get(id string) (*Request, error)
	Update(id string, r *Request) error
}

type Request struct {
	gorm.Model
	Name  string
	Email string
}
