package produto

import (
	"errors"
	"fmt"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ProductType string  `json:"type"`
	Count       int     `json:"count"`
	Price       float64 `json:"price"`
}

var ListProdutos []Product

type Repository interface {
	GetAll() ([]Product, error)
	Create(id int, name, productType string, count int, price float64) (Product, error)
	LastID() int
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	DeleteProduct(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repository) GetAll() ([]Product, error) {
	return ListProdutos, nil
}

func (repository) Create(id int, name, productType string, count int, price float64) (Product, error) {
	p := Product{id, name, productType, count, price}
	ListProdutos = append(ListProdutos, p)
	return p, nil
}

func (repository) LastID() int {
	var id int = 1
	if len(ListProdutos) > 0 {
		id = ListProdutos[len(ListProdutos)-1].ID + 1
	}
	return id
}

func (repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	p := Product{id, name, productType, count, price}
	for i := range ListProdutos {
		if ListProdutos[i].ID == p.ID {
			ListProdutos[i] = p
			return ListProdutos[i], nil
		}
	}
	return Product{}, errors.New("ID não encontrado")
}

func (repository) UpdateName(id int, name string) (Product, error) {
	for i := range ListProdutos {
		if ListProdutos[i].ID == id {
			ListProdutos[i].Name = name
			return ListProdutos[i], nil
		}
	}
	return Product{}, fmt.Errorf("produto %d não encontrado", id)
}

func (repository) DeleteProduct(id int) error {
	for i := range ListProdutos {
		if ListProdutos[i].ID == id {
			ListProdutos = append(ListProdutos[:i], ListProdutos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("produto %d não encontrado", id)
}
