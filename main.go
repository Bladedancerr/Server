package main

import (
	"fmt"

	"github.com/Bladedancerr/server/server"
)

func main() {
	opts := server.ServerOpts{
		ListenAddr: "localhost:3000",
	}
	s := server.NewTCPServer(opts)
	if err := s.Start(); err != nil {
		fmt.Println(err)
	}

	// port := ":3000"
	// conn, err := net.ListenPacket("udp", port)
	// if err != nil {
	// 	panic(err)
	// }
	// defer conn.Close()
	// fmt.Println("udp server started on port: ", 3000)
	// buf := make([]byte, 1024)
	// for {
	// 	n, addr, err := conn.ReadFrom(buf)
	// 	if err != nil {
	// 		fmt.Println("Error reading:", err)
	// 		continue
	// 	}
	// 	fmt.Printf("Received %d bytes from %s: %s\n", n, addr, string(buf[:n]))
	// }
}
