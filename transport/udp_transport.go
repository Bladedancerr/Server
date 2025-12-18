package transport

import "github.com/Bladedancerr/server/utils"

// udp transport
// must implement transport interface
type UDPTransport struct {
	listenAddr string
	requestch  chan utils.Request
	quitch     chan struct{}
}

func NewUDPTransport(listenAddr string) *UDPTransport {
	return &UDPTransport{
		listenAddr: listenAddr,
		requestch:  make(chan utils.Request, 100),
		quitch:     make(chan struct{}),
	}
}

func (t *UDPTransport) Listen() error {
	return nil
}

func (t *UDPTransport) Close() error {
	return nil
}

func (t *UDPTransport) WriteLoop() {

}

func (t *UDPTransport) ListenAddr() string {
	return t.listenAddr
}

func (t *UDPTransport) Requests() <-chan utils.Request {
	return t.requestch
}
