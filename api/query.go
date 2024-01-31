package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/margostino/anfield-api/auth"
	"github.com/margostino/anfield-api/graph"
	"github.com/margostino/anfield-api/metrics"
	// "github.com/margostino/anfield-api/graph/generated"
	// "github.com/margostino/anfield-api/metrics"
)

var server = newServer()

func Query(w http.ResponseWriter, r *http.Request) {
	if auth.IsAuthorized(r) {
		go metrics.PublishRequest(r)
		server.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func newServer() *handler.Server {
	c := graph.Config{Resolvers: &graph.Resolver{}}
	return handler.NewDefaultServer(graph.NewExecutableSchema(c))
}
