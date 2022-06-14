package server

import (
	"go-reserve/models"
	"go-reserve/server/audit"
)

type commodityManaer struct{}

var _ models.CommodityManager = (*commodityManaer)(nil)

func NewCommodityManager() *commodityManaer {
	return &commodityManaer{}
}

func (m *commodityManaer) List() ([]*models.Commodity, error) {
	var cs []*models.Commodity
	if err := db.Find(&cs).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return cs, nil
}

func (m *commodityManaer) Create(c *models.Commodity) error {
	defer func() {
		audit.NewAuditLog().
			WithAction(audit.ActionCommodityCreate).
			Notify()
	}()

	return db.Create(&c).Error
}

func (m *commodityManaer) Get(id string) (*models.Commodity, error) {
	var commodity *models.Commodity
	if err := db.Find(&commodity, id).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return commodity, nil
}

func (m *commodityManaer) Update(id string, c *models.Commodity) error {
	defer func() {
		audit.NewAuditLog().
			WithAction(audit.ActionCommodityUpdate).
			Notify()
	}()

	return nil
}

func (m *commodityManaer) Delete(id string) error {
	defer func() {
		audit.NewAuditLog().
			WithAction(audit.ActionCommodityDelete).
			Notify()
	}()

	return db.Delete(&models.Commodity{}, id).Error
}
