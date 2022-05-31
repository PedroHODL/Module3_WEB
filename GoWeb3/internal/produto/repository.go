package produto

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
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repository) GetAll() ([]Product, error) {
	return ListProdutos, nil
}

func (r repository) Create(id int, name, productType string, count int, price float64) (Product, error) {
	p := Product{id, name, productType, count, price}
	p.ID = r.LastID()
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
