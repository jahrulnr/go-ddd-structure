package httpserver

import (
	"context"
	"ddd-impl/config"
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":80"
	_defaultShutdownTimeout = 5 * time.Second
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
	cfg             *config.Config
}

func New(h http.Handler, cfg *config.Config, opts ...Options) *Server {
	h2s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}
	handler := h2c.NewHandler(h, h2s)

	s := &Server{
		server: &http.Server{
			Handler:      handler,
			ReadTimeout:  _defaultReadTimeout,
			WriteTimeout: _defaultWriteTimeout,
			Addr:         _defaultAddr,
		},
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
		cfg:             cfg,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.Start()

	return s
}

func (s *Server) Start() {
	go func() {
		if s.cfg.Server.UseSSL {
			cert := fmt.Sprintf("%s%s", s.cfg.App.BaseDir, s.cfg.Server.SSLCert)
			key := fmt.Sprintf("%s%s", s.cfg.App.BaseDir, s.cfg.Server.SSLKey)
			log.Printf("Server run use tls is listening on %s \n", s.server.Addr)
			s.notify <- s.server.ListenAndServeTLS(cert, key)
		} else {
			log.Printf("Server is listening on %s \n", s.server.Addr)
			s.notify <- s.server.ListenAndServe()
		}
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		return err
	}

	return err
}
