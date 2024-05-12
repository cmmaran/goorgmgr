package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/cmmaran/goorgmgr/internal/layer"
	"github.com/gorilla/mux"
)

type Server struct {
	bl layer.Layer
}

func NewServer(l layer.Layer) *Server {
	return &Server{
		bl: l,
	}
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func (s *Server) Start() {
	log.Printf("Server started")
	router := s.newRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func (s *Server) newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

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
			s.CreateEmployee,
		},
		Route{
			"GetEmployee",
			strings.ToUpper("Get"),
			"/api/v1/employee",
			s.GetEmployee,
		},
		Route{
			"CreateUsersWithListInput",
			strings.ToUpper("Post"),
			"/api/v1/employee/createWithList",
			s.CreateUsersWithListInput,
		},
		Route{
			"DeleteEmployee",
			strings.ToUpper("Delete"),
			"/api/v1/employee/{id}",
			s.DeleteEmployee,
		},
		Route{
			"GetEmployeeByID",
			strings.ToUpper("Get"),
			"/api/v1/employee/{id}",
			s.GetEmployeeByID,
		},
		Route{
			"UpdateEmployee",
			strings.ToUpper("Put"),
			"/api/v1/employee/{id}",
			s.UpdateEmployee,
		},
	}

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
