package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

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
    fmt.Println("Starting server on :8080")
    http.ListenAndServe(":8080", r)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request){
	fmt.Fprintln(rw, "Hello, world")
}

func signHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "signHandler")
}