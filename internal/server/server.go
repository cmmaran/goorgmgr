package server

import (
	"log"
	"net/http"
)

func Start() {
	log.Printf("Server started")
	router := newRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
