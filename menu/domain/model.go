package domain

import "time"

type Menu struct {
	ID          string    `gorm:"column:id;primaryKey"`
	MenuName    string    `gorm:"column:menu_name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	CreatedBy   string    `gorm:"column:created_by"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	UpdatedBy   string    `gorm:"column:updated_by"`
}

func (Menu) TableName() string {
	return "menu"
}
