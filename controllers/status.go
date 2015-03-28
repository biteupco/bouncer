package controllers

import "net/http"

func GetStatusHandler(rw http.ResponseWriter, r *http.Request) {
	// simulate check ...
	rw.WriteHeader(http.StatusOK)
}
