package handlers

import (
	"net/http"
	"products/models"
	"products/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.ProductRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Sanitize input to prevent XSS
		request.Name = utils.SanitizeInput(request.Name)
		request.Description = utils.SanitizeInput(request.Description)

		if _, err := uuid.Parse(request.CategoryID.String()); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID for CategoryID"})
			return
		}

		productID := uuid.New()

		createProduct := models.Product{
			ID:          productID,
			Name:        request.Name,
			Description: request.Description,
			CategoryID:  request.CategoryID,
		}

		if err := db.Create(&createProduct).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, createProduct)
	}
}

func GetProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		var count int64
		query := db.Preload("Category")

		keyword := c.Query("keyword")
		categoryID := c.Query("category_id")

		if keyword != "" {
			query = query.Where("name LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
		}

		if categoryID != "" {
			query = query.Where("category_id = ?", categoryID)
		}

		// Pagination
		limit, offset := utils.GetPaginationParams(c)
		query = query.Offset(offset).Limit(limit).Order("created_at DESC")

		if err := query.Find(&products).Count(&count).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"products": products,
			"count":    count,
		})
	}
}
