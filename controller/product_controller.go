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

func (p *ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")
	products, err := p.productUseCase.GetProduct(id)
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
	product, err := p.productUseCase.Create(new_product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var updated model.Product
	err := c.BindJSON(&updated)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	product, err := pc.productUseCase.UpdateProduct(id, updated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, product)
}

func (p *ProductController) DeleteById(c *gin.Context) {
	id := c.Param("id")
	product, err := p.productUseCase.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}
