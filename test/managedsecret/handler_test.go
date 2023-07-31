package managedsecret

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"smanager/internal/httputils"
	"smanager/internal/managedsecret"
	"smanager/internal/middleware"
	"smanager/test/fixture"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestManagedSecretHandlerCreateAndListSecret(t *testing.T) {
	router, _ := prepare()
	request := `
	{
		"name": "CreatedsecretName",
		"value": "CreatedsecretValue"
	}
	`
	sendRequest(router, "POST", bytes.NewBuffer([]byte(request)))

	responseData := sendRequest(router, "GET", bytes.NewBuffer([]byte(request)))

	var result httputils.HttpResponseDto[[]managedsecret.ManagedSecretDto]
	json.Unmarshal(responseData, &result)
	secretNames := make([]string, 0)
	for _, secret := range result.Data {
		secretNames = append(secretNames, secret.Name)
	}
	assert.Contains(t, secretNames, "CreatedsecretName")

}

func sendRequest(router *gin.Engine, method string, bodyBuffer *bytes.Buffer) []byte {
	req, _ := http.NewRequest(method, "/managedSecret", bodyBuffer)
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7IklkIjoxLCJVc2VybmFtZSI6InVzZXJuYW1lRm9yVGVzdHMifX0.hjH5ae2it81F-D4WRHpbCp4SjBf5hmOBOAsCUEIICaY")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	return responseData
}

func prepare() (*gin.Engine, fixture.DBFixture) {
	dbFixture := fixture.RunDBFixture()
	router := gin.Default()
	managedSecretService := managedsecret.NewManagedSercertService(dbFixture.ManagedSecretRepo)
	handlerPovider := managedsecret.NewHandlerConfigProvider(managedSecretService)
	middlewareMaps := map[middleware.MiddlewareType]gin.HandlerFunc{
		middleware.Secured: fixture.NewTestAuthMiddleware(),
	}
	httputils.RegisterRoutesWithMiddleware(router, handlerPovider.GetHandlers(), middlewareMaps)
	return router, dbFixture
}
