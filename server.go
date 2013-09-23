package main

import (
	"flag"
	"fmt"
	"github.com/armed/restcache/cache"
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

	http.Handle("/", newRouter())

	log.Printf("Server started at port: %d", *port)

	host := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(host, nil))
}
