package domain

import "time"

type User struct {
	ID        int64     `gorm:"column:id;primaryKey"`
	Email     string    `gorm:"column:email"`
	Nik       string    `gorm:"column:nik"`
	Password  string    `gorm:"column:password"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy string    `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy string    `gorm:"column:updated_by"`
}

func (User) TableName() string {
	return "users"
}

type UserRole struct {
	ID        int64     `gorm:"column:id;primaryKey"`
	UserID    int64     `gorm:"column:user_id"`
	RoleID    int       `gorm:"column:role_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy string    `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy string    `gorm:"column:updated_by"`
}

func (UserRole) TableName() string {
	return "user_role"
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	RoleID         int    `json:"roleID"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	NIK            string `json:"nik"`
	Password       string `json:"password"`
	ReTypePassword string `json:"retype_password"`
}

type LoginDataResponse struct {
	Token string   `json:"token"`
	Menus []string `json:"menus"`
}
