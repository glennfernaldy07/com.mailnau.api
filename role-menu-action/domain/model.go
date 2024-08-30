package domain

type RoleMenuAction struct {
	RoleID   int64 `gorm:"column:role_id;foreignKey"`
	MenuID   int64 `gorm:"column:menu_id;foreignKey"`
	ActionID int64 `gorm:"column:action_id;foreignKey"`
}

type Menu struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"size:50;not null"`
}

type Action struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"size:50;not null"`
}

func (RoleMenuAction) TableName() string {
	return "role_menu_action"
}

type MenuOrActionDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
