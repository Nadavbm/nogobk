package server

import (
	"fmt"
	"net/http"

	"github.com/nadavbm/nogobk/api/dat"
)

func profileHandler(w http.ResponseWriter, r *http.Request) {
	u := &dat.User{}
	fmt.Fprintf(w, "profile handler for user: %s\nurl path: %s\n", u.Name, r.URL.Path[1:])
}
