package server

import "go-reserve/models"

type commodityManaer struct{}

var _ models.CommodityManager = (*commodityManaer)(nil)

func NewCommodityManager() *commodityManaer {
	return &commodityManaer{}
}

func (m *commodityManaer) List() ([]*models.Commodity, error) {
	var cs []*models.Commodity
	if err := _db.Find(&cs).Error; err != nil {
		return nil, err
	}
	return cs, nil
}

func (m *commodityManaer) Create(c *models.Commodity) error {
	return _db.Create(&c).Error
}

func (m *commodityManaer) Get(id string) (*models.Commodity, error) {
	var commodity *models.Commodity
	if err := _db.Find(&commodity, id).Error; err != nil {
		return nil, err
	}
	return commodity, nil
}

func (m *commodityManaer) Update(id string, c *models.Commodity) error {
	return nil
}

func (m *commodityManaer) Delete(id string) error {
	return _db.Delete(&models.Commodity{}, id).Error
}