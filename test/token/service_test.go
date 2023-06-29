package token

import (
	"fmt"
	"smanager/internal/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndValidateToken(t *testing.T) {
	service := token.NewTokenService("thesecret")
	claimsMap := make(map[string]interface{})
	claimsMap["theKey"] = "theValue"
	token, err := service.CreateToken(claimsMap)
	assert.Nil(t, err)
	claims, err := service.GetClaims(token)
	assert.Nil(t, err)
	fmt.Println(token)
	assert.Equal(t, claims["theKey"], "theValue")

}
