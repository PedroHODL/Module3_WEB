package produto

import (
	"errors"
	"fmt"

	"github.com/PedroHODL/Module3_WEB.git/GoWeb3/pkg/store"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ProductType string  `json:"type"`
	Count       int     `json:"count"`
	Price       float64 `json:"price"`
}
type Repository interface {
	GetAll() ([]Product, error)
	Create(id int, name, productType string, count int, price float64) (Product, error)
	LastID() int
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	DeleteProduct(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r repository) GetAll() ([]Product, error) {
	var ListProdutos []Product
	r.db.Read(&ListProdutos)
	return ListProdutos, nil
}

func (r repository) Create(id int, name, productType string, count int, price float64) (Product, error) {
	var ListProdutos []Product
	r.db.Read(&ListProdutos)

	p := Product{id, name, productType, count, price}
	ListProdutos = append(ListProdutos, p)

	r.db.Write(ListProdutos)
	return p, nil
}

func (r repository) LastID() int {
	var ListProdutos []Product
	r.db.Read(&ListProdutos)

	var id int = 0
	if len(ListProdutos) > 0 {
		id = ListProdutos[len(ListProdutos)-1].ID + 1
	}
	return id
}

func (r repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	var ListProdutos []Product
	r.db.Read(&ListProdutos)

	p := Product{id, name, productType, count, price}
	for i := range ListProdutos {
		if ListProdutos[i].ID == p.ID {
			ListProdutos[i] = p
			r.db.Write(ListProdutos)
			return ListProdutos[i], nil
		}
	}
	return Product{}, errors.New("ID não encontrado")
}

func (r repository) UpdateName(id int, name string) (Product, error) {
	var ListProdutos []Product
	r.db.Read(&ListProdutos)

	for i := range ListProdutos {
		if ListProdutos[i].ID == id {
			ListProdutos[i].Name = name
			r.db.Write(ListProdutos)
			return ListProdutos[i], nil
		}
	}
	return Product{}, fmt.Errorf("produto %d não encontrado", id)
}

func (r repository) DeleteProduct(id int) error {
	var ListProdutos []Product
	r.db.Read(&ListProdutos)

	for i := range ListProdutos {
		if ListProdutos[i].ID == id {
			ListProdutos = append(ListProdutos[:i], ListProdutos[i+1:]...)
			r.db.Write(ListProdutos)
			return nil
		}
	}
	return fmt.Errorf("produto %d não encontrado", id)
}
