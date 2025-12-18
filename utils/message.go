package utils

import "net"

type Message struct {
	Addr net.Addr
	Data []byte
}
