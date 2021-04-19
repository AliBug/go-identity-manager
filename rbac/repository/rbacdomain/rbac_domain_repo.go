package rbacdomain

import (
	"context"
	"log"

	"github.com/alibug/go-user-casbin/domain"
	"github.com/casbin/casbin/v2"
)

type casbinRepo struct {
	enforcer *casbin.Enforcer
}

// NewCasbinRepository -
func NewCasbinRepository(enforcer *casbin.Enforcer) domain.RBACinDomainRepository {
	return &casbinRepo{enforcer}
}

func (c *casbinRepo) AddRoleForUserInDomain(ctx context.Context, ur domain.UserRoleInDomain) (bool, error) {
	// 按道理应确定 role 存在
	roleExist := c.enforcer.HasNamedPolicy("p", ur.GetRole())
	if !roleExist {
		log.Printf("role %s not exist", ur.GetRole())

	}
	return c.enforcer.AddRoleForUserInDomain(ur.GetUserID(), ur.GetRole(), ur.GetDomain())
}

func (c *casbinRepo) DeleteRoleForUserInDomain(ctx context.Context, ur domain.UserRoleInDomain) (bool, error) {
	return c.enforcer.DeleteRoleForUserInDomain(ur.GetUserID(), ur.GetRole(), ur.GetDomain())
}

func (c *casbinRepo) DeleteRolesForUserInDomain(ctx context.Context, ur domain.UserRoleInDomain) (bool, error) {
	return c.enforcer.DeleteRolesForUserInDomain(ur.GetUserID(), ur.GetDomain())
}
