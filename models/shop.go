package models

import "gorm.io/gorm"

type ShopManager interface {
	List() ([]*Shop, error)

	Create(product *Shop) error
	Get(id string) (*Shop, error)
	Update(id string, product *Shop) error
	Delete(id string) error
}

type Shop struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Description string `json:"description"`
}
