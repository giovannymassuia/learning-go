package main

import (
	"04-sqlc-transactions/internal/db"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := db.New(tx)
	err = fn(q) // execute function in transaction
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil { // rollback transaction
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

func (c *CourseDB) CreateCrouseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	//queries := db.New(dbConn)

	courseArgs := CourseParams{
		ID:          "course-1",
		Name:        "Course 1",
		Description: sql.NullString{String: "Course 1 Description", Valid: true},
	}
	categoryArgs := CategoryParams{
		ID:          "category-1",
		Name:        "Category 1",
		Description: sql.NullString{String: "Category 1 Description", Valid: true},
	}

	courseDB := NewCourseDB(dbConn)
	err = courseDB.CreateCrouseAndCategory(ctx, categoryArgs, courseArgs)
	if err != nil {
		panic(err)
	}

}
