package main

import (
	"flag"
	"fmt"
	"github.com/Feinot/metric/cmd/server/handler"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

var (
	host *string
)

func init() {

}

func main() {

	Server()
}
func Server() {
	host = flag.String("a", ":8080", "")
	flag.Parse()

	r := chi.NewRouter()
	fmt.Println(*host)
	r.Post("/update/{type}/{name}/{value}", handler.RequestUpdateHandle)

	r.Get("/value/{type}/{name}", handler.RequestValueHandle)

	r.Get("/", handler.HomeHandle)

	if err := http.ListenAndServe(*host, r); err != nil {
		log.Fatal(err)
	}
}
