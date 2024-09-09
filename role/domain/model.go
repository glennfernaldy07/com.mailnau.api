package domain

import (
	"time"

	rma_domain "com.mailnau.api/role-menu-action/domain"
)

type Role struct {
	ID        int                 `gorm:"column:id;primaryKey"`
	RoleName  string              `gorm:"column:role_name"`
	Eselon    string              `gorm:"column:eselon"`
	OPD       string              `gorm:"column:opd"`
	Menus     []rma_domain.Menu   `gorm:"many2many:role_menu_action;joinForeignKey:RoleID;joinReferences:MenuID"`
	Actions   []rma_domain.Action `gorm:"many2many:role_menu_action;joinForeignKey:RoleID;joinReferences:ActionID"`
	CreatedAt time.Time           `gorm:"column:created_at"`
	CreatedBy string              `gorm:"column:created_by"`
	UpdatedAt time.Time           `gorm:"column:updated_at"`
	UpdatedBy string              `gorm:"column:updated_by"`
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

type CreateRoleBodyRequest struct {
	Name      string   `json:"name" validate:"required,max=50"`
	Echelon   string   `json:"echelon" validate:"required,oneof='II' 'III' 'IV' 'V'"`
	MenuIDS   []string `json:"menu_ids" validate:"required,gt=0"`
	ActionIDS []string `json:"action_ids" validate:"required,gt=0"`
}

type GetRolesListQueryRequest struct {
	Page  int `json:"page" validate:"required,min=1"`
	Limit int `json:"limit" validate:"required"`
}

type RoleDTO struct {
	ID       int                          `json:"id"`
	RoleName string                       `json:"roleName"`
	Echelon  string                       `json:"echelon"`
	Menus    []rma_domain.MenuOrActionDTO `json:"menus"`
	Actions  []rma_domain.MenuOrActionDTO `json:"actions"`
}
