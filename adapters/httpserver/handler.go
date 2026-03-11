package httpserver

import (
	"fmt"
	"net/http"

	"github.com/and-volkov/learn-go-with-tests.at.study/domain/interactons"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, interactons.Greet(name)) //nolint:errcheck
}
