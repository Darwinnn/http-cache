package main

import (
	"flag"
	"log"
	"net/http"
	"sync"
)

var (
	DefaultTTL int64
	addr       string
)

func init() {
	flag.Int64Var(&DefaultTTL, "ttl", 4294967295, "default time-to-live of cache objects")
	flag.StringVar(&addr, "addr", ":8080", "address to listen on")
}

func main() {
	flag.Parse()

	// this is the in-mem cache variable that will store all the data
	cache := &Cache{
		Mutex: &sync.Mutex{},
		Map:   make(map[string]CacheElement),
	}
	cache.StartCleanUpWorker()

	router := cache.BuildRouter()

	log.Printf("Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
