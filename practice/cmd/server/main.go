package main

import (
	"net/http"

	"github.com/s0ran/go-rookie-gym/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", handler.UserHandler)
	http.ListenAndServe(":8080", mux)
}
