package main

import (
	"fmt"
	"io"
	"net"
)

// server
type Server interface {
	Start() error
	Stop() error
}

type ServerOpts struct {
	ListenAddr string
}

// tcp server
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

// transport
type Transport interface {
	Listen() error
	Close() error
	Messages() <-chan []byte
}

// tcp transport
type TCPTransport struct {
	listenAddr string
	messagech  chan []byte
	quitch     chan struct{}
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddr: listenAddr,
		messagech:  make(chan []byte, 100),
		quitch:     make(chan struct{}),
	}
}

func (t *TCPTransport) ListenAddr() string {
	return t.listenAddr
}

func (t *TCPTransport) Messages() <-chan []byte {
	return t.messagech
}

func (t *TCPTransport) Listen() error {
	ln, err := net.Listen("tcp", t.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	go t.acceptLoop(ln)
	<-t.quitch

	close(t.messagech)
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

	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("connection dropped: ", conn.RemoteAddr())
			return
		}
		if err != nil {
			fmt.Println("read error: ", err)
			return
		}

		if _, err := conn.Write(buf[:n]); err != nil {
			fmt.Println(err)
		}
		select {
		case t.messagech <- append([]byte(nil), buf[:n]...):
		case <-t.quitch:
			return
		}
	}
}
func (t *TCPTransport) Close() error {
	close(t.quitch)
	return nil
}
