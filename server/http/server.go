package http

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Server is an HTTP test server.
type Server struct {
	addr    string
	router  *mux.Router
	timeout time.Duration
}

// NewServer returns a test server.
func NewServer(addr string, options ...Option) *Server {
	s := Server{
		addr:    addr,
		router:  mux.NewRouter(),
		timeout: 15 * time.Second, // TODO: modify timeout using options
	}

	for _, option := range options {
		option(&s)
	}

	return &s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) ListenAndServe() error {
	srv := http.Server{
		Addr:         s.addr,
		ReadTimeout:  s.timeout,
		WriteTimeout: s.timeout,
		Handler:      s.router,
	}

	return srv.ListenAndServe()
}

// An Option is options func that can be set up when creating a new server.
type Option func(*Server)
