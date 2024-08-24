package tests

import (
	"net/http"
	"net/http/httptest"
	"products/database"
	"products/handlers"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationCreateProduct(t *testing.T) {
	// Inisialisasi database
	db := database.Connect()

	// Setup Gin context dan handler
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/products", handlers.CreateProduct(db))

	// Buat permintaan HTTP
	reqBody := `{"name": "Product1", "description": "Description1", "category_id": "123e4567-e89b-12d3-a456-426614174001"}`
	req, _ := http.NewRequest("POST", "/products", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Lakukan permintaan
	router.ServeHTTP(w, req)

	// Periksa respons
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Product1")
}

func TestIntegrationGetProducts(t *testing.T) {
	// Inisialisasi database
	db := database.Connect()

	// Setup Gin context dan handler
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/products", handlers.GetProducts(db))

	// Buat permintaan HTTP
	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()

	// Lakukan permintaan
	router.ServeHTTP(w, req)

	// Periksa respons
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "products")
}
