package router

import (
	"net/http"
)

// Router embedding the standard Go HTTP ServeMux routing capabilities
type Router struct {
	*http.ServeMux
}

// NewRouter instantiates a clean, dedicated multiplexer gateway engine
func NewRouter() *Router {
	return &Router{
		ServeMux: http.NewServeMux(),
	}
}

// HandleFunc registers a specific pattern matching endpoint onto our router
func (r *Router) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.ServeMux.HandleFunc(pattern, handler)
}
