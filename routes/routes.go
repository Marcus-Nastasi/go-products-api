package routes

import (
	"github.com/Marcus-Nastasi/go-products-api/controller"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, controller controller.ProductController) {
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"status": "pong"})
	})

	products := router.Group("/products")
	{
		products.GET("", controller.GetProducts)
	}
}
