package models

import "gorm.io/gorm"

var (
	StatusNew     = "new"
	StatusApprove = "approve"
	StatusDeny    = "deny"
)

type RequestManager interface {
	List() ([]*Request, error)
	Create(r *Request) error
	Get(id string) (*Request, error)
	Update(id string, r *Request) error
}

type Request struct {
	gorm.Model
	Name   string
	Email  string
	Status string // 管理员审核状态
}
