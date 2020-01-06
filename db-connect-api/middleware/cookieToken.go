package middleware

import (
	"context"
	"fmt"
	"member-db-api/util"

	"github.com/gin-gonic/gin"
)

func CookieAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("authCookie")
		fmt.Println(cookie, err)
		if err == nil {
			// DO auth logic here
			ctx := context.WithValue(c.Request.Context(), "authCookie", cookie.Value)
			c.Request = c.Request.WithContext(ctx)
		} else {
			fmt.Println(err)
		}
		ctx := context.WithValue(c.Request.Context(), "authCookie", util.BaseSuccessResponse)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
