package domain

import (
	"com.mailnau.api/common"
	"context"
)

type Service interface {
	Register(ctx context.Context, req RegisterRequest) (common.GeneralResponse, error)
	LoginByEmail(ctx context.Context, req LoginByEmail) (common.GeneralResponse, error)
	LoginByNIK(ctx context.Context, req LoginByNIK) (common.GeneralResponse, error)
}

type Repository interface {
	StoreUser(ctx context.Context, model User) (User, error)
	FindUserByEmail(ctx context.Context, email string) (User, error)
	FindUserByNIK(ctx context.Context, nik string) (User, error)

	StoreUserRole(ctx context.Context, role UserRole) error
}

type CacheRepository interface {
	StoreAccessToken(ctx context.Context, userID string, token string) error
	GetAccessToken(ctx context.Context, accessToken string) (bool, error)
}
