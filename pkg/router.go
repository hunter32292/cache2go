package cache

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router to handle the index of the REST api
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// Add a logger to the http.Handler
	var handler http.Handler
	for _, route := range routes {
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
