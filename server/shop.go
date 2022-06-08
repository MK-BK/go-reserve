package server

import (
	"go-reserve/models"
)

type ShopManager struct{}

var _ models.ShopManager = (*ShopManager)(nil)

func NewShopManager() *ShopManager { return &ShopManager{} }

func (m *ShopManager) List() ([]*models.Shop, error) {
	var shops []*models.Shop
	if err := db.Find(&shops).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return shops, nil
}

func (m *ShopManager) Create(shop *models.Shop) error {
	return db.Create(&shop).Error
}

func (m *ShopManager) Get(id string) (*models.Shop, error) {
	var shop *models.Shop
	if err := db.Find(&shop, id).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return shop, nil
}

func (m *ShopManager) Update(id string, shop *models.Shop) error {
	return nil
}

func (m *ShopManager) Delete(id string) error {
	return db.Delete(&models.Shop{}, id).Error
}
