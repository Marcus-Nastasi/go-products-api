package main

import (
	"github.com/Marcus-Nastasi/go-products-api/controller"
	"github.com/Marcus-Nastasi/go-products-api/db"
	"github.com/Marcus-Nastasi/go-products-api/repository"
	"github.com/Marcus-Nastasi/go-products-api/routes"
	"github.com/Marcus-Nastasi/go-products-api/usecases"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	db, dbErr := db.ConnectDb()
	if dbErr != nil {
		panic("Error on database " + dbErr.Error())
	}
	productRepo := repository.NewProductRepository(db)
	productUsecase := usecases.NewProductUseCase(productRepo)
	productController := controller.NewProductController(productUsecase)

	routes.Setup(server, productController)

	server.Run(":8000")
}
