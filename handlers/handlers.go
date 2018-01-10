package handlers

import (
	"github.com/gorilla/mux"
)

// Router regiser necessary routes and returns an instance of a router

func Router() *mux.Router {
	r := mux.Router()
	r.HandleFunc("/home", home).Methods("GET")
	return r
}