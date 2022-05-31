package models

import "gorm.io/gorm"

const (
	UserAdmin = iota
	UserProductAdmin
	UserProductNormal
	UserNormal
)

type UserManaer interface {
	Login(user *User) (string, error)
	List() ([]*User, error)
	Get(id string) (*User, error)
	Create(user *User) error
	Update(id string, user *User) error
	Delete(id string) error
}

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	PassWord string `json:"password" binding:"required"`
	Auth     int    `json:"auth" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Enable   bool
}
