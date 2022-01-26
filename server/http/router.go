package http

import "github.com/gorilla/mux"

// Router route the service using github.com/gorilla/mux.
type Router func(*mux.Router)

// Routers registers routers.
func Routers(routers ...Router) Option {
	return func(s *Server) {
		for _, router := range routers {
			router(s.router)
		}
	}
}
