package main

import (
	"crud-go/controller"
	"crud-go/db"
	"crud-go/repository"
	"crud-go/use_case"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}

	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUseCase := use_case.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)

	server.POST("/products", ProductController.CreateProduct)
	server.GET("/products", ProductController.GetAllProducts)
	server.GET("/products/:id", ProductController.GetProductById)

	server.Run(":8080")

}
