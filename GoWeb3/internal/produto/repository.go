package produto

import (
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
	AvailableID() int
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

	for i := range ListProdutos {
		if ListProdutos[i].ID+1 == id {
			post := make([]Product, len(ListProdutos[i+1:]))
			copy(post, ListProdutos[i+1:])
			ListProdutos = append(ListProdutos[:i+1], p)
			ListProdutos = append(ListProdutos, post...)
			break
		}
	}

	if id == 1 {
		prod := []Product{p}
		ListProdutos = append(prod, ListProdutos...)
	}
	r.db.Write(ListProdutos)
	return p, nil
}

func (r repository) AvailableID() int {
	var ListProdutos []Product
	r.db.Read(&ListProdutos)

	for prevI := range ListProdutos[:len(ListProdutos)-1] {
		i := prevI + 1
		if ListProdutos[i].ID != (ListProdutos[prevI].ID + 1) {
			id := ListProdutos[prevI].ID + 1
			return id
		}
	}
	return r.LastID()
}

func (r repository) LastID() int {
	var ListProdutos []Product
	r.db.Read(&ListProdutos)

	if len(ListProdutos) == 0 || ListProdutos[0].ID != 1 {
		return 1
	}
	return ListProdutos[len(ListProdutos)-1].ID + 1
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
	return Product{}, fmt.Errorf("produto %d não encontrado", id)
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
