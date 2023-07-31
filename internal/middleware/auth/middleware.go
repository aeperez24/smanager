package auth

import (
	"context"
	"net/http"
	"smanager/internal/user"

	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware(ts TokenClaimService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		claims, err := ts.GetClaims(token)
		if err != nil || claims["user"] == nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		userInContext := claims["user"].(map[string]interface{})
		userName, okUsername := userInContext["Username"].(string)

		idAsFloat, okId := userInContext["Id"].(float64)
		if !okId || !okUsername {
			c.Abort()
			c.JSON(http.StatusUnauthorized, "")

			return
		}

		userID := int(idAsFloat)
		newContext := context.WithValue(c.Request.Context(), "user", user.UserDTO{Id: userID,
			Username: userName})
		c.Request = c.Request.WithContext(newContext)
		c.Next()
	}
}

type TokenClaimService interface {
	GetClaims(token string) (map[string]interface{}, error)
}
