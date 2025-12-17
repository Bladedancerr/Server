package main

// transport
type Transport interface {
	Listen() error
	Close() error
	Messages() <-chan []byte
}
