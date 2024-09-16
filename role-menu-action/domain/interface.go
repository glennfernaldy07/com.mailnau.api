package domain

import (
	"context"

	"com.mailnau.api/common"
)

type Service interface {
	GetAllMenus(ctx context.Context) (common.GeneralResponse, error)
	GetAllActions(ctx context.Context) (common.GeneralResponse, error)

	GetRoleMenuActionByRoleID(ctx context.Context, roleID int) ([]RoleMenuAction, error)
}

type Repository interface {
	FindAllMenus(ctx context.Context) ([]Menu, error)
	FindAllActions(ctx context.Context) ([]Action, error)

	FindRoleMenuActionByRoleID(ctx context.Context, roleID int) ([]RoleMenuAction, error)
}

type CacheRepository interface {
	StoreListMenu(ctx context.Context, listMenu []Menu) error
	StoreListAction(ctx context.Context, listAction []Action) error
	StoreRoleMenuAction(ctx context.Context, roleMenuAction []RoleMenuAction, roleID int) error

	GetListMenu(ctx context.Context) ([]Menu, error)
	GetListAction(ctx context.Context) ([]Action, error)
	GetRoleMenuActionByRoleID(ctx context.Context, roleID int) ([]RoleMenuAction, error)
}
