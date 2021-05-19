package body

// RoleForUser -
type RoleForUser struct {
	User   string `json:"user,omitempty" binding:"required"`
	Role   string `json:"role,omitempty"`
	Domain string `json:"domain,omitempty"`
}

// GetUser -
func (r *RoleForUser) GetUser() string {
	return r.User
}

// GetRole -
func (r *RoleForUser) GetRole() string {
	return r.Role
}

// GetDomain -
func (r *RoleForUser) GetDomain() string {
	return r.Domain
}
