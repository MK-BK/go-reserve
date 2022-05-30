package user

import (
	"fmt"
	"go-reserve/common/db"
	"go-reserve/models"
)

type manager struct{}

var _db = db.GetDB()

var _ models.UserManaer = (*manager)(nil)

func NewManger() *manager { return &manager{} }

func (m *manager) Login(userBody *models.User) error {
	var user *models.User
	if err := _db.Find(&user, userBody.Name, userBody.PassWord).Error; err != nil {
		return err
	}
	return nil
}

func (m *manager) Get(id string) (*models.User, error) {
	fmt.Println("++++++ uerID:", id)
	var user *models.User
	if err := _db.Find(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (m *manager) Create(user *models.User) error {
	return _db.Create(user).Error
}

func (m *manager) List() ([]*models.User, error) {
	var users []*models.User
	if err := _db.Find(users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (m *manager) Update(id string, user *models.User) error { return nil }

func (m *manager) Delete(id string) error {
	return _db.Delete(&models.User{}, id).Error
}
