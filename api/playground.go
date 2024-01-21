package handler

import (
	"github.com/99designs/gqlgen/graphql/playground"
	// "github.com/margostino/anfield-api/metrics"
	"net/http"
	"os"
)

var playgroundUrl = os.Getenv("PLAYGROUND_ENDPOINT")
var playgroundServer = playground.Handler("GraphQL playground", playgroundUrl)

func Playground(w http.ResponseWriter, r *http.Request) {
	// go metrics.PublishRequest(r)
	playgroundServer.ServeHTTP(w, r)
}
