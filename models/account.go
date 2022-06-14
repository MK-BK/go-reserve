package models

import "gorm.io/gorm"

const (
	AccountAdmin = iota
	AccountProductAdmin
	// AccountProductNormal
	AccountNormal
)

type AccountManager interface {
	Login(account *Account) (string, error)
	Create(account *Account) error
	List() ([]*Account, error)
	Get(id string) (*Account, error)
	Update(id string, user *Account) error
	Delete(id string) error
}

type Account struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Auth     int
	Email    string
	Enable   bool
	Status   string // status 存储在redis中
}
