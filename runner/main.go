package main

import (
	"log"
	"net/http"

	handler "github.com/margostino/anfield-api/api"
)

func main() {
	http.HandleFunc("/hello", handler.Hello)
	log.Println("Starting anfield-api server in :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
