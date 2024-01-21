package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/margostino/anfield-api/fetcher"
	"github.com/margostino/anfield-api/graph"
	"github.com/margostino/anfield-api/graph/generated"
	"github.com/margostino/anfield-api/metrics"
)

var server = newServer()

func Query(w http.ResponseWriter, r *http.Request) {
	go metrics.PublishRequest(r)
	server.ServeHTTP(w, r)
}

func newServer() *handler.Server {
	c := generated.Config{Resolvers: &graph.Resolver{Fetcher: fetcher.Fetch}}
	return handler.NewDefaultServer(generated.NewExecutableSchema(c))
}
