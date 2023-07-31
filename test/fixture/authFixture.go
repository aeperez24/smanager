package fixture

import (
	"smanager/internal/middleware/auth"
	"smanager/internal/token"

	"github.com/gin-gonic/gin"
)

func NewTestAuthMiddleware() gin.HandlerFunc {
	return auth.NewAuthMiddleware(token.NewTokenService("key"))
}
