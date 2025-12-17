package main

//udp server
// must implement server
type UDPServer struct {
	transport Transport
	opts      ServerOpts
}

func NewUDPServer(opts ServerOpts) *UDPServer {
	return &UDPServer{
		opts:      opts,
		transport: NewUDPTransport(opts.ListenAddr),
	}
}

func (s *UDPServer) Start() error {
	return nil
}

func (s *UDPServer) Stop() error {
	return nil
}
