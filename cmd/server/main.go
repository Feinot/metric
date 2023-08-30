package main

import (
	"github.com/Feinot/metric/cmd/server/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/update/{type}/{name}/{value}", handler.RequestUpdateHandle)
	r.HandleFunc("/value/", handler.RequestValueHandle)

	r.HandleFunc("/", handler.HomeHandle)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
