//go:build wireinject
// +build wireinject

package main

import (
	"17-di/product"
	"database/sql"
	"github.com/google/wire"
)

func NewUseCase(db *sql.DB) *product.UseCase {
	wire.Build(product.NewUseCase, product.NewRepository)
	return &product.UseCase{}
}
