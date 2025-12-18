package transport

import "github.com/Bladedancerr/server/utils"

// transport
type Transport interface {
	Listen() error
	Close() error
	WriteLoop()
	Requests() <-chan utils.Request
}
