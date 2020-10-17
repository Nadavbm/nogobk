package server

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "root handler\nurl path: %s\n", r.URL.Path[1:])
}
