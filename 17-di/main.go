package main

import (
	"17-di/product"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file::memory")
	if err != nil {
		panic(err)
	}

	repository := product.NewRepository(db)

	useCase := product.NewUseCase(repository)

	getProduct, err := useCase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	println(getProduct.Name)
}
