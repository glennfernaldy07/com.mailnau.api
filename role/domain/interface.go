package domain

import (
	"context"

	"com.mailnau.api/common"
)

type Service interface {
	AddRole(ctx context.Context, name, echelon string, menuIDs, actionIDs []string) (common.GeneralResponse, error)
	GetRolesWithMenuAndActions(ctx context.Context, limit, page int) (common.GeneralResponse, error)
	GetRoleByID(ctx context.Context, id int) (Role, error)

	GetListMenuByRoleID(ctx context.Context, roleID int) ([]string, error)
}

type Repository interface {
	FindRoleByID(ctx context.Context, id int) (Role, error)
	FindRolesWithMenuAndActions(ctx context.Context, limit, offset int) ([]Role, error)
	CreateRole(ctx context.Context, name, echelon string, menuIDs, actionIDs []string) (int, error)
	CountRolesRecords(ctx context.Context) (int64, error)

	FindRoleMenuByRoleID(ctx context.Context, roleID int) ([]RoleMenu, error)
}

type CacheRepository interface {
	StoreRoleByID(ctx context.Context, id int, role Role) error
	GetRoleByID(ctx context.Context, id int) (Role, error)

	StoreListMenuByRoleID(ctx context.Context, id int, listMenu []string) error
	GetListMenuByRoleID(ctx context.Context, id int) ([]string, error)
}
