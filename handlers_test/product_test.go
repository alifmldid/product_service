package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"products/handlers"
	"products/models"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	// Setup in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.Product{})

	// Generate a valid UUID
	categoryID := uuid.New().String()

	// Simulate Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Simulate a POST request
	jsonBody := `{"name": "Product1", "description": "Description1", "category_id": "` + categoryID + `"}`
	c.Request = httptest.NewRequest("POST", "/products", strings.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the handler
	handler := handlers.CreateProduct(db)
	handler(c)

	// Check response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, w.Code)
	}

	// Optionally, validate the JSON response
}

func TestGetProducts(t *testing.T) {
	// Setup in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.Product{})

	// Generate valid UUIDs
	product1ID := uuid.New()
	product2ID := uuid.New()
	categoryID := uuid.New()

	// Insert test data
	products := []models.Product{
		{
			ID:          product1ID,
			Name:        "Product1",
			Description: "Description1",
			CategoryID:  categoryID,
			CreatedAt:   time.Now().Add(-time.Hour),
		},
		{
			ID:          product2ID,
			Name:        "Product2",
			Description: "Description2",
			CategoryID:  categoryID,
			CreatedAt:   time.Now(),
		},
	}

	for _, product := range products {
		if err := db.Create(&product).Error; err != nil {
			t.Fatalf("Failed to insert product: %v", err)
		}
	}

	// Simulate Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Simulate GET request with no filters
	c.Request = httptest.NewRequest("GET", "/products", nil)

	// Call the handler
	handler := handlers.GetProducts(db)
	handler(c)

	// Check response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, w.Code)
	}

	// Optionally, validate the JSON response
}
