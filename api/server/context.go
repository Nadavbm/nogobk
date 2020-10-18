package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/nadavbm/nogobk/api/dat"
	"github.com/nadavbm/nogobk/pkg/logger"
	"go.uber.org/zap"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	RBody   []byte
	Log     *logger.Logger
	ReqTime time.Time
	User    *dat.User
	Session *dat.Session
	Context *dat.Context
}

type CtxHandler func(ctx *Context)

func NewRequestContext(l *logger.Logger, w http.ResponseWriter, r *http.Request) *Context {
	l.Info("new context")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		l.Error("request cannot convert to byte:", zap.Error(err))
	}

	ctx := Context{
		Writer:  w,
		Request: r,
		Log:     l,
		ReqTime: time.Now(),
		RBody:   b,
	}

	fmt.Println("the context request:", &ctx.Request, "\nthe request body:", &ctx.Request.Body, "\nthe request URL:", &ctx.Request.URL, "\nrequest context:", ctx.Request.Context())
	return &ctx
}

func (c *Context) GetContext() error {
	ctx, err := dat.NewDatabaseContext()
	if err != nil {
		return err
	}

	c.Context = ctx
	return nil
}

func ContextHandler(l *logger.Logger, h CtxHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("context handler:", h)
		ctx := NewRequestContext(l, w, r)

		ctx.Request.ParseForm()

		fmt.Println("handler context body", ctx)
		h(ctx)
	}
}

// Helpers

func (c *Context) RequestUnmarshal(j interface{}) error {
	err := json.Unmarshal(c.RBody, j)
	if err == io.EOF {
		return err
	}
	fmt.Println("unmarshal request body:", j)
	return err
}
