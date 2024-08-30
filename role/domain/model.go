package domain

import (
	"time"

	rma_domain "com.mailnau.api/role-menu-action/domain"
)

type Role struct {
	ID        int64               `gorm:"column:id;primaryKey"`
	Name      string              `gorm:"column:name"`
	Echelon   string              `gorm:"column:echelon"`
	Menus     []rma_domain.Menu   `gorm:"many2many:role_menu_action;joinForeignKey:RoleID;joinReferences:MenuID"`
	Actions   []rma_domain.Action `gorm:"many2many:role_menu_action;joinForeignKey:RoleID;joinReferences:ActionID"`
	CreatedAt time.Time           `gorm:"column:created_at"`
	UpdatedAt time.Time           `gorm:"column:updated_at"`
}

func (Role) TableName() string {
	return "roles"
}

type CreateRoleResponse struct {
	RoleId int64 `json:"roleId"`
}

type RoleDTO struct {
	ID       int64                        `json:"id"`
	RoleName string                       `json:"roleName"`
	Echelon  string                       `json:"echelon"`
	Menus    []rma_domain.MenuOrActionDTO `json:"menus"`
	Actions  []rma_domain.MenuOrActionDTO `json:"actions"`
}
