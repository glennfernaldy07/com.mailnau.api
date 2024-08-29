package domain

type RoleMenuAction struct {
	RoleID   int64 `gorm:"column:role_id;foreignKey"`
	MenuID   int64 `gorm:"column:menu_id;foreignKey"`
	ActionID int64 `gorm:"column:action_id;foreignKey"`
}

func (RoleMenuAction) TableName() string {
	return "role_menu_action"
}
