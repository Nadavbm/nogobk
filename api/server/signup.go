package server

import (
	"fmt"

	"github.com/nadavbm/nogobk/api/dat"

	"go.uber.org/zap"
)

type SignUp struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Confirm  string `json:"confirm"`
}

func signupHandler(ctx *Context) {
	l := ctx.Log

	var s SignUp

	err := ctx.RequestUnmarshal(&s)
	if err != nil {
		return
	}

	u := dat.User{
		Name:     s.Name,
		Email:    s.Email,
		Password: s.Password,
	}
	err = u.CreateUsers(l)
	if err != nil {
		l.Info("there was an error while creating a user", zap.Error(err))
	}
	l.Info("user created", zap.String("with email address - ", u.Email))

	// html template
	fmt.Fprintf(ctx.Writer, "signup handler\nurl path: %s\n", ctx.Request.URL.Path[1:])

}
