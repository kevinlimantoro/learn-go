package main

import (
	"context"
	"fmt"
	"log"
	"member-db-api/controller"
	"member-db-api/loader"
	"member-db-api/util"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func init() {
	gin.SetMode(os.Getenv(gin.EnvGinMode))
	defer timeTrack(time.Now(), "server load")
	if e := godotenv.Load(); e != nil {
		fmt.Println(e)
	}
	loader.InitAllContext()
}
func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func tokenAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("ckiftk")
		fmt.Println(cookie, err)
		if err == nil {
			// DO auth logic here
			ctx := context.WithValue(c.Request.Context(), "ckiftk", cookie.Value)
			c.Request = c.Request.WithContext(ctx)
		}
		ctx := context.WithValue(c.Request.Context(), "ckiftk", util.BaseSuccessResponse)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
func main() {
	r := gin.New()
	r.Use(tokenAuth())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("/api/v1")

	api.GET("/members", gin.WrapF(controller.GetAllMember))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("server_port"), r))
}
