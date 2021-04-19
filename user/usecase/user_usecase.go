package usecase

import (
	"context"
	"time"

	"github.com/alibug/go-user-casbin/domain"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func (u *userUsecase) GetByIDUc(c context.Context, id string) (res domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.userRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}
