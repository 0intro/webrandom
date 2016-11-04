package main

import (
	"flag"
	"net/http"

	"github.com/0intro/webrandom/handlers"
)

var (
	httpAddr = flag.String("http", "localhost:8080", "HTTP listen address")
)

func main() {
	flag.Parse()

	h := handlers.New()

	http.HandleFunc("/", h.Root)
	http.HandleFunc("/null", h.Null)
	http.HandleFunc("/random/", h.Random)
	http.HandleFunc("/zero/", h.Zero)

	http.ListenAndServe(*httpAddr, Recovery(Log(http.DefaultServeMux)))
}
