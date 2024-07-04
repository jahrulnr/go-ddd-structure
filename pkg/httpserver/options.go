package httpserver

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

type Options func(*Server)

func Port(host, port string) Options {
	return func(s *Server) {
		s.server.Addr = net.JoinHostPort(host, port)
	}
}

func ReadTimeout(timeout time.Duration) Options {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Options {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

func ShutdownTimeout(timeout time.Duration) Options {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}

func UseSsl() Options {
	return func(s *Server) {
		if !s.cfg.Server.UseHTTP2 {
			s.server.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)
		}

		if s.cfg.Server.UseSSL && s.cfg.App.Environment == "preproduction" || s.cfg.App.Environment == "production" {
			cfgClient := &tls.Config{
				MinVersion:       tls.VersionTLS12,
				CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
				CipherSuites: []uint16{
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
					tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_RSA_WITH_AES_256_CBC_SHA,
					// Best disabled, as they don't provide Forward Secrecy,
					tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				},
			}
			s.server.TLSConfig = cfgClient
		}
	}
}
