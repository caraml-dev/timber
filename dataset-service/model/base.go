package model

import "time"

// Base is base model
type Base struct {
	// Unique ID of the entity
	ID int64 `gorm:"primaryKey,autoIncrement"`
	// Project ID that own the entity
	ProjectID int64
	// CreatedAt creation timestamp
	CreatedAt time.Time
	// UpdatedAt last update timestamp
	UpdatedAt time.Time
}
