package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

// results
func resultNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func resultError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, `{ "message": "%s" }`, err)
}

func resultOK(w http.ResponseWriter, value string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, value)
}

// favicon
func favicon(w http.ResponseWriter, r *http.Request) {
	resultNotFound(w)
}

// handlers
func putData(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Println(err)
		resultError(w, err)
	}

	vars := mux.Vars(r)

	err = instance.Put(vars["key"], string(bytes), vars["duration"])

	if err != nil {
		log.Println(err)
		resultError(w, err)
	} else {
		resultNotFound(w)
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if value, ok := instance.Get(vars["key"]); ok {
		resultOK(w, value)
	} else {
		resultNotFound(w)
	}
}
