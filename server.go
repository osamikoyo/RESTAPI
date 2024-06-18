package _package

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpserver *http.Server
}

func (s Server) Run(port string) error {
	s.httpserver = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpserver.ListenAndServe()
}
func (s Server) shutsown(ctx context.Context) error {
	return s.httpserver.Shutdown(ctx)
}
