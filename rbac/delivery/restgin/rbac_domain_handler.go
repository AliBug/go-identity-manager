package restgin

import (
	"net/http"

	"github.com/alibug/go-user-casbin/domain"
	"github.com/alibug/go-user-casbin/rbac/repository/body"
	"github.com/gin-gonic/gin"
)

// RBACinDomainHandler -
type RBACinDomainHandler struct {
	rbacUC domain.RBACinDomainUseCase
}

/*
type RBACWithDomainUseCase interface {
	AddRoleForUserInDomain(userRoleinDomain UserRoleinDomain) (bool, error)
	DeleteRoleForUserInDomain(userRoleinDomain UserRoleinDomain) (bool, error)
	DeleteRolesForUserInDomain(userRoleinDomain UserRoleinDomain) (bool, error)
}

// RBACWithDomainRepository -
type RBACWithDomainRepository interface {
	AddRoleForUserInDomain(user string, role string, domain string) (bool, error)
	DeleteRoleForUserInDomain(user string, role string, domain string) (bool, error)
	DeleteRolesForUserInDomain(user string, domain string) (bool, error)
}
*/

// NewRBACinDomainHandler - new casbin handler
func NewRBACinDomainHandler(route *gin.Engine, rbacUC domain.RBACinDomainUseCase) {
	handler := &RBACinDomainHandler{rbacUC: rbacUC}
	group := route.Group("/domain")
	group.POST("/addroleforuser", handler.AddRoleForUserInDomain)
	// route.POST("/delroleforuser", handler.)
}

// AddRoleForUserInDomain - adds authorization rules to the current named policy
//					- 添加 策略
func (h *RBACinDomainHandler) AddRoleForUserInDomain(c *gin.Context) {
	var body body.RoleForUser

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	ctx := c.Request.Context()
	ok, err := h.rbacUC.AddRoleForUserInDomainUc(ctx, &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	if ok {
		c.JSON(http.StatusAccepted, gin.H{"ok": true})
		return
	}
	c.JSON(http.StatusConflict, gin.H{"message": "rules already exist"})
}
