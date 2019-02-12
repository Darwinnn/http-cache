package main

import (
	"flag"
	"log"

	"github.com/valyala/fasthttp"
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
	var cache Cache
	//cache.StartCleanUpWorker()

	router := cache.BuildRouter()

	log.Printf("Listening on %s\n", addr)
	log.Fatal(fasthttp.ListenAndServe(addr, router.Handler))
}
