package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var port int

func init() {
	const (
		defaultPort = 8080
		portUsage   = "port number"
	)
	flag.IntVar(&port, "p", defaultPort, portUsage)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", HomeHandler).Methods("GET")

	signup := r.Path("/signup").Subrouter()
	signup.Methods("POST").HandlerFunc(signHandler)

	// Posts collection
	/*
	   signup := r.Path("/signup").Subrouter()
	   signup.Methods("GET").HandlerFunc(PostsIndexHandler)
	   signup.Methods("POST").HandlerFunc(PostsCreateHandler)

	   // Posts singular
	   login := r.PathPrefix("/login/{id}").Subrouter()
	   login.Methods("GET").Path("/edit").HandlerFunc(PostEditHandler).Schemes("http")
	   login.Methods("GET").HandlerFunc(PostShowHandler)
	   login.Methods("PUT", "POST").HandlerFunc(PostUpdateHandler)
	   login.Methods("DELETE").HandlerFunc(PostDeleteHandler)
	*/
	fmt.Printf("Starting server on port:%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello, world")
}

func signHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "signHandler")
}
