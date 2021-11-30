package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancal := context.WithTimeout(c.Request.Context(), t)
		defer cancal()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
