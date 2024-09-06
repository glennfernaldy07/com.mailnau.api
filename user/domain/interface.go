package domain

import (
	"com.mailnau.api/common"
	"com.mailnau.api/common/snap/snapauth"
	"context"
)

type Service interface {
	Register(ctx context.Context, req RegisterRequest) (common.GeneralResponse, error)
	Login(ctx context.Context, req LoginRequest) (common.GeneralResponse, error)
	GetUserByUsernameAndPassword(ctx context.Context, username, password string) (*common.GeneralResponse, error)
}

type Repository interface {
	StoreUser(ctx context.Context, model User) (User, error)
	FindUserByEmail(ctx context.Context, email string) (User, error)

	StoreUserRole(ctx context.Context, role UserRole) error
}

type CacheRepository interface {
	StoreAccessToken(ctx context.Context, userID string, dt snapauth.AccessTokenResponse) error
	GetAccessToken(ctx context.Context, accessToken string) (bool, error)
}
