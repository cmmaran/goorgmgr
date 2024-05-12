/*
 * Org Manager - OpenAPI 3.0
 *
 * This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about Swagger at [https://swagger.io](https://swagger.io). In the third iteration of the pet store, we've switched to the design first approach! You can now help us improve the API whether it's by making changes to the definition itself or to the code. That way, with time, we can improve the API in general, and expose some of the new features in OAS3.  _If you're looking for the Swagger 2.0/OAS 2.0 version of Petstore, then click [here](https://editor.swagger.io/?url=https://petstore.swagger.io/v2/swagger.yaml). Alternatively, you can load via the `Edit > Load Petstore OAS 2.0` menu option!_  Some useful links: - [The Pet Store repository](https://github.com/swagger-api/swagger-petstore) - [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)
 *
 * API version: 1.0.11
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

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

func newRouter() *mux.Router {
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
		"/api/v1/",
		Index,
	},

	Route{
		"CreateEmployee",
		strings.ToUpper("Post"),
		"/api/v1/employee",
		CreateEmployee,
	},

	Route{
		"CreateUsersWithListInput",
		strings.ToUpper("Post"),
		"/api/v1/employee/createWithList",
		CreateUsersWithListInput,
	},

	Route{
		"DeleteEmployee",
		strings.ToUpper("Delete"),
		"/api/v1/employee/{id}",
		DeleteEmployee,
	},

	Route{
		"GetEmployeeByID",
		strings.ToUpper("Get"),
		"/api/v1/employee/{id}",
		GetEmployeeByID,
	},

	Route{
		"LoginUser",
		strings.ToUpper("Get"),
		"/api/v1/employee",
		LoginUser,
	},

	Route{
		"UpdateUser",
		strings.ToUpper("Put"),
		"/api/v1/employee/{id}",
		UpdateUser,
	},
}
