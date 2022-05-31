package blabla

import (
	"github.com/GoWeb2&3-Arquitecture/internal/produto/repository"
)

type Services interface {
	GetAll() ([]Product, error)
	Create(id int, name, productType string, count int, price float64) (Product, error)
}

type service struct {
	repository Repository
}

func newService(r Repository) Services {
	s := service{r}
	return &s
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
}
