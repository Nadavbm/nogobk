package server

import (
	"fmt"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login handler\nurl path: %s\n", r.URL.Path[1:])
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "logout handler\nurl path: %s\n", r.URL.Path[1:])
}
