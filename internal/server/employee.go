package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cmmaran/goorgmgr/api/model"
	"github.com/gorilla/mux"
)

func renderJSON(w http.ResponseWriter, data any) (int, error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(b))
	return w.Write(b)
}

func (s *Server) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var rec model.Employee
	err := dec.Decode(&rec)
	if err != nil {
		log.Print(err)
	}

	emp := s.bl.Employee().CreateEmployee(&rec)
	renderJSON(w, emp)
}

func (s *Server) CreateUsersWithListInput(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (s *Server) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		// ... handle error
		panic(err)
	}

	emp := s.bl.Employee().DeleteEmployee(int64(i))
	renderJSON(w, emp)
}

func (s *Server) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		// ... handle error
		panic(err)
	}

	emp := s.bl.Employee().GetEmployeeByID(int64(i))
	renderJSON(w, emp)
}

func (s *Server) GetEmployee(w http.ResponseWriter, r *http.Request) {
	emps := s.bl.Employee().GetEmployee()
	renderJSON(w, emps)
}

func (s *Server) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var rec model.Employee
	err := dec.Decode(&rec)
	if err != nil {
		log.Print(err)
	}

	id := mux.Vars(r)["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		// ... handle error
		panic(err)
	}

	emp := s.bl.Employee().UpdateEmployee(int64(i), &rec)
	renderJSON(w, emp)
}
