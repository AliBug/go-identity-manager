package domain

import (
	"context"
	"time"
)

// User ...
type User interface {
	GetUserID() string
	GetDisplayName() string
	SetUpdatedTime(*time.Time)
}

// UserUsecase ...
type UserUsecase interface {
	GetByIDUc(ctx context.Context, id string) (User, error)
}

// UserRepository ...
type UserRepository interface {
	GetByID(ctx context.Context, id string) (User, error)
}
