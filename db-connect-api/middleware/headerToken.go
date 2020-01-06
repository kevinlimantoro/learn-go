package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

func HeaderAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("x-authentication")
		if auth != "" {
			// DO auth logic here
			ctx := context.WithValue(c.Request.Context(), "authHeader", auth)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		} else {
			respondWithError(c, 401, "Unauthorized")
		}
	}
}
