package utils

import (
	"fmt"
	"net"
)

type Reader interface {
	Read() (*Message, error)
}

type Writer interface {
	Write(Message) (int, error)
}

// implements reader interface
type TCPReader struct {
	conn net.Conn
	buf  []byte
}

func NewTCPReader(conn net.Conn) *TCPReader {
	return &TCPReader{
		conn: conn,
		buf:  make([]byte, 2048),
	}
}

func (r *TCPReader) Read() (*Message, error) {
	n, err := r.conn.Read(r.buf)
	if err != nil {
		return nil, err
	}
	return &Message{
		Addr: r.conn.RemoteAddr(),
		Data: r.buf[:n],
	}, nil
}

// udp reader
// implements reader interface
type UDPReader struct {
	conn net.PacketConn
	buf  []byte
}

func NewUDPReader(conn net.PacketConn) *UDPReader {
	return &UDPReader{
		conn: conn,
		buf:  make([]byte, 2048),
	}
}

func (r *UDPReader) Read() (*Message, error) {
	n, addr, err := r.conn.ReadFrom(r.buf)
	if err != nil {
		return nil, err
	}
	return &Message{
		Addr: addr,
		Data: r.buf[:n],
	}, nil
}

// implements writer interface
type ConsoleWriter struct{}

func NewConsoleWriter() *ConsoleWriter {
	return &ConsoleWriter{}
}

func (w *ConsoleWriter) Write(msg Message) (int, error) {
	n, err := fmt.Println(string(msg.Data))
	return n, err
}

// implements writer interface
type TCPEchoWriter struct {
	conn net.Conn
}

func NewTCPEchoWriter(conn net.Conn) *TCPEchoWriter {
	return &TCPEchoWriter{
		conn: conn,
	}
}

func (w *TCPEchoWriter) Write(msg Message) (int, error) {
	n, err := w.conn.Write(msg.Data)
	return n, err
}

// implements writer interface
type UDPEchoWriter struct {
	conn net.PacketConn
}

func NewUDPEchoWriter(conn net.PacketConn) *UDPEchoWriter {
	return &UDPEchoWriter{
		conn: conn,
	}
}

func (w *UDPEchoWriter) Write(msg Message) (int, error) {
	if msg.Addr == nil {
		return 0, fmt.Errorf("udp echo requires a destination address in the message")
	}
	n, err := w.conn.WriteTo(msg.Data, msg.Addr)
	return n, err
}

// multiwriter
// implements writer interface
type MultiWriter struct {
	writers []Writer
}

func NewMultiWriter(writers ...Writer) *MultiWriter {
	return &MultiWriter{
		writers: writers,
	}
}

func (w *MultiWriter) Write(msg Message) (int, error) {
	for _, writer := range w.writers {
		if _, err := writer.Write(msg); err != nil {
			return 0, err
		}
	}
	return len(msg.Data), nil
}
