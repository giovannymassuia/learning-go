package main

import (
	"github.com/giovannymassuia/learning-go/07-api/configs"
	"github.com/giovannymassuia/learning-go/07-api/internal/entity"
	"github.com/giovannymassuia/learning-go/07-api/internal/infra/database"
	"github.com/giovannymassuia/learning-go/07-api/internal/infra/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)

	http.ListenAndServe(":"+config.WebServerPort, r)
}
