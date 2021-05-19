package usecase

import (
	"context"
	"time"

	"github.com/alibug/go-identity-manager/domain"
)

type usersUsecase struct {
	usersRepo      domain.UsersRepository
	contextTimeout time.Duration
}

// NewUserUsecase will create new an userUsecase object representation of domain.ArticleUsecase interface
func NewUsersUsecase(repo domain.UsersRepository, timeout time.Duration) domain.UsersUsecase {
	return &usersUsecase{
		usersRepo:      repo,
		contextTimeout: timeout,
	}
}

func (u *usersUsecase) ListUsersUC(c context.Context, limit int64, skip int64) (users []domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	users, err = u.usersRepo.ListUsers(ctx, limit, skip)
	if err != nil {
		return
	}

	return
}

// DeleteUserUC -
func (u *usersUsecase) DeleteUserUC(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.usersRepo.DeleteUserByID(ctx, userID)
}
