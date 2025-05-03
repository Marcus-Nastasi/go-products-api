package controller

import (
	"net/http"

	"github.com/Marcus-Nastasi/go-products-api/model"
	"github.com/gin-gonic/gin"
)

type productController struct {
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, []model.Product{
		{ID: 10, Name: "Oil", Price: 77.48},
	})
}
