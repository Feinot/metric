package main

import (
	"fmt"
	"github.com/Feinot/metric/cmd/server/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/update/counter", handler.RequestHandle)
	http.HandleFunc("/update/gauge", handler.RequestHandle)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
