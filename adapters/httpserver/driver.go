// Package httpserver
package httpserver

import (
	"io"
	"log"
	"net/http"
)

type Driver struct {
	BaseURL string
	Client  *http.Client
}

func (d Driver) Greet(name string) (string, error) {
	client := d.Client
	if client == nil {
		client = http.DefaultClient
	}

	res, err := client.Get(d.BaseURL + "/greet?name=" + name)
	if err != nil {
		return "", err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Printf("warning: failed to close response body: %v", err)
		}
	}()

	greeting, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(greeting), nil
}
