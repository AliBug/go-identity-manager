package restgin

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// CasbinHandler -
type CasbinHandler struct {
	enforcer *casbin.Enforcer
}

// NewCasbinHandler - new casbin handler
func NewCasbinHandler(route *gin.Engine, e *casbin.Enforcer) {
	handler := &CasbinHandler{enforcer: e}
	route.POST("/policies", handler.AddNamedPolicies)
}

// NamedPolicies - struct for Policies
type NamedPolicies struct {
	Policies [][]string
}

// AddNamedPolicies - adds authorization rules to the current named policy
//					- 添加 策略
func (h *CasbinHandler) AddNamedPolicies(c *gin.Context) {
	var policies NamedPolicies

	if err := c.ShouldBind(&policies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	ok, err := h.enforcer.AddNamedPolicies("p", policies.Policies)
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
