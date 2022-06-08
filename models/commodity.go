package models

import "gorm.io/gorm"

type CommodityManager interface {
	List() ([]*Commodity, error)
	Create(c *Commodity) error
	Get(id string) (*Commodity, error)
	Update(id string, c *Commodity) error
	Delete(id string) error
}

type Commodity struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	ShopID      uint   `json:"shopID" binding:"required"`
	Description string
}
