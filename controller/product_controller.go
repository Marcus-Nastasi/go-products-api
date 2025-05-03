package controller

import (
	"net/http"

	"github.com/Marcus-Nastasi/go-products-api/model"
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

func (p *ProductController) Create(c *gin.Context) {
	var new_product model.Product
	err := c.BindJSON(&new_product)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	id, err := p.productUseCase.Create(new_product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, map[string]int{"product": id})
}
