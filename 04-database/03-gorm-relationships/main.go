package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primary_key"`
	Name     string
	Products []Product // has many
}

type Product struct {
	ID            int `gorm:"primary_key"`
	Name          string
	Price         float64
	CategoryID    int          // belongs to
	Category      Category     // has one
	SerialNumbers SerialNumber // has one
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primary_key"`
	Number    string
	ProductID int // belongs to
}

func main() {
	dsn := "docker:docker@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// Create Category
	category := Category{Name: "Mouses"}
	db.Create(&category)

	// Create Product
	product := Product{Name: "Mouse Classic", Price: 2000, CategoryID: category.ID}
	db.Create(&product)

	// Create Serial Number
	serialNumber := SerialNumber{Number: "123456789", ProductID: product.ID}
	db.Create(&serialNumber)

	// Read Product
	//var products []Product
	//db.Preload("Category").Preload("SerialNumbers").Find(&products)
	//for _, product := range products {
	//	fmt.Println(product.Name, product.Price, product.Category, product.SerialNumbers)
	//}

	// Read Categories with Products
	//var categories []Category
	//err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	//if err != nil {
	//	panic(err)
	//}
	//for _, category := range categories {
	//	fmt.Println(category.Name)
	//	for _, product := range category.Products {
	//		fmt.Println("--", product.Name, product.Price)
	//	}
	//}

	// Read Categories with Products and Serial Number
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumbers").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name)
		for _, product := range category.Products {
			fmt.Println("--", product.Name, product.Price, "Serial Numbers:", product.SerialNumbers.Number)
		}
	}
}
