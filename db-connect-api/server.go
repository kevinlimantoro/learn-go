package main

import (
	"fmt"
	"log"
	"member-db-api/controller"
	"member-db-api/loader"
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
func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("/api/v1")

	api.GET("/members", gin.WrapF(controller.GetAllMember))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("server_port"), r))
}
