package main

import (
	"03-sqlc/internal/db"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"time"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// create new category
	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "New Category",
		Description: sql.NullString{String: "New Category Description", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	// list all categories
	allCategories := listAllCategories(queries, ctx)

	// update category
	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          allCategories[0].ID,
		Name:        "Updated Category " + time.Now().String(),
		Description: sql.NullString{String: "Updated Category Description", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	// list all categories
	listAllCategories(queries, ctx)

	// delete category
	err = queries.DeleteCategory(ctx, allCategories[0].ID)
	if err != nil {
		panic(err)
	}

	// list all categories
	listAllCategories(queries, ctx)
}

func listAllCategories(queries *db.Queries, ctx context.Context) []db.Category {
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Printf("%+v\n", category)
	}

	if len(categories) == 0 {
		fmt.Println("No categories found")
	}

	return categories
}
