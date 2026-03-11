// Package main
package main

import (
	"log"
	"net/http"

	"github.com/and-volkov/learn-go-with-tests.at.study/adapters/httpserver"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
