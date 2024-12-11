package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	//err = db.AutoMigrate(&Product{})
	//if err != nil {
	//	panic("failed to migrate table")
	//}

	// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	var product Product
	err = db.First(&product, 1).Error
	// db.First(&product, 1) // find product with integer primary key
	//  db.First(&product, "code = ?", "D42") // find product with code D42
	if err != nil {
		log.Fatal("Failed to find product:", err)
	}
	fmt.Println(product)
}

func createProduct(db *gorm.DB, code string, price uint) {
	product := Product{Code: code, Price: price}

	err := db.Create(&product).Error
	if err != nil {
		log.Fatal("Failed to create product:", err)
	}
	log.Printf("Created product: %+v\n", product)
}
