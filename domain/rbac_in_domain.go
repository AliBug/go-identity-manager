package domain

// UserRoleInDomain -
type UserRoleInDomain interface {
	GetDomain() string
	GetUser() string
	GetRole() string
}

// RoleInDomain -
type RoleInDomain interface {
	GetDomain() string
	GetRole() string
}

// RBACinDomainUseCase -
type RBACinDomainUseCase interface {
	AddRoleForUserInDomainUC(userRoleinDomain UserRoleInDomain) (bool, error)
	DeleteRoleForUserInDomainUC(userRoleinDomain UserRoleInDomain) (bool, error)
	DeleteRolesForUserInDomainUC(userRoleinDomain UserRoleInDomain) (bool, error)
	GetDomainsForUserUC(name string) ([]string, error)
	GetRolesForUserInDomainUC(name string, domain string) []string
	GetRolesInDomainsForUserUC(name string) (map[string][]string, error)
}
