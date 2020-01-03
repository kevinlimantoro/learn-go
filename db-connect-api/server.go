package main

import (
	"fmt"
	"log"
	"member-db-api/controller"
	"member-db-api/loader"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func init() {
	if e := godotenv.Load(); e != nil {
		fmt.Println(e)
	}
	loader.InitAllContext()
	defer timeTrack(time.Now(), "server load")
}
func main() {
	r := mux.NewRouter().StrictSlash(true)
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	})
	api.HandleFunc("/members", controller.GetAllMember).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("server_port"), r))
}
