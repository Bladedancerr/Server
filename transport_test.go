package main

import "testing"

func TestNewServer(t *testing.T) {
	opts := ServerOpts{
		ListenAddr: "localhost:3000",
	}
	s := NewTCPServer(opts)
	s.Start()
}
