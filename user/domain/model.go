package domain

import (
	comdb "com.mailnau.api/common/db"
	_rma_domain "com.mailnau.api/role-menu-action/domain"
	"time"
)

type User struct {
	Email    string `gorm:"column:email"`
	Nik      string `gorm:"column:nik"`
	Password string `gorm:"column:password"`
	Status   string `gorm:"column:status"`
	RoleID   int    `gorm:"column:role_id"`
	comdb.Base
}

func (User) TableName() string {
	return "users"
}

// DEPRECATED
type UserRole struct {
	ID        int64     `gorm:"column:id;primaryKey"`
	UserID    string    `gorm:"column:user_id"`
	RoleID    int       `gorm:"column:role_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy string    `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy string    `gorm:"column:updated_by"`
}

func (UserRole) TableName() string {
	return "user_role"
}

type LoginByEmail struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginByNIK struct {
	NIK      string `json:"nik" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	RoleID         int    `json:"roleID" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email" validate:"required"`
	NIK            string `json:"nik" validate:"required"`
	Password       string `json:"password" validate:"required"`
	ReTypePassword string `json:"retype_password" validate:"required"`
}

type LoginDataResponse struct {
	Token       string                       `json:"token"`
	TokenType   string                       `json:"token_type"`
	ExpiresIn   string                       `json:"expires_in"`
	UserID      string                       `json:"user_id"`
	Status      string                       `json:"status"`
	MenuActions []_rma_domain.RoleMenuAction `json:"menu_actions"`
}
