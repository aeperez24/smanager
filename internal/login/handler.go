package login

import (
	"context"
	"net/http"
	"smanager/internal/common"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	LoginService ILoginService
}

func (ah *AuthHandler) Login(c *gin.Context) {
	request := make(map[string]string)
	er1 := c.ShouldBindJSON(request)
	responseDto := common.HttpResponseDto[LoginResponse]{}

	if er1 != nil {
		c.JSON(http.StatusBadRequest, responseDto)
		return
	}
	userName := request["username"]
	password := request["password"]

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
