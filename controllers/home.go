package controllers

import (
	"fmt"
	"net/http"
)

func GetHomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello, world")
}
