package restgin

import (
	"net/http"

	"fmt"
	"strconv"

	"github.com/alibug/go-identity-manager/domain"
	"github.com/alibug/go-identity-utils/status"
	"github.com/gin-gonic/gin"
)

// UsersHandler -
type UsersHandler struct {
	usersUseCase        domain.UsersUsecase
	rolesManagerUseCase domain.RoleManagerUseCase
}

// NewRoleManagerHandler - new rolemanager handler
func NewRoleManagerHandler(route *gin.Engine, usersUC domain.UsersUsecase, rolesUC domain.RoleManagerUseCase) {
	handler := &UsersHandler{usersUseCase: usersUC, rolesManagerUseCase: rolesUC}
	group := route.Group("/users")
	group.GET("/limit/:limit/skip/:skip", handler.ListUsers)
	group.DELETE("/:userID", handler.DeleteUserByID)
}

// GetPathInt64 - convert string to int64
func GetPathInt64(c *gin.Context, name string) (int64, error) {
	val := c.Params.ByName(name)
	if val == "" {
		return 0, fmt.Errorf("%w: %s param required", status.ErrBadParamInput, name)
	}
	return strconv.ParseInt(val, 10, 64)
}

// ListUsers -
func (u *UsersHandler) ListUsers(c *gin.Context) {
	limit, err := GetPathInt64(c, "limit")
	if err != nil {
		c.JSON(http.StatusBadRequest, status.ResponseError{Message: err.Error()})
		return
	}
	skip, err := GetPathInt64(c, "skip")
	if err != nil {
		c.JSON(http.StatusBadRequest, status.ResponseError{Message: err.Error()})
		return
	}

	ctx := c.Request.Context()
	users, err := u.usersUseCase.ListUsersUC(ctx, limit, skip)
	if err != nil {
		c.JSON(status.GetStatusCode(err), status.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// DeleteUserByID - 删除用户 以及 用户相关的角色
func (u *UsersHandler) DeleteUserByID(c *gin.Context) {
	userID := c.Params.ByName("userID")
	ctx := c.Request.Context()
	err := u.usersUseCase.DeleteUserUC(ctx, userID)
	if err != nil {
		c.JSON(status.GetStatusCode(err), status.ResponseError{Message: err.Error()})
		return
	}
	_, err = u.rolesManagerUseCase.DeleteUser(userID)
	if err != nil {
		c.JSON(status.GetStatusCode(err), status.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
