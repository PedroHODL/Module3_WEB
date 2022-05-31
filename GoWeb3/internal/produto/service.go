package produto

type Services interface {
	GetAll() ([]Product, error)
	Create(name, productType string, count int, price float64) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Services {
	s := service{r}
	return &s
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) Create(name, productType string, count int, price float64) (Product, error) {
	id := s.repository.LastID()
	ps, err := s.repository.Create(id, name, productType, count, price)
	if err != nil {
		return Product{}, err
	}
	return ps, nil
}
