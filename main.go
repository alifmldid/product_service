package main

import (
	"products/database"
	"products/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	db := database.Connect()

	// Set up the router
	router := gin.Default()

	// Routes
	router.POST("/products", handlers.CreateProduct(db))
	router.GET("/products", handlers.GetProducts(db))

	// Run the server
	router.Run(":8080")
}
