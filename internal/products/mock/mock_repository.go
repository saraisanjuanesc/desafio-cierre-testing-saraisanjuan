package mock

import (
	"errors"

	"desafio-cierre-testing-saraisanjuan/internal/products/domain"
)

type MockRepository struct {
	DataMock []domain.Product
	ErrRead  string
	ErrWrite string
}

func (m *MockRepository) GetAllBySeller(sellerID string) ([]domain.Product, error) {
	var results []domain.Product
	if m.ErrRead != "" {
		return []domain.Product{}, errors.New(m.ErrRead)
	}
	for _, value := range m.DataMock {
		if value.SellerID == sellerID {
			results = append(results, value)
		}
	}
	return results, nil
}
