package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nadavbm/nogobk/api/dat"
	"go.uber.org/zap"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func loginHandler(ctx *Context) {
	l := ctx.Log

	var c Credentials

	err := ctx.RequestUnmarshal(&c)
	if err != nil {
		return
	}

	u, err := dat.GetUserByEmail(l, c.Email)
	if err != nil {
		l.Info("failed to get user by email: ", zap.Error(err))
	}

	expire := time.Now()
	expire = expire.Add(3 * time.Minute)
	s := dat.Session{
		UserId:  u.Id,
		Token:   "sometoken",
		Csrf:    "csrf token",
		Expires: expire,
	}

	err = s.CreateSession(l)
	if err != nil {
		l.Info("failed to create session in database:", zap.Error(err))
	}

	l.Info("new user login", zap.String("email:", u.Email))

	// html template
	fmt.Fprintf(ctx.Writer, "login handler\nurl path: %s\n", ctx.Request.URL.Path[1:])
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "logout handler\nurl path: %s\n", r.URL.Path[1:])
}
