package main

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

type endPoint struct {
	method string
	url    string
}

var endpoints = []endPoint{
	{"GET", "/"},
	{"GET", "/status"},
}

func hasEndpoint(r *mux.Router, ep endPoint) bool {
	req, err := http.NewRequest(ep.method, ep.url, nil)
	if err != nil {
		panic(err)
	}

	var match mux.RouteMatch
	ok := r.Match(req, &match)
	return ok
}

func TestAppRoutes(t *testing.T) {
	r := setupServer()

	for _, ep := range endpoints {
		ok := hasEndpoint(r, ep)
		if !ok {
			t.Errorf("%s %s endpoint is not in router", ep.method, ep.url)
		}
	}
}
