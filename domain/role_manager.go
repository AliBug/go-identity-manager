package domain

// Policy - policy interface
type Policy interface {
	GetPolcy() []string
}

// RoleManagerUseCase - List and manage roles in domain
type RoleManagerUseCase interface {
	GetPolicies() [][]string
	GetPoliciesInDomain(string) [][]string
	AddNamedPolicy(Policy) (bool, error)
	RemoveNamedPolicy(Policy) (bool, error)
	RemoveFilteredNamedPolicy(domain string) (bool, error)
	DeleteUser(userID string) (bool, error)
}
