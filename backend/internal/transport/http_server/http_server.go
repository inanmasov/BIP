package http_server

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port, cert, key string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
		//MaxHeaderBytes: 1 << 20, // 1 MB
		//ReadTimeout:    10 * time.Second,
		//WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServeTLS(cert, key)
	// return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
