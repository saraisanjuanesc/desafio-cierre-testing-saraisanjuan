package products

import (
	"log"

	"desafio-cierre-testing-saraisanjuan/internal/products/domain"
)

type Service interface {
	GetAllBySeller(sellerID string) ([]domain.Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAllBySeller(sellerID string) ([]domain.Product, error) {
	data, err := s.repo.GetAllBySeller(sellerID)
	if err != nil {
		log.Println("error in repository", err.Error(), "sellerId:", sellerID)
		return nil, err
	}
	return data, err
}
