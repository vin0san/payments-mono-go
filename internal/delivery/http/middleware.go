package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Prefer incoming request-id header if exists (for distributed tracing)
		reqID := c.GetHeader("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}

		c.Set("trace_id", reqID)
		c.Writer.Header().Set("X-Request-ID", reqID)

		c.Next()
	}
}
