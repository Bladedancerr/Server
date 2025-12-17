package main

// server
type Server interface {
	Start() error
	Stop() error
}

// options
type ServerOpts struct {
	ListenAddr string
}
