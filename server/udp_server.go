package server

import transport "github.com/Bladedancerr/server/transport"

//udp server
// must implement server
type UDPServer struct {
	transport transport.Transport
	opts      ServerOpts
}

func NewUDPServer(opts ServerOpts) *UDPServer {
	return &UDPServer{
		opts:      opts,
		transport: transport.NewUDPTransport(opts.ListenAddr),
	}
}

func (s *UDPServer) Start() error {
	return s.transport.Listen()
}

func (s *UDPServer) Stop() error {
	return s.transport.Close()
}
