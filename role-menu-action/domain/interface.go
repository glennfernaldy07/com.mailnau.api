package domain

import (
	"context"

	"com.mailnau.api/common"
)

type Service interface {
	GetAllMenus(ctx context.Context) (*common.BaseResponse[[]MenuOrActionDTO], error)
	GetAllActions(ctx context.Context) (*common.BaseResponse[[]MenuOrActionDTO], error)
}

type Repository interface {
	FindAllMenus(ctx context.Context) ([]Menu, error)
	FindAllActions(ctx context.Context) ([]Action, error)
}
