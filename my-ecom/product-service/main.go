package main

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	// "product-service/internal/stores/postgres"


	// 	/*
	// 		//------------------------------------------------------//
	// 	                Setting up DB & Migrating tables
	// 		//------------------------------------------------------//
	// */

	// slog.Info("Migrating tables for user-service if not already done")
	// db, err := postgres.OpenDB()
	// if err != nil {
	// 	return err
	// }
	// defer db.Close()
	// err = postgres.RunMigrations(db)
	// if err != nil {
	// 	return err
	// }

	// port := os.Getenv("PORT")

	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ping successful"})
	})

	if err := r.Run(":" + os.Getenv("APP_PORT")); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
