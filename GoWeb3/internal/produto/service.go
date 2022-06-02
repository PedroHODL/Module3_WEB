package produto

type Services interface {
	GetAll() ([]Product, error)
	Create(name, productType string, count int, price float64) (Product, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	DeleteProduct(id int) error
	LastID() int
}

type service struct {
	repository Repository
}

func NewService(r Repository) Services {
	s := service{r}
	return &s
}

func (s service) LastID() int {
	return s.repository.LastID()
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) Create(name, productType string, count int, price float64) (Product, error) {
	id := s.repository.AvailableID()
	ps, err := s.repository.Create(id, name, productType, count, price)
	if err != nil {
		return Product{}, err
	}
	return ps, nil
}

func (s *service) Update(id int, name, productType string, count int, price float64) (Product, error) {
	ps, err := s.repository.Update(id, name, productType, count, price)
	if err != nil {
		return Product{}, err
	}
	return ps, nil
}

func (s *service) UpdateName(id int, name string) (Product, error) {
	ps, err := s.repository.UpdateName(id, name)
	if err != nil {
		return Product{}, err
	}

	return ps, nil
}

func (s *service) DeleteProduct(id int) error {
	err := s.repository.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
