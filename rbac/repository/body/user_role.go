package body

// RoleForUser -
type RoleForUser struct {
	UserID string `json:"id,omitempty" binding:"required"`
	Role   string `json:"role,omitempty"`
	Domain string `json:"domain,omitempty"`
}

// GetUserID -
func (r *RoleForUser) GetUserID() string {
	return r.UserID
}

// GetRole -
func (r *RoleForUser) GetRole() string {
	return r.Role
}

// GetDomain -
func (r *RoleForUser) GetDomain() string {
	return r.Domain
}
