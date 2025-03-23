// Package httpserver provides an HTTP server implementation.
package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Server wraps an HTTP server instance.
type Server struct {
	httpServer *http.Server
}

// New creates a new HTTP server instance.
func New(addr string, router *gin.Engine) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:              addr,
			Handler:           router,
			ReadHeaderTimeout: time.Second,
		},
	}
}

// Start launches the HTTP server.
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// Stop gracefully shuts down the server.
func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
