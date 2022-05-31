package server

import (
	"context"
	"go-reserve/models"
	"go-reserve/server/audit"
)

type UserManager struct{}

var _ models.UserManaer = (*UserManager)(nil)

func NewUserManager() *UserManager { return &UserManager{} }

func (m *UserManager) Login(body *models.User) (string, error) {
	var user *models.User
	if err := _db.Find(&user, body.Name, body.PassWord).Error; err != nil {
		return "", err
	}

	// TODO
	if user.Enable {
		// not allow to login
		return "", nil
	}
	return "", nil
}

func (m *UserManager) List() ([]*models.User, error) {
	var users []*models.User
	if err := _db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (m *UserManager) Get(id string) (*models.User, error) {
	var user models.User

	if err := _db.Find(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *UserManager) Create(user *models.User) error {
	defer func() {
		audit.NewAuditLog().
			InContext(context.Background()).
			Notify()
	}()

	if user.Auth == models.UserAdmin {
		return nil
	}

	if user.Auth == models.UserProductAdmin {
		Env.RequestManager.Create(&models.Request{
			Name:  user.Name,
			Email: user.Email,
		})
		return nil
	}

	return _db.Create(user).Error
}

func (m *UserManager) Update(id string, user *models.User) error {
	return nil
}

func (m *UserManager) Delete(id string) error {
	return _db.Delete(&models.User{}, id).Error
}
