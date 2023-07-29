package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

const endpoint string = "/test"
const tokenHeaderName = "Authorization"

func TestTokenInformationOnContext(t *testing.T) {
	handler := newTestHandler(t)
	router := gin.Default()
	router.POST(endpoint, handler)
	req, _ := http.NewRequest("POST", endpoint, nil)
	token := "bearer tokeeen"
	req.Header.Add(tokenHeaderName, token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

}

func newTestHandler(t *testing.T) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get(tokenHeaderName)
		fmt.Printf("\n %s\n", token)
		ctx.JSON(http.StatusOK, "")
	}
}
