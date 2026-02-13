package transport

import (
	"fmt"
	"net"

	"github.com/Bladedancerr/server/utils"
)

// udp transport
// implements transport interface
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
	conn, err := net.ListenPacket("udp", t.listenAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	go t.readLoop(conn)
	go t.WriteLoop()
	<-t.quitch

	close(t.requestch)
	return nil
}

func (t *UDPTransport) readLoop(conn net.PacketConn) {
	reader := utils.NewUDPReader(conn)
	writer := utils.NewMultiWriter(utils.NewConsoleWriter(), utils.NewUDPEchoWriter(conn))

	for {
		msg, err := reader.Read()
		if err != nil {
			select {
			case <-t.quitch:
				return
			default:
				fmt.Println("udp read error:", err)
				continue
			}
		}

		req := utils.NewRequest(*msg, writer)
		select {
		case t.requestch <- *req:
		case <-t.quitch:
			return
		}
	}
}

func (t *UDPTransport) Close() error {
	close(t.quitch)
	return nil
}

func (t *UDPTransport) WriteLoop() {
	for req := range t.requestch {
		if _, err := req.Writer.Write(req.Message); err != nil {
			fmt.Println("write error: ", err)
			return
		}
	}
}

func (t *UDPTransport) ListenAddr() string {
	return t.listenAddr
}

func (t *UDPTransport) Requests() <-chan utils.Request {
	return t.requestch
}
