package domain

import (
	"context"

	"com.mailnau.api/common"
)

type Service interface {
	AddRole(ctx context.Context, name, echelon string, menuIDs, actionIDs []int64) (*common.BaseResponse[CreateRoleResponse], error)
	GetRolesWithMenuAndActions(ctx context.Context, limit, page int) (*common.BaseResponse[[]RoleDTO], error)
}

type Repository interface {
	CreateRole(ctx context.Context, name, echelon string, menuIDs, actionIDs []int64) (int64, error)
	FindRolesWithMenuAndActions(ctx context.Context, limit, offset int) ([]Role, error)
	CountRolesRecords(ctx context.Context) (int64, error)
}
