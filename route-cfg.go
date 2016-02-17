//all configuration of the routing table goes into here
package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

//routes configuration
var routes = Routes{
	Route{"Index", "GET", "/", index},
	Route{"Todos", "GET", "/todos", todos},
}
