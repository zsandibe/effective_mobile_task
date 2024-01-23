package server

import (
	"net/http"

	"github.com/zsandibe/effective_mobile_task/config"
	"github.com/zsandibe/effective_mobile_task/pkg"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(config config.Config, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           config.Server.Host + ":" + config.Server.Port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		RequestTimeout: config.Server.RequestTimeout,
	}
	pkg.InfoLog.Printf("Starting server on  %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
