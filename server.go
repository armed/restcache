package main

import (
	"flag"
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

var (
	instance cache.Cache
	port     = flag.Uint("port", 8080, "http port to listen")
)

func main() {
	flag.Parse()
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

	log.Printf("Server started at port: %d", *port)

	host := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(host, nil))
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
