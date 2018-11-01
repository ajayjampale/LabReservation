/*
 * Lab Resource Management
 *
 * This specification defines the APIs provided by the application to manage/share the lab resources by reserving and releasing it.
 *
 * API version: 1.0
 * Contact: mithun.bs@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

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

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"CreateResourceType",
		strings.ToUpper("Post"),
		"/v1/resourcetype",
		CreateResourceType,
	},

	Route{
		"DeleteResourceType",
		strings.ToUpper("Delete"),
		"/v1/resourcetype",
		DeleteResourceType,
	},

	Route{
		"GetResourceTypes",
		strings.ToUpper("Get"),
		"/v1/resourcetype",
		GetResourceTypes,
	},
}
