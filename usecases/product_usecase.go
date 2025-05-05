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

func (pu *ProductUseCase) GetProduct(id string) (model.Product, error) {
	product, err := pu.repository.GetProduct(id)
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}
	return product, nil
}

func (pu *ProductUseCase) Create(new_product model.Product) (model.Product, error) {
	product, err := pu.repository.Create(new_product)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (pu *ProductUseCase) UpdateProduct(id string, product model.Product) (model.Product, error) {
	product, err := pu.repository.UpdateProduct(id, product)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (pu *ProductUseCase) DeleteById(id string) (model.Product, error) {
	product, err := pu.repository.DeleteById(id)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}
