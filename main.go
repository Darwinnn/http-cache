package main

import (
	"flag"
	"log"

	"github.com/valyala/fasthttp"
)

var (
	DefaultTTL int64
	addr       string
	MemOpt     bool
	MaxSize    int
)

func init() {
	flag.Int64Var(&DefaultTTL, "ttl", 4294967295, "default time-to-live of cache objects")
	flag.StringVar(&addr, "addr", ":8080", "address to listen on")
	flag.BoolVar(&MemOpt, "memopt", false, "Memory optimization (might increase CPU usage)")
	flag.IntVar(&MaxSize, "maxsize", 4294967295, "maximum upload size (in bytes)")
}

func main() {
	flag.Parse()

	// this is the in-mem cache variable that will store all the data
	var cache Cache
	cache.StartCleanUpWorker()

	router := cache.BuildRouter()

	log.Printf("Listening on %s\n", addr)
	server := &fasthttp.Server{
		MaxRequestBodySize: MaxSize,
		ReduceMemoryUsage:  MemOpt,
		Handler:            router.Handler,
	}
	log.Fatal(server.ListenAndServe(addr))
}
