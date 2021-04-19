package domain

import "context"

// RBACinDomainUseCase -
type RBACinDomainUseCase interface {
	AddRoleForUserInDomainUc(ctx context.Context, userRoleinDomain UserRoleInDomain) (bool, error)
	DeleteRoleForUserInDomainUc(ctx context.Context, userRoleinDomain UserRoleInDomain) (bool, error)
	DeleteRolesForUserInDomainUc(ctx context.Context, userRoleinDomain UserRoleInDomain) (bool, error)
}

// RBACinDomainRepository -
type RBACinDomainRepository interface {
	AddRoleForUserInDomain(ctx context.Context, userRoleinDomain UserRoleInDomain) (bool, error)
	DeleteRoleForUserInDomain(ctx context.Context, userRoleinDomain UserRoleInDomain) (bool, error)
	DeleteRolesForUserInDomain(ctx context.Context, userRoleinDomain UserRoleInDomain) (bool, error)
}
