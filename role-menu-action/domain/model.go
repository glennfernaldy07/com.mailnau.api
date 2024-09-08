package domain

type RoleMenuAction struct {
	RoleID   int    `gorm:"column:role_id;foreignKey"`
	MenuID   string `gorm:"column:menu_id;foreignKey"`
	ActionID string `gorm:"column:action_id;foreignKey"`
}

type Menu struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"column:menu_name;size:255;not null"`
}

func (Menu) TableName() string {
	return "menu"
}

type Action struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"column:action_name;size:255;not null"`
}

func (Action) TableName() string {
	return "action"
}

func (RoleMenuAction) TableName() string {
	return "role_menu_action"
}

type MenuOrActionDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
