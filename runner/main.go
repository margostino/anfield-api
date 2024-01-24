package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/margostino/anfield-api/api"
)

func main() {
	http.HandleFunc("/hello", handler.Hello)
	http.HandleFunc("/query", handler.Query)
	http.HandleFunc("/playground", handler.Playground)

	helloEndpoint := "http://localhost:8080/hello"
	queryEndpoint := "http://localhost:8080/query"
	playgroundEndpoint := "http://localhost:8080/playground"
	endpoints := fmt.Sprintf("Endpoints:\n👋 %s\n⚡️ %s\n🔎 %s", helloEndpoint, queryEndpoint, playgroundEndpoint)
	log.Println("Starting anfield-api server...")
	log.Println(endpoints)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
