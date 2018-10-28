package rest

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"app-server/infra/rest/handlers"
)

//Route defines a unique route for a REST request
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	QueryPairs  []string
}

//Routes define an array of routes
type Routes []Route

//NewLRMRouter returns a new Router which routes the LRM REST request to the unique Handler
func NewLRMRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = handlers.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler).
			Queries(route.QueryPairs...)

	}

	return router
}

//Index returns a welcome message
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
/*
	Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: Index,
	},
*/
/*	Route{
		Name:        "Shutdown",
		Method:      strings.ToUpper("POST"),
		Pattern:     "/shutdown",
		HandlerFunc: Handler,
		QueryPairs:  []string{"id", "{id}"},
	},*/

}

