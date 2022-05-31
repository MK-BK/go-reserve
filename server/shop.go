package server

import (
	"go-reserve/models"
)

type ProductManager struct{}

func NewShopManager() *ProductManager { return &ProductManager{} }

var _ models.ShopManager = (*ProductManager)(nil)

func (m *ProductManager) List() ([]*models.Shop, error) {
	var products []*models.Shop
	if err := _db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (m *ProductManager) Create(product *models.Shop) error {
	return _db.Create(&product).Error
}

func (m *ProductManager) Get(id string) (*models.Shop, error) {
	var product *models.Shop
	if err := _db.Find(&product, id).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (m *ProductManager) Update(id string, product *models.Shop) error {
	return nil
}

func (m *ProductManager) Delete(id string) error {
	return _db.Delete(&models.Shop{}, id).Error
}
