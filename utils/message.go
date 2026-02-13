package utils

import "net"

// message created when something is written in connection
type Message struct {
	Addr net.Addr
	Data []byte
}
