package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/and-volkov/learn-go-with-tests.at.study/adapters"
	"github.com/and-volkov/learn-go-with-tests.at.study/adapters/httpserver"
	"github.com/and-volkov/learn-go-with-tests.at.study/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "8080"
		dockerFilePath = "./cmd/httpserver/Dockerfile"
		baseURL        = fmt.Sprintf("http://localhost:%s", port)
		driver         = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, driver)
}
