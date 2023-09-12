package main

import (
	"flag"
	"github.com/Feinot/metric/cmd/server/handler"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	host string
)

func init() {

}

func main() {

	Server()
}
func GetConfig() []string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	if os.Getenv("ADDRESS") != "" {
		return strings.Split(os.Getenv("ADDRESS"), "localhost")
	}
	flag.StringVar(&host, "a", "localhost:8080", "")

	flag.Parse()
	return strings.Split(host, "localhost")
}
func Server() {

	q := GetConfig()
	r := chi.NewRouter()

	r.Post("/update/{type}/{name}/{value}", handler.RequestUpdateHandle)

	r.Get("/value/{type}/{name}", handler.RequestValueHandle)

	r.Get("/", handler.HomeHandle)

	if err := http.ListenAndServe(q[1], r); err != nil {
		log.Fatal(err)
	}
}
