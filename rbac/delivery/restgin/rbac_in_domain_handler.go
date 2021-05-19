package restgin

import (
	"net/http"

	"github.com/alibug/go-identity-manager/domain"
	"github.com/alibug/go-identity-manager/rbac/body"
	"github.com/alibug/go-identity-utils/status"
	"github.com/gin-gonic/gin"
)

// RBACinDomainHandler -
type RBACinDomainHandler struct {
	rbacUC domain.RBACinDomainUseCase
}

// NewRBACinDomainHandler - new casbin handler
func NewRBACinDomainHandler(route *gin.Engine, rbacUC domain.RBACinDomainUseCase) {
	handler := &RBACinDomainHandler{rbacUC: rbacUC}
	group := route.Group("/domain")
	group.POST("/add-role-4u", handler.AddRoleForUserInDomain)
	group.POST("/del-role-4u", handler.DeleteRoleForUserInDomain)
	group.POST("/del-roles-4u", handler.DeleteRolesForUserInDomain)
	group.GET("/domains-4u/:user", handler.GetDomainsForUser)
	// group.POST("/roles-4u", handler.GetRolesForUserInDomain)
	group.GET("/roles-4u/:user", handler.GetRolesInDomainsForUser)
}

// AddRoleForUserInDomain - adds authorization rules to the current named policy
//					- 添加 策略
func (h *RBACinDomainHandler) AddRoleForUserInDomain(c *gin.Context) {
	var body body.RoleForUser

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, status.ResponseError{Message: err.Error()})
		return
	}

	if body.Domain == "" || body.Role == "" {
		c.JSON(http.StatusBadRequest, status.ResponseError{Message: "Domain or role required"})
		return
	}

	ok, err := h.rbacUC.AddRoleForUserInDomainUC(&body)
	if err != nil {
		c.JSON(status.GetStatusCode(err), status.ResponseError{Message: err.Error()})
		return
	}
	if ok {
		c.JSON(http.StatusAccepted, gin.H{"ok": true})
		return
	}
	c.JSON(http.StatusConflict, status.ResponseError{Message: "User already has the role"})
}

// DeleteRoleForUserInDomain -
func (h *RBACinDomainHandler) DeleteRoleForUserInDomain(c *gin.Context) {
	var body body.RoleForUser

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, status.ResponseError{Message: err.Error()})
		return
	}

	if body.Domain == "" || body.Role == "" {
		c.JSON(http.StatusBadRequest, status.ResponseError{Message: "Domain or role required"})
		return
	}

	ok, err := h.rbacUC.DeleteRoleForUserInDomainUC(&body)
	if err != nil {
		c.JSON(status.GetStatusCode(err), status.ResponseError{Message: err.Error()})
		return
	}

	if ok {
		c.JSON(http.StatusAccepted, gin.H{"ok": true})
		return
	}
	c.JSON(http.StatusConflict, status.ResponseError{Message: "User does not have the role"})
}

// DeleteRolesForUserInDomain -
func (h *RBACinDomainHandler) DeleteRolesForUserInDomain(c *gin.Context) {
	var body body.RoleForUser

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, status.ResponseError{Message: err.Error()})
		return
	}

	if body.Domain == "" {
		c.JSON(http.StatusBadRequest, status.ResponseError{Message: "Domain required"})
		return
	}

	ok, err := h.rbacUC.DeleteRolesForUserInDomainUC(&body)
	if err != nil {
		c.JSON(status.GetStatusCode(err), status.ResponseError{Message: err.Error()})
		return
	}

	if ok {
		c.JSON(http.StatusAccepted, gin.H{"ok": true})
		return
	}

	c.JSON(http.StatusConflict, status.ResponseError{Message: "User does not have any roles"})
}

// GetDomainsForUser -
func (h *RBACinDomainHandler) GetDomainsForUser(c *gin.Context) {
	name := c.Params.ByName("user")

	if name == "" {
		c.JSON(http.StatusBadRequest, status.ResponseError{Message: "Empty name"})
		return
	}

	roles, err := h.rbacUC.GetDomainsForUserUC(name)
	if err != nil {
		c.JSON(status.GetStatusCode(err), status.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roles})
}

// GetRolesForUserInDomain ---
func (h *RBACinDomainHandler) GetRolesForUserInDomain(c *gin.Context) {
	var body body.RoleForUser

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, status.ResponseError{Message: err.Error()})
		return
	}

	roles := h.rbacUC.GetRolesForUserInDomainUC(body.User, body.Domain)
	c.JSON(http.StatusOK, gin.H{"data": roles})
}

// GetRolesInDomainsForUser -
func (h *RBACinDomainHandler) GetRolesInDomainsForUser(c *gin.Context) {
	name := c.Params.ByName("user")

	if name == "" {
		c.JSON(http.StatusBadRequest, status.ResponseError{Message: "Empty name"})
		return
	}
	result, err := h.rbacUC.GetRolesInDomainsForUserUC(name)
	if err != nil {
		c.JSON(status.GetStatusCode(err), status.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}
