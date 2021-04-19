package usecase

import (
	"context"
	"time"

	"github.com/alibug/go-user-casbin/domain"
)

type rbacInDomainUsecase struct {
	rbacRepo       domain.RBACinDomainRepository
	contextTimeout time.Duration
}

// NewRBACinDomainUsecase -
func NewRBACinDomainUsecase(rbacRepo domain.RBACinDomainRepository, timeout time.Duration) domain.RBACinDomainUseCase {
	return &rbacInDomainUsecase{rbacRepo: rbacRepo, contextTimeout: timeout}
}

func (r *rbacInDomainUsecase) AddRoleForUserInDomainUc(c context.Context, ur domain.UserRoleInDomain) (bool, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.rbacRepo.AddRoleForUserInDomain(ctx, ur)
}

func (r *rbacInDomainUsecase) DeleteRoleForUserInDomainUc(c context.Context, ur domain.UserRoleInDomain) (bool, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.rbacRepo.DeleteRoleForUserInDomain(ctx, ur)
}

func (r *rbacInDomainUsecase) DeleteRolesForUserInDomainUc(c context.Context, ur domain.UserRoleInDomain) (bool, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.rbacRepo.DeleteRolesForUserInDomain(ctx, ur)
}
