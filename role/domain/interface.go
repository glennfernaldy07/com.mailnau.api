package domain

import (
	"context"

	"com.mailnau.api/common"
)

type Service interface {
	AddRole(ctx context.Context, name, echelon string, menuIDs, actionIDs []int64) (*common.BaseResponse[CreateRoleResponse], error)
}

type Repository interface {
	CreateRole(ctx context.Context, name, echelon string, menuIDs, actionIDs []int64) (int64, error)
}
