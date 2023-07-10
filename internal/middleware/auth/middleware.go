package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("running authMiddleware")
		c.Next()
	}
}
