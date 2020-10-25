package server

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/nadavbm/nogobk/pkg/logger"
	"go.uber.org/zap"
)

type Server struct {
	Logger     *logger.Logger
	Mux        *http.ServeMux
	HTTPServer *http.Server
}

var tpl *template.Template

func NewServer(l *logger.Logger, listenAddress string) *Server {
	s := &Server{
		Logger: l,
	}

	r, err := CreateApiRouter(l)
	if err != nil {
		l.Error("failed to create mux router")
	}

	s.Mux = http.NewServeMux()
	s.Mux.Handle("/", r)

	s.HTTPServer = &http.Server{
		Addr: listenAddress,
	}

	return s
}

func (s *Server) Run() error {
	logger := logger.SetLogger()

	err := s.HTTPServer.ListenAndServe()
	if err != nil {
		logger.Error("cannot run http server - listen and serve", zap.Error(err))
	}

	return nil
}

func CreateApiRouter(l *logger.Logger) (*mux.Router, error) {
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/login", ContextHandler(l, loginHandler)).Methods("POST", "GET")
	r.HandleFunc("/signup", ContextHandler(l, signupHandler)).Methods("POST", "GET")
	r.HandleFunc("/profile/{id}", ContextHandler(l, profileHandler)).Methods("GET")
	r.HandleFunc("/logout", logoutHandler)

	http.Handle("/", r)
	return r, nil
}
