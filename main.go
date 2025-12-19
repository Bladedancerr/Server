package main

import (
	"fmt"

	"github.com/Bladedancerr/server/server"
)

func main() {
	opts := server.ServerOpts{
		ListenAddr: ":3000",
	}
	// s := server.NewTCPServer(opts)
	// if err := s.Start(); err != nil {
	// 	fmt.Println(err)
	// }
	s := server.NewUDPServer(opts)
	if err := s.Start(); err != nil {
		fmt.Println(err)
	}
}
