package main

import (
	"fmt"
	"log"
	"os"

	"github.com/margostino/anfield-api/tooling"
)

func main() {
	if len(os.Args) < 2 {
		message := fmt.Sprintf("Command Not Found!\n" +
			"Commands available: \n" +
			"- server: to start a local server\n" +
			"- schema-gen: to generate a new Graphql Schema")
		log.Panicln(message)
	}
	switch action := os.Args[1]; action {
	case "server":
		tooling.RunLocalServer()
	case "schema-gen":
		tooling.GenerateSchema()
	default:
		log.Printf("command not valid: %s\n", action)
	}
}
