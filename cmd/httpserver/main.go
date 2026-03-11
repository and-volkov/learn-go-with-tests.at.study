// Package main
package main

import (
	"log"
	"net/http"

	gospecsgreet "github.com/and-volkov/learn-go-with-tests.at.study"
)

func main() {
	handler := http.HandlerFunc(gospecsgreet.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
