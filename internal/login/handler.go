package login

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	LoginService ILoginService
}

func (ah *AuthHandler) Login(c *gin.Context) {
	request := make(map[string]string)
	er1 := c.ShouldBindJSON(request)
	if er1 != nil {
		c.String(http.StatusInternalServerError, er1.Error())
		return
	}
	userName := request["username"]
	password := request["password"]

	token, err := ah.LoginService.LoginWithUsernameAndPassword(c.Request.Context(), userName, password)
	if err != nil {
		err, ok := err.(InvalidUserError)
		if ok {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}
		c.String(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ILoginService interface {
	LoginWithUsernameAndPassword(ctx context.Context, username, password string) (string, error)
}
