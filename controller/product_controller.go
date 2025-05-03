package controller

import (
	"github.com/Marcus-Nastasi/go-products-api/usecases"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUseCase usecases.ProductUseCase
}

func NewProductController(usecase usecases.ProductUseCase) ProductController {
	return ProductController{
		productUseCase: usecase,
	}
}

func (p *ProductController) GetProducts(c *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		panic("error getting products")
	}
	c.JSON(200, products)
}
