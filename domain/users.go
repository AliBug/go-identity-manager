package domain

import (
	"context"
)

// User ...
type User interface {
	GetUserID() string
	GetDisplayName() string
}

// UsersUsecase ...
type UsersUsecase interface {
	ListUsersUC(ctx context.Context, limit int64, skip int64) ([]User, error)
	DeleteUserUC(c context.Context, userID string) error
}

// UsersRepository ...
type UsersRepository interface {
	ListUsers(ctx context.Context, limit int64, skip int64) ([]User, error)
	DeleteUserByID(ctx context.Context, userID string) error
}
