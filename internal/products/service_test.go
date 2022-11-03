package products

import (
	"errors"
	"testing"

	"desafio-cierre-testing-saraisanjuan/internal/products/domain"
	"desafio-cierre-testing-saraisanjuan/internal/products/mock"

	"github.com/stretchr/testify/assert"
)

var data = []domain.Product{
	{
		ID:          "mock1",
		SellerID:    "FAX89",
		Description: "generic product",
		Price:       982.75,
	},
	{
		ID:          "mock2",
		SellerID:    "NCJSA2",
		Description: "generic product",
		Price:       723.34,
	},
	{
		ID:          "mock3",
		SellerID:    "FAX89",
		Description: "generic product",
		Price:       1334.54,
	},
}

func Test_GetAllBySeller_Service(t *testing.T) {
	idSeller := "FAX89"
	productsExpected := []domain.Product{
		{
			ID:          "mock1",
			SellerID:    "FAX89",
			Description: "generic product",
			Price:       982.75,
		},
		{
			ID:          "mock3",
			SellerID:    "FAX89",
			Description: "generic product",
			Price:       1334.54,
		},
	}
	//resultExpectedErr := []domain.Product{}
	myMockR := mock.MockRepository{DataMock: data}
	service := NewService(&myMockR)

	results, err := service.GetAllBySeller(idSeller)

	assert.Nil(t, err)
	assert.Equal(t, productsExpected, results)

	//Fail
	myMockR.ErrRead = "ID not found"
	errorExpected := errors.New("ID not found")
	service = NewService(&myMockR)

	results, err = service.GetAllBySeller(idSeller)

	assert.NotNil(t, err)
	assert.Nil(t, results)
	assert.Equal(t, errorExpected, err)

}
