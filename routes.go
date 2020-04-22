package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ValueDump",
		"GET",
		"/dump",
		ValueDump,
	},
	Route{
		"ValueFind",
		"GET",
		"/find/{value}",
		ValueFind,
	},
	Route{
		"ValueAdd",
		"POST",
		"/add",
		ValueAdd,
	},
	Route{
		"ValueRemove",
		"DELETE",
		"/remove/{value}",
		ValueRemove,
	},
}
