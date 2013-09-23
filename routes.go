package main

import (
	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/favicon.ico", favicon)
	r.HandleFunc("/{key}/{duration}", putData).
		Methods("POST")
	r.HandleFunc("/{key}", putData).
		Methods("POST")
	r.HandleFunc("/{key}", getData).
		Methods("GET")

	return r
}
