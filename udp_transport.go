package main

// udp transport
// must implement transport interface
type UDPTransport struct {
	listenAddr string
	messagech  chan []byte
	quitch     chan struct{}
}

func NewUDPTransport(listenAddr string) *UDPTransport {
	return &UDPTransport{
		listenAddr: listenAddr,
		messagech:  make(chan []byte, 100),
		quitch:     make(chan struct{}),
	}
}

func (t *UDPTransport) Listen() error {
	return nil
}

func (t *UDPTransport) Close() error {
	return nil
}

func (t *UDPTransport) ListenAddr() string {
	return t.listenAddr
}

func (t *UDPTransport) Messages() <-chan []byte {
	return t.messagech
}
