package domain

import (
	"context"
)

type Service interface {
	GetRoleByID(ctx context.Context, id int) (Role, error)

	GetListMenuByRoleID(ctx context.Context, roleID int) ([]string, error)
}

type Repository interface {
	FindRoleByID(ctx context.Context, id int) (Role, error)

	FindRoleMenuByRoleID(ctx context.Context, roleID int) ([]RoleMenu, error)
}

type CacheRepository interface {
	StoreRoleByID(ctx context.Context, id int, role Role) error
	GetRoleByID(ctx context.Context, id int) (Role, error)

	StoreListMenuByRoleID(ctx context.Context, id int, listMenu []string) error
	GetListMenuByRoleID(ctx context.Context, id int) ([]string, error)
}
