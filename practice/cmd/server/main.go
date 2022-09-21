package main

import (
	"fmt"
	"net/http"

	"github.com/s0ran/go-rookie-gym/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", handler.UserHandler)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Errorf("failed to start server: %v", err)
	}
}
