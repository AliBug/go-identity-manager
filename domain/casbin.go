package domain

// UserRoleInDomain - Casbin RBAC model for domains
type UserRoleInDomain interface {
	GetDomain() string
	GetUserID() string
	GetRole() string
}
