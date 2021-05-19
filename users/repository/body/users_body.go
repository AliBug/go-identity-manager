package body

import (
	"time"

	"github.com/alibug/go-identity-utils/converter"
)

// UserBody - Contain Register Body
type UserBody struct {
	ID converter.StrToObjectID `bson:"_id,omitempty" json:"id,omitempty"` // 用户ID
	// RegisterBody
	Account     string     `json:"account" bson:"account" binding:"required"`
	Displayname string     `json:"displayname"  bson:"displayname" binding:"required"`
	CreatedAt   *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   *time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"` // 更新时间
}

// GetUserID - implement domain.User
func (u *UserBody) GetUserID() string {
	return string(u.ID)
}

// GetDisplayName - implement domain.User
func (u *UserBody) GetDisplayName() string {
	return string(u.Displayname)
}
