package product

import "database/sql"

type RepositoryInterface interface {
	GetProduct(id int) (Product, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetProduct(id int) (Product, error) {
	return Product{
		ID:   id,
		Name: "Product " + string(id),
	}, nil
}
