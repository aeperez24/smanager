package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"smanager/internal/middleware/auth"
	"smanager/internal/token"
	"smanager/internal/user"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const endpoint string = "/test"
const tokenHeaderName = "Authorization"
const username = "username"
const id = 1

func TestTokenInformationOnContext(t *testing.T) {
	handler := newTestHandler(t)
	router := gin.Default()
	tokenService := token.NewTokenService("the secret")
	authMiddleware := auth.NewAuthMiddleware(tokenService)
	router.POST(endpoint, authMiddleware, handler)
	req, _ := http.NewRequest("POST", endpoint, nil)
	claims := map[string]interface{}{
		"user": user.UserDTO{Id: id, Username: username},
	}
	token, _ := tokenService.CreateToken(claims)
	req.Header.Add(tokenHeaderName, token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

}

func newTestHandler(t *testing.T) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		fmt.Println(ctx.Request.Context().Value("user"))
		user := ctx.Request.Context().Value("user").(user.UserDTO)
		ctx.JSON(http.StatusOK, "")
		assert.Equal(t, id, user.Id)
		assert.Equal(t, username, user.Username)

	}
}
