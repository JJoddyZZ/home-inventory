package server

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Server struct {
	Server *http.Server
}

func NewServer(port string, h http.Handler) *Server {
	return &Server{
		Server: &http.Server{
			Addr:    ":" + port,
			Handler: h,
		},
	}
}

func (s *Server) Serve() {
	err := s.Server.ListenAndServe()
	if err == http.ErrServerClosed {
		log.Info().Msg("normal shutdown, no more connections accepted")
	} else {
		log.Error().Err(err).Msg("abnormal shutdown")
	}
}

func (s *Server) Shutdown(ctx context.Context) {
	if err := s.Server.Shutdown(ctx); err != nil {
		log.Panic().Err(err).Msg("abnormal shutdown")
	} else {
		log.Info().Msg("application shutdowned")
	}
}
