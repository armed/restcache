package main

import (
	"fmt"
	"github.com/armed/restcache/cache"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	defaultDuration = "1h"
)

var instance cache.Cache

func main() {
	instance = cache.New(defaultDuration)

	r := mux.NewRouter()

	r.HandleFunc("/favicon.ico", notFound)

	r.HandleFunc("/{key}/{duration}", putData).
		Methods("POST")

	r.HandleFunc("/{key}", putData).
		Methods("POST")

	r.HandleFunc("/{key}", getData).
		Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func putData(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	r.Body.Close()

	vars := mux.Vars(r)

	err = instance.Put(vars["key"], string(bytes), vars["duration"])

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{ "message": "%s" }`, err)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if value, ok := instance.Get(vars["key"]); ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, value)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
