package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Create(c *gin.Context) {
	var createUserRequest CreateUserDTO
	er1 := c.ShouldBindJSON(&createUserRequest)
	if er1 != nil {
		c.String(http.StatusInternalServerError, er1.Error())
		return
	}

	res, err := h.service.CreateUser(c.Request.Context(), createUserRequest)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
