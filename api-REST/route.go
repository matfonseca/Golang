package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRoute() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"Get",
		"/",
		Index,
	},
	Route{
		"MovieList",
		"Get",
		"/movies",
		MovieList,
	},
	Route{
		"MovieShow",
		"Get",
		"/movie/{id}",
		MovieShow,
	},
	Route{
		"MovieAdd",
		"POST",
		"/movie",
		MovieAdd,
	},
}
