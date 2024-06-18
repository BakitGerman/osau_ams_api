package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         cfg.HTPP.Host + ":" + strconv.Itoa(int(cfg.HTPP.Port)),
			Handler:      handler,
			ReadTimeout:  cfg.HTPP.ReadTimeout,
			WriteTimeout: cfg.HTPP.WriteTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
