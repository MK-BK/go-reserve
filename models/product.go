package models

type ProductManage interface {
	Create(*Product) error
	Get(id string) (*Product, error)
	Update(id string, product *Product) error
	List() ([]*Product, error)
	Delete(id string) error
}

type Product struct {
	ID      uint
	Name    string
	Address string
}
