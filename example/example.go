package main

import "net/http"

// Hello is an example endpoint
func Hello(w http.ResponseWriter, r *http.Request) {
	flannel.Write(w, http.StatusOK, "hello")
}
