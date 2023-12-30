package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type RouterServerOptions struct {
	Port         string
	WriteTimeout *time.Duration
	ReadTimeout  *time.Duration
}

type HTTPServer struct {
	options RouterServerOptions
	router  *mux.Router
}

func NewRouterServer(options RouterServerOptions) *HTTPServer {
	return &HTTPServer{
		options: options,
		router:  mux.NewRouter(),
	}
}

func (s *HTTPServer) Router() *mux.Router {
	return s.router
}

func (s *HTTPServer) Start() error {
	srv := &http.Server{
		Handler: s.router,
		Addr:    s.options.Port,
	}

	if s.options.WriteTimeout != nil {
		srv.WriteTimeout = *s.options.WriteTimeout
	}

	if s.options.ReadTimeout != nil {
		srv.ReadTimeout = *s.options.ReadTimeout
	}

	return srv.ListenAndServe()

}
