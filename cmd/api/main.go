package main

import (
	"net/http"

	"github.com/pellegrino/vivsemo-api/internal/vivsemoapi"
)

func main() {
	twirpHandler := vivsemoapi.NewVivsemoApiServer()
	// You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
	mux := http.NewServeMux()
	// The generated code includes a method, PathPrefix(), which
	// can be used to mount your service on a mux.
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	http.ListenAndServe(":8080", mux)
}
