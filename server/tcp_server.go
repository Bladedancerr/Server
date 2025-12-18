package server

import (
	"github.com/Bladedancerr/server/transport"
)

// tcp server
// implements server interface
type TCPServer struct {
	transport transport.Transport
	opts      ServerOpts
}

func NewTCPServer(opts ServerOpts) *TCPServer {
	return &TCPServer{
		opts:      opts,
		transport: transport.NewTCPTransport(opts.ListenAddr),
	}
}

func (s *TCPServer) Start() error {
	return s.transport.Listen()
}

func (s *TCPServer) Stop() error {
	return s.transport.Close()
}
