package controller

import (
	"net/http"

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
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
