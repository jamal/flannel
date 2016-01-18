package main

import (
	"log"
	"net/http"

	"github.com/jamal/flannel"
)

var routes = []flannel.Route{
	flannel.Route{
		"GET",
		"/hello",
		Hello,
	},
}

func main() {
	r := flannel.New()
	r.Use(FunMiddleware)
	r.Use(OtherMiddleware)
	r.AddRoutes(routes)
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Hello is an example endpoint
func Hello(w http.ResponseWriter, r *http.Request) {
	flannel.LogInfo(r, "Hello was called")
	flannel.Write(w, http.StatusOK, "hello")
}

// FunMiddleware is an example of a middleware method.
func FunMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flannel.LogInfo(r, "Running fun middleware")
		next(w, r)
	}
}

// OtherMiddleware is another example of a middleware method.
func OtherMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flannel.LogInfo(r, "Running other middleware")
		next(w, r)
	}
}
