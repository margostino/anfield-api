package tooling

import (
	"log"
	"net/http"

	handler "github.com/margostino/anfield-api/api"
)

func RunLocalServer() {
	http.HandleFunc("/hello", handler.Hello)
	http.HandleFunc("/query", handler.Query)
	http.HandleFunc("/playground", handler.Playground)
	log.Println("Starting anfield-api Server in :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
