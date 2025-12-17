package main

import "fmt"

// tcp server
// implements server interface
type TCPServer struct {
	transport Transport
	opts      ServerOpts
}

func NewTCPServer(opts ServerOpts) *TCPServer {
	return &TCPServer{
		opts:      opts,
		transport: NewTCPTransport(opts.ListenAddr),
	}
}

func (s *TCPServer) Start() error {
	// for testing
	go func() {
		for msg := range s.transport.Messages() {
			fmt.Println("received:", string(msg))
		}
	}()

	return s.transport.Listen()
}

func (s *TCPServer) Stop() error {
	return s.transport.Close()
}
