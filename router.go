package flannel

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Router thing
type Router struct {
	*mux.Router
	middleware []Middleware
}

// Route defines a single route for the router
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

// New returns a new instance of Router.
func New() *Router {
	router := &Router{
		mux.NewRouter(),
		make([]Middleware, 0),
	}
	router.StrictSlash(true)
	return router
}

// Use registers a new Middleware for this router.
func (r *Router) Use(middleware Middleware) {
	r.middleware = append(r.middleware, middleware)
}

// AddRoute registers a route.
func (r *Router) AddRoute(route Route) {
	r.Methods(route.Method).
		Path(route.Path).
		Handler(r.handler(route.Handler))
}

// AddRoutes registers a set of routes with this router.
func (r *Router) AddRoutes(routes []Route) {
	for _, route := range routes {
		r.AddRoute(route)
	}
}

func (r *Router) handler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Prepare handler stack
		stack := handler
		for i := len(r.middleware) - 1; i >= 0; i-- {
			stack = r.middleware[i](stack)
		}

		// Create new request context
		start := time.Now()
		rid := genReqID()
		setReqID(req, rid)
		cw := &responseWriter{w, 200}
		cw.Header().Set("Request-Id", rid)

		stack(cw, req)

		logAccess(req, "%s %s status=%d remote=%s time=%s",
			req.Method,
			req.RequestURI,
			cw.Status,
			req.RemoteAddr,
			time.Since(start),
		)

		deleteContext(req)
	}
}

// Write helper to marshal an object to JSON and write it to the ResponseWriter
func Write(w http.ResponseWriter, code int, out interface{}) {
	response, err := json.Marshal(out)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// ResponseWriter wraps http.ResponseWriter but adds a few things we need for
// some of the middleware
type responseWriter struct {
	http.ResponseWriter
	Status int
}

// WriteHeader logs the Status code and calls ResponseWriter.WriteHeader
func (w *responseWriter) WriteHeader(code int) {
	w.Status = code
	w.ResponseWriter.WriteHeader(code)
}

func genReqID() string {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		LogError(nil, "failed to generate request id: %v", err)
	}
	return hex.EncodeToString(b)
}
