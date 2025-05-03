package main

import (
	"github.com/Marcus-Nastasi/go-products-api/controller"
	"github.com/Marcus-Nastasi/go-products-api/usecases"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	productUsecase := usecases.NewProductUseCase()
	ctr := controller.NewProductController(productUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"status": "pong"})
	})
	server.GET("/products", ctr.GetProducts)

	server.Run(":8000")
}
