package token

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type tokenService struct {
	secretKey     string
	signingMethod jwt.SigningMethod
	keyFunc       func(token *jwt.Token) (interface{}, error)
}

func NewTokenService(secretKey string) *tokenService {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid Token")
		}
		return []byte(secretKey), nil
	}
	return &tokenService{secretKey, jwt.SigningMethodHS256, keyFunc}
}

func (service *tokenService) CreateToken(claims map[string]interface{}) (string, error) {
	claimsMap := jwt.MapClaims{}
	for k, v := range claims {
		claimsMap[k] = v
	}

	token := jwt.NewWithClaims(service.signingMethod, claimsMap)

	return token.SignedString([]byte(service.secretKey))
}

func (service *tokenService) GetClaims(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, service.keyFunc)
	if err != nil {
		return nil, fmt.Errorf("GetClaims:%w", err)
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}
	claims := parsedToken.Claims.(jwt.MapClaims)

	return claims, nil
}
