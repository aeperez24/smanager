package login

import (
	"context"
	"fmt"
	"net/http"
	"smanager/internal/common"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginService ILoginService
}

func (ah *LoginHandler) Login(c *gin.Context) {
	var request LoginRequest
	er1 := c.ShouldBindJSON(&request)
	responseDto := common.HttpResponseDto[LoginResponse]{}

	if er1 != nil {
		fmt.Print(er1.Error())
		c.JSON(http.StatusBadRequest, responseDto)
		return
	}
	userName := request.Username
	password := request.Password

	token, err := ah.LoginService.LoginWithUsernameAndPassword(c.Request.Context(), userName, password)
	if err != nil {
		err, ok := err.(InvalidUserError)
		if ok {
			responseDto.ErrorMessage = err.Error()
			c.JSON(http.StatusUnauthorized, responseDto)
			return
		}
		responseDto.ErrorMessage = "internal error"
		c.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	responseDto.Data = LoginResponse{Token: token}
	c.JSON(http.StatusOK, responseDto)
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ILoginService interface {
	LoginWithUsernameAndPassword(ctx context.Context, username, password string) (string, error)
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
