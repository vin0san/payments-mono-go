package http

import (
	"strings"

	"pye/pkg/response"
	"pye/pkg/security"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			response.Error(c, 401, "unauthorized", "missing token", "")
			c.Abort()
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		userID, err := security.VerifyToken(token)
		if err != nil {
			response.Error(c, 401, "unauthorized", "invalid token", "")
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
