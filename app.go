package main

import (
	"flag"
	"fmt"
	"net/http"
	"log"

	"github.com/gobbl/bouncer/controllers"
	"github.com/gobbl/bouncer/models"
	"github.com/gorilla/mux"
)

var port int

func init() {
	const (
		defaultPort = 8080
		portUsage   = "port number"
	)
	flag.IntVar(&port, "port", defaultPort, portUsage)
}

func setupServer() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", controllers.GetHomeHandler).Methods("GET")
	r.HandleFunc("/status", controllers.GetStatusHandler).Methods("GET")

	/*
		signup := r.Path("/signup").Subrouter()
		signup.Methods("POST").HandlerFunc(signHandler)

		// Posts collection
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
	return r

}

func main() {
	flag.Parse()

	r := setupServer()

	models.Start()
	defer models.Close()

	fmt.Printf("Starting server on port:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
