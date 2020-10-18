package server

import (
	"fmt"
)

type SignUp struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Confirm  string `json:"confirm"`
}

func signupHandler(ctx *Context) {
	fmt.Println("in login handler:", ctx)
	var s SignUp

	err := ctx.RequestUnmarshal(&s)
	if err != nil {
		return
	}

	fmt.Println("signup handler after unmarshal:", s)
	//ctx.User, err = ctx.Users.CreateUser(l, s.Name, s.Email, s.Password)
	fmt.Println("dat user:", ctx.User)

}
