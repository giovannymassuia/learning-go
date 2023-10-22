//go:build wireinject
// +build wireinject

package main

import (
	"17-di/product"
	"database/sql"
	"github.com/google/wire"
)

var setRepositoryDependency = wire.NewSet(
	product.NewRepository,
	wire.Bind(new(product.RepositoryInterface), new(*product.Repository)),
)

func NewUseCase(db *sql.DB) *product.UseCase {
	wire.Build(
		setRepositoryDependency,
		product.NewUseCase,
	)
	return &product.UseCase{}
}
