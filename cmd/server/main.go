package main

import (
	"github.com/Feinot/metric/cmd/server/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/update/counter/", handler.RequestHandle)
	http.HandleFunc("/update/gauge/", handler.RequestHandle)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
