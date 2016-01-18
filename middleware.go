package flannel

import "net/http"

// Middleware definition
type Middleware func(next http.HandlerFunc) http.HandlerFunc

// middleware is executed in the order listed
var middleware = []Middleware{}
