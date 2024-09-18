package db

import (
	"github.com/google/uuid"
	"gopkg.in/jinzhu/gorm.v1"
	"time"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primary_key;"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy string    `gorm:"column:created_by" json:"created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy string    `gorm:"column:updated_by" json:"updated_by"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}
