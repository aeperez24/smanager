package login

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	httputil "smanager/internal/httputils"
	"smanager/internal/login"
	"smanager/internal/token"
	"smanager/internal/user"
	"smanager/test/fixture"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	router, _ := prepare()
	request := fmt.Sprintf(`
	{
		"username": "%s",
		"password": "%s"
	}
	`, fixture.TEST_USERNAME, fixture.TEST_PASSWORD)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(request)))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	expectedResponse :=
		`{"data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7IklkIjoxLCJVc2VybmFtZSI6InVzZXJuYW1lRm9yVGVzdHMifX0.hjH5ae2it81F-D4WRHpbCp4SjBf5hmOBOAsCUEIICaY"},"errorMessage":""}`
	result := string(responseData)
	assert.Equal(t, expectedResponse, result)
}

func prepare() (*gin.Engine, fixture.DBFixture) {
	dbFixture := fixture.RunDBFixture()
	router := gin.Default()
	userService := user.NewUserService(dbFixture.UserRepo)
	tokenService := token.NewTokenService("key")
	ls := login.NewLoginService(userService, tokenService)
	lhc := login.NewLoginHandlerConfigProvider(ls)
	httputil.RegisterRoutes(router, lhc.GetHandlers())
	return router, dbFixture
}
