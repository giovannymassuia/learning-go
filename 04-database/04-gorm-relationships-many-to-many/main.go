package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primary_key"`
	Name     string
	Products []Product `gorm:"many2many:product_categories;"`
}

type Product struct {
	ID         int `gorm:"primary_key"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:product_categories;"`
	gorm.Model
}

func main() {
	dsn := "docker:docker@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// Create Category
	//category := Category{Name: "Kitchen"}
	//db.Create(&category)
	//category2 := Category{Name: "Home"}
	//db.Create(&category2)

	// Create Product with Categories
	//product := Product{Name: "Knife", Price: 2000, Categories: []Category{category, category2}}
	//db.Create(&product)

	var categories []Category
	db.Preload("Products").Find(&categories)
	for _, category := range categories {
		println(category.Name)
		for _, product := range category.Products {
			println("-> ", product.Name)
		}
	}

	println("=====================================")

	var products []Product
	db.Preload("Categories").Find(&products)
	for _, product := range products {
		println(product.Name)
		for _, category := range product.Categories {
			println("-> ", category.Name)
		}
	}
}
