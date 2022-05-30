package product

import (
	"go-reserve/common/db"
	"go-reserve/models"
)

var _db = db.GetDB()

type manage struct{}

func NewManager() *manage { return &manage{} }

var _ models.ProductManage = (*manage)(nil)

func (m *manage) Create(product *models.Product) error {
	return _db.Create(&product).Error
}

func (m *manage) Get(id string) (*models.Product, error) {
	var product *models.Product
	if err := _db.Find(&product, id).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (m *manage) List() ([]*models.Product, error) {
	var products []*models.Product
	if err := _db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (m *manage) Update(id string, product *models.Product) error { return nil }

func (m *manage) Delete(id string) error {
	return _db.Delete(&models.Product{}, id).Error
}
