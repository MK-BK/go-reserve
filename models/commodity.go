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
	Price       int    `json:"price" binding:"required"`
	ShopID      int    `json:"shopID" binding:"required"`
	Description string
}
