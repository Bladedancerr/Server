package transport

import (
	"fmt"
	"io"
	"net"

	"github.com/Bladedancerr/server/utils"
)

// tcp transport
// implements transport interface
type TCPTransport struct {
	listenAddr string
	requestch  chan utils.Request
	quitch     chan struct{}
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddr: listenAddr,
		requestch:  make(chan utils.Request, 100),
		quitch:     make(chan struct{}),
	}
}

func (t *TCPTransport) Listen() error {
	ln, err := net.Listen("tcp", t.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	go t.WriteLoop()
	go t.acceptLoop(ln)
	<-t.quitch

	close(t.requestch)
	return nil
}

func (t *TCPTransport) acceptLoop(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-t.quitch:
				return
			default:
				fmt.Println(err)
				continue
			}
		}

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("new connection: ", conn.RemoteAddr())
	reader := utils.NewTCPReader(conn)
	writer := utils.NewMultiWriter(utils.NewConsoleWriter(), utils.NewTCPEchoWriter(conn))

	for {
		msg, err := reader.Read()
		if err == io.EOF {
			fmt.Println("connection dropped: ", conn.RemoteAddr())
			return
		}
		if err != nil {
			fmt.Println("read error: ", err)
			return
		}

		req := utils.NewRequest(*msg, writer)

		select {
		case t.requestch <- *req:
		case <-t.quitch:
			return
		}
	}
}

func (t *TCPTransport) WriteLoop() {
	for req := range t.requestch {
		if _, err := req.Writer.Write(req.Message); err != nil {
			fmt.Println("write error: ", err)
			return
		}
	}
}

func (t *TCPTransport) Close() error {
	close(t.quitch)
	return nil
}

func (t *TCPTransport) ListenAddr() string {
	return t.listenAddr
}

func (t *TCPTransport) Requests() <-chan utils.Request {
	return t.requestch
}
