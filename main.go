package main

import (
	"net/http"

	"github.com/GolangJavierRestApi/database"
	"github.com/GolangJavierRestApi/product"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaseConnection := database.InitDB()

	defer databaseConnection.Close()

	var productRepository = product.NewRepository(databaseConnection)
	var productsService product.Service
	productsService = product.NewService(productRepository)
	r := chi.NewRouter()
	r.Mount("/products", product.MakeHttpHandLer(productsService))

	http.ListenAndServe(":3000", r)
}
