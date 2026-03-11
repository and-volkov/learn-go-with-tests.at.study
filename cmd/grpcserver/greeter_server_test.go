package main_test

import (
	"fmt"
	"testing"

	"github.com/and-volkov/learn-go-with-tests.at.study/adapters"
	"github.com/and-volkov/learn-go-with-tests.at.study/adapters/grpcserver"
	"github.com/and-volkov/learn-go-with-tests.at.study/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port   = "50051"
		driver = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, "grpcserver")
	specifications.GreetSpecification(t, &driver)
}
