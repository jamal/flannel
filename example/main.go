package main

import (
	"fmt"
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
	flannel.Route{
		"POST",
		"/login",
		Login,
	},
}

var counter int

type myContext struct {
	Name    string
	Counter int
}

func main() {
	r := flannel.New()
	r.Use(NameMiddleware)
	r.Use(CounterMiddleware)
	r.AddRoutes(routes)
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Hello is an example endpoint
func Hello(w http.ResponseWriter, r *http.Request) {
	ctx := &myContext{}
	flannel.Context(r, &ctx)
	flannel.LogInfo(r, "Hello was called by %s", ctx.Name)

	flannel.Write(w, http.StatusOK, fmt.Sprintf("Hello, %s. I've been called %d times.", ctx.Name, ctx.Counter))
}

// NameMiddleware is an example of a middleware method.
func NameMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &myContext{}
		flannel.Context(r, &ctx)
		ctx.Name = "World"

		next(w, r)
	}
}

// CounterMiddleware is another example of a middleware method.
func CounterMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &myContext{}
		flannel.Context(r, &ctx)
		counter++
		ctx.Counter = counter

		next(w, r)
	}
}
