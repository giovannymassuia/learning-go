package main

import (
	"database/sql"
	"fmt"

	// the '_' is used to import the package without using it
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "docker:docker@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Go", 20)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product.Price = 30
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	product, err = selectOneProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %+v\n", product)

	fmt.Printf("Finding all products...\n")
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		fmt.Printf("Product: %+v\n", product)
	}

	fmt.Printf("Deleting product with ID: %s\n", product.ID)
	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done!")
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted product with ID: %s\n", product.ID)

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Printf("Last insert ID: %d\n", lastInsertID)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("Rows affected: %d\n", rowsAffected)

	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name=?, price=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("Rows affected: %d\n", rowsAffected)

	return nil
}

func selectOneProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product Product
	// parameters are passed in the same order as select statement
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func selectAllProducts(db *sql.DB) ([]*Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("Rows affected: %d\n", rowsAffected)

	return nil
}
