package function_option

import "time"

type Server struct {
	timeout time.Duration
}

func (s *Server) DefaultServer() Server {
	return Server{
		timeout: 10 * time.Second,
	}
}

type Option func(*Server)

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func NewServer(opts ...Option) Server {
	server := Server{}
	server = server.DefaultServer()
	for _, opt := range opts {
		opt(&server)
	}
	return server
}
