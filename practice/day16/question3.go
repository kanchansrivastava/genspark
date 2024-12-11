package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"fmt"
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
	err = db.AutoMigrate(&Product{}) // created table product
	if err != nil {
		panic("failed to migrate table")
	}

	code := "D423"
	createProduct(db, code, 100) // insert a record
	getProduct(db, 1) // fetch the record with id = 1
	updateProduct(db, code) // updates the record with code
	deleteProduct(db, "D423") // deletes the record
}

func createProduct(db *gorm.DB, code string, price uint) {
	product := Product{Code: code, Price: price}

	err := db.Create(&product).Error
	if err != nil {
		log.Fatal("Failed to create product:", err)
	}
	log.Printf("Created product: %+v\n", product)
}


func updateProduct(db *gorm.DB, code string) {
	result := db.Model(&Product{}).Where("code = ?", code).Update("Price", 150)

	if result.Error != nil {
		log.Fatal("Failed to update product:", result.Error)
	}

	if result.RowsAffected == 0 {
		log.Printf("No product found with code: %s\n", code)
		return
	}

	log.Printf("Updated product with code: %s\n", code)
}


func deleteProduct(db *gorm.DB, code string) {
	result := db.Where("code = ?", code).Delete(&Product{}) // to delete with id err := db.Delete(&Product{}, id)
	if result.Error != nil {
		log.Fatal("Failed to delete product:", result.Error)
	}
	log.Printf("Deleted product: with code %+v\n", code)
}

func getProduct(db *gorm.DB, id int) {
	var product Product
	err := db.First(&product, id).Error
	// db.First(&product, 1) // find product with integer primary key
	//  db.First(&product, "code = ?", "D42") // find product with code D42
	if err != nil {
		log.Fatal("Failed to find product:", err)
	}
	fmt.Println(product)
}