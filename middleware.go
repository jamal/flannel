package main

import "net/http"

// Middleware definition
type Middleware func(handler http.HandlerFunc) http.HandlerFunc

// middleware is executed in the order listed
var middleware = []Middleware{}
