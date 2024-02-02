package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/David-Ackerman/imersao-full-cycle/goapi/internal/database"
	"github.com/David-Ackerman/imersao-full-cycle/goapi/internal/service"
	"github.com/David-Ackerman/imersao-full-cycle/goapi/internal/webserver"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersaoFullCycle")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	categoryDb := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDb)

	productDb := database.NewProductDB(db)
	productService := service.NewProductService(*productDb)

	webCategoryHandler := webserver.NewCategoryHandler(categoryService)
	webProductHandler := webserver.NewProductHandler(productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Get("/category", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)

	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/product", webProductHandler.GetProducts)
	c.Get("/product/category/{categoryID}", webProductHandler.GetProductByCategoryID)
	c.Post("/product", webProductHandler.CreateProduct)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)

}
