package server

import (
	"fmt"
	"net/http"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func loginHandler(ctx *Context) {
	fmt.Println("in login handler:", ctx)
	var c Credentials

	err := ctx.RequestUnmarshal(&c)
	if err != nil {
		return
	}

	fmt.Println("login handler after unmarshal:", c)
}

func loginnHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login handler\nurl path: %s\n", r.URL.Path[1:])
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "logout handler\nurl path: %s\n", r.URL.Path[1:])
}
