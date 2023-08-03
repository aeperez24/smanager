package user

import (
	"context"
	"net/http"
	"smanager/internal/httputils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service IUserService
}

func NewUserHandler(service IUserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Create(c *gin.Context) {
	var createUserRequest CreateUserDTO
	er1 := c.ShouldBindJSON(&createUserRequest)
	var response httputils.HttpResponseDto[*UserDTO]
	if er1 != nil {
		response.ErrorMessage = er1.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	res, err := h.service.CreateUser(c.Request.Context(), createUserRequest)
	if err != nil {
		response.ErrorMessage = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return

	}
	response.Data = res
	c.JSON(http.StatusOK, response)
}

type IUserService interface {
	CreateUser(ctx context.Context, req CreateUserDTO) (*UserDTO, error)
}
