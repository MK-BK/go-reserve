package models

type UserManaer interface {
	Login(user *User) error
	List() ([]*User, error)
	Create(user *User) error
	Get(id string) (*User, error)
	Update(id string, user *User) error
	Delete(id string) error
}

type User struct {
	ID       uint32
	Name     string `json:"name" binding:"required"`
	PassWord string `json:"password" binding:"required"`
	Auth     int
}
