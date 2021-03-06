package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/nadavbm/nogobk/api/dat"
	"github.com/nadavbm/nogobk/pkg/logger"
	"go.uber.org/zap"
)

type Context struct {
	*http.Request
	Writer  http.ResponseWriter
	RBody   []byte
	Log     *logger.Logger
	ReqTime time.Time
	User    *dat.User
	Session *dat.Session
	dat.Context
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
	return &ctx
}

func (c *Context) GetDBContext() error {
	ctx, err := dat.NewDatabaseContext()
	if err != nil {
		return err
	}

	c.Context = *ctx
	return nil
}

func ContextHandler(l *logger.Logger, h CtxHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := NewRequestContext(l, w, r)

		/*
			err := ctx.GetDBContext()
			if err != nil {
				l.Error("fail to get db context:", zap.Error(err))
				return
			}
		*/

		ctx.Request.ParseForm()

		h(ctx)
	}
}

// Helpers

func (c *Context) RequestUnmarshal(j interface{}) error {
	err := json.Unmarshal(c.RBody, j)
	if err == io.EOF {
		return err
	}

	return err
}
