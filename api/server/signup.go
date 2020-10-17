package server

import (
	"fmt"
	"net/http"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "signup handler\nurl path: %s\n", r.URL.Path[1:])
}
