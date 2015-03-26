package auth

import (
	"fmt"
	"net/http"
)

func SignupHandler(rw http.ResponseWriter) {
	fmt.Fprintln(rw, "you are signing up")
}