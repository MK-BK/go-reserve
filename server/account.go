package server

import (
	"context"
	"crypto/md5"
	"fmt"
	"go-reserve/models"
	"go-reserve/server/audit"
)

type AccountManager struct{}

var _ models.AccountManager = (*AccountManager)(nil)

func NewAccountManager() *AccountManager { return &AccountManager{} }

func (m *AccountManager) Login(body *models.Account) (string, error) {
	defer func() {
		audit.NewAuditLog().
			WithAction(audit.ActionAccountLogin).
			Notify()
	}()

	password := fmt.Sprintf("%x", md5.Sum([]byte(body.Password)))

	var account models.Account
	if err := db.Where("name=?", body.Name).Where("password=?", password).First(&account).Error; err != nil {
		return "", err
	}

	if account.Enable {
		return "", ErrorAccountEnable
	}

	return "", nil
}

func (m *AccountManager) Create(account *models.Account) error {
	defer func() {
		audit.NewAuditLog().
			WithAction(audit.ActionAccountRegister).
			InContext(context.Background()).
			Notify()
	}()

	if account.Auth == models.AccountAdmin {
	}

	// if account.Auth == models.AccountProductAdmin {
	// 	Env.RequestManager.Create(&models.Request{
	// 		Name:  account.Name,
	// 		Email: account.Email,
	// 	})
	// 	return nil
	// }

	has := md5.Sum([]byte(account.Password))
	account.Password = fmt.Sprintf("%x", has)

	return db.Create(account).Error
}

func (m *AccountManager) List() ([]*models.Account, error) {
	var accounts []*models.Account
	if err := db.Find(&accounts).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return accounts, nil
}

func (m *AccountManager) Get(id string) (*models.Account, error) {
	var account models.Account

	if err := db.Find(&account, id).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return &account, nil
}

func (m *AccountManager) Update(id string, user *models.Account) error {
	return nil
}

func (m *AccountManager) Delete(id string) error {
	defer func() {
		audit.NewAuditLog().
			WithAction(audit.ActionAccountDelete).
			Notify()
	}()
	return db.Delete(&models.Account{}, id).Error
}
