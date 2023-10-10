package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Price float64
}

func main() {
	dsn := "docker:docker@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// Create
	//db.Create(&Product{Name: "Go", Price: 20})

	// Create in batch
	//db.Create([]Product{
	//	{Name: "Python", Price: 30},
	//	{Name: "JavaScript", Price: 40},
	//})

	// Select One
	//var product Product
	//db.First(&product, 1)
	//db.First(&product, "name = ?", "Go")
	//fmt.Println(product)

	// Select All
	//var products []Product
	//db.Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	// Select All with limit and offset
	//var products []Product
	//db.Limit(2).Offset(2).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	// Select All with where
	//var products []Product
	//db.Where("price > ?", 30).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	// Update
	var product Product
	db.First(&product, 1)
	product.Price = 50
	db.Save(&product)

	var updatedProduct Product
	db.First(&updatedProduct, 1)
	fmt.Println(updatedProduct)

	// Delete
	db.Delete(&product)

}
