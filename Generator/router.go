package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route Defines the routes available and the handler function responsible for
// processing them.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes A collection of defined Route(s)
type Routes []Route

var routes = Routes{
	Route{
		"Generate",
		"POST",
		"/Generate",
		Generate,
	},
}

// NewRouter Creates a new router and establishes the routes defined.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
