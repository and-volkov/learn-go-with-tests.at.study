package interactons_test

import (
	"testing"

	"github.com/and-volkov/learn-go-with-tests.at.study/domain/interactons"
	"github.com/and-volkov/learn-go-with-tests.at.study/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactons.Greet),
	)
}
