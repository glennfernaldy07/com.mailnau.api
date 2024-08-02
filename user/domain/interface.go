package domain

import "context"

type Service interface {
	//Login(ctx context.Context)
	GetUserByUsernameAndPassword(ctx context.Context, username, password string) (*UserResponse, error)
}

type Repository interface {
	FindUserByUsernameAndPassword(ctx context.Context, username, password string) (*User, error)
}
