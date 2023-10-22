package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file::memory")
	if err != nil {
		panic(err)
	}

	useCase := NewUseCase(db)

	getProduct, err := useCase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	println(getProduct.Name)
}
