package gospecsgreet_test

import (
	"testing"

	gospecsgreet "github.com/and-volkov/learn-go-with-tests.at.study"
	"github.com/and-volkov/learn-go-with-tests.at.study/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(gospecsgreet.Greet),
	)
}
