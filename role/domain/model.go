package domain

import "time"

type Role struct {
	ID        int       `gorm:"column:id;primaryKey"`
	RoleName  string    `gorm:"column:role_name"`
	Eselon    string    `gorm:"column:eselon"`
	OPD       string    `gorm:"column:opd"`
	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy string    `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy string    `gorm:"column:updated_by"`
}

func (Role) TableName() string {
	return "role"
}

type RoleMenu struct {
	ID        int64     `gorm:"column:id;primaryKey"`
	RoleID    int64     `gorm:"column:role_id"`
	MenuID    string    `gorm:"column:menu_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy string    `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy string    `gorm:"column:updated_by"`
}

func (RoleMenu) TableName() string {
	return "role_menu"
}
