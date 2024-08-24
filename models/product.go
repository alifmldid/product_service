package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryID  uuid.UUID `json:"category_id"`
}

type Product struct {
	ID          uuid.UUID       `gorm:"primaryKey"`
	Name        string          `gorm:"size:255"`
	Description string          `gorm:"type:text"`
	CategoryID  uuid.UUID       `gorm:"index"`
	Category    ProductCategory `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time
}

type ProductCategory struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string    `gorm:"size:255"`
}
