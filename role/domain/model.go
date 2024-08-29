package domain

import "time"

type Role struct {
	ID        int64     `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name"`
	Echelon   string    `gorm:"column:echelon"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Role) TableName() string {
	return "roles"
}

type CreateRoleResponse struct {
	RoleId int64 `json:"roleId"`
}
