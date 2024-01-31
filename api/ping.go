package handler

import (
	"fmt"
	"net/http"

	"github.com/margostino/anfield-api/auth"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	if auth.IsAuthorized(r) {
		fmt.Fprintf(w, "pong\n")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
