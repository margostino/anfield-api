package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/margostino/anfield-api/api"
)

func main() {
	http.HandleFunc("/ping", handler.Hello)
	http.HandleFunc("/query", handler.Query)
	http.HandleFunc("/playground", handler.Playground)

	pingEndpoint := "http://localhost:8080/ping"
	queryEndpoint := "http://localhost:8080/query"
	playgroundEndpoint := "http://localhost:8080/playground"
	endpoints := fmt.Sprintf("Endpoints:\n👋 %s\n⚡️ %s\n🔎 %s", pingEndpoint, queryEndpoint, playgroundEndpoint)
	log.Println("Starting anfield-api server...")
	log.Println(endpoints)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
