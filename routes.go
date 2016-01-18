package main

import "net/http"

// Route defines a single route for the router
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

var routes = []Route{
	Route{
		"GET",
		"/hello",
		Hello,
	},
}
