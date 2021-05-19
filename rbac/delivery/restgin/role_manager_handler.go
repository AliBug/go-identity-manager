package restgin

import (
	"log"
	"net/http"

	"github.com/alibug/go-identity-manager/domain"
	"github.com/alibug/go-identity-manager/rbac/body"
	"github.com/alibug/go-identity-utils/status"
	"github.com/gin-gonic/gin"
)

// RoleManagerHandler -
type RoleManagerHandler struct {
	roleManager domain.RoleManagerUseCase
}

// NewRoleManagerHandler - new rolemanager handler
func NewRoleManagerHandler(route *gin.Engine, roleManager domain.RoleManagerUseCase) {
	handler := &RoleManagerHandler{roleManager: roleManager}
	group := route.Group("/roles")
	group.GET("/policies", handler.GetPolicies)
	group.GET("/:domain", handler.GetPoliciesInDomain)
	group.POST("/add-policy", handler.AddNamedPolicy)
	group.POST("/remove-policy", handler.RemoveNamedPolicy)
	group.DELETE("/:domain", handler.RemoveFilteredNamedPolicy)
}

// GetPoliciesInDomain -
func (h *RoleManagerHandler) GetPoliciesInDomain(c *gin.Context) {
	domainName := c.Param("domain")
	roles := h.roleManager.GetPoliciesInDomain(domainName)
	c.JSON(http.StatusOK, gin.H{"data": roles})
}

// AddNamedPolicy -
func (h *RoleManagerHandler) AddNamedPolicy(c *gin.Context) {
	var reqBody body.Policy

	if err := c.ShouldBind(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, status.ResponseError{Message: err.Error()})
		return
	}

	if len(reqBody.Params) == 0 {
		c.JSON(http.StatusBadRequest, status.ResponseError{Message: "Policy params required"})
		return
	}

	for _, v := range reqBody.Params {
		if v == "" {
			c.JSON(http.StatusBadRequest, status.ResponseError{Message: "Policy param required"})
			return
		}
	}

	ok, err := h.roleManager.AddNamedPolicy(&reqBody)
	if err != nil {
		c.JSON(status.GetStatusCode(err), status.ResponseError{Message: err.Error()})
		return
	}
	if ok {
		c.JSON(http.StatusAccepted, gin.H{"ok": true})
		return
	}
	c.JSON(http.StatusConflict, status.ResponseError{Message: "Policy already exists"})
}

// RemoveNamedPolicy -
func (h *RoleManagerHandler) RemoveNamedPolicy(c *gin.Context) {
	var reqBody body.Policy
	if err := c.ShouldBind(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, status.ResponseError{Message: err.Error()})
		return
	}
	ok, err := h.roleManager.RemoveNamedPolicy(&reqBody)
	if err != nil {
		c.JSON(status.GetStatusCode(err), status.ResponseError{Message: err.Error()})
		return
	}
	if ok {
		c.JSON(http.StatusAccepted, gin.H{"ok": true})
		return
	}
	c.JSON(http.StatusNotFound, status.ResponseError{Message: "Policy not found"})
}

// RemoveFilteredNamedPolicy -
func (h *RoleManagerHandler) RemoveFilteredNamedPolicy(c *gin.Context) {
	domainName := c.Param("domain")
	log.Println("del :", domainName)
	ok, err := h.roleManager.RemoveFilteredNamedPolicy(domainName)
	if err != nil {
		c.JSON(status.GetStatusCode(err), status.ResponseError{Message: err.Error()})
		return
	}
	if ok {
		c.JSON(http.StatusAccepted, gin.H{"ok": true})
		return
	}
	c.JSON(http.StatusNotFound, status.ResponseError{Message: "Not match policy"})
}

// GetPolicies - 获取 所有策略
func (h *RoleManagerHandler) GetPolicies(c *gin.Context) {
	roles := h.roleManager.GetPolicies()
	c.JSON(http.StatusOK, gin.H{"data": roles})
}
