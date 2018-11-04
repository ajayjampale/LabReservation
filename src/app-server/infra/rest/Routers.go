package rest

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"app-server/infra/rest/handlers"
	"strings"
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
		fmt.Println("Registering ", route.Name, " with pattern ", route.Pattern)
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
	fmt.Fprintf(w, "Lab Resource Management Application!")
}

var routes = Routes{

	Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/v1/",
		HandlerFunc: Index,
	},

	Route{
		Name:        "GetResourceTypes",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/v1/resourcetypes",
		HandlerFunc: handlers.GetResourceTypes,
		QueryPairs:  []string{"parent_resource_type", "{parent_resource_type}"},
	},

}

