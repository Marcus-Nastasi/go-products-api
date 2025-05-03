package usecases

import (
	"fmt"

	"github.com/Marcus-Nastasi/go-products-api/model"
	"github.com/Marcus-Nastasi/go-products-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repository repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repository,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	products, err := pu.repository.GetProducts()
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}
	return products, nil
}

func (pu *ProductUseCase) Create(new_product model.Product) (int, error) {
	id, err := pu.repository.Create(new_product)
	if err != nil {
		return 0, err
	}
	return id, nil
}
