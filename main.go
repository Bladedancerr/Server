package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Bladedancerr/server/server"
)

func main() {
	StartTCPServer()
	// StartUDPServer()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		tokens := []string{
			"you", "got", "message", "from", "server", "and", "be", "happy", "about", "it", "!",
		}

		for _, token := range tokens {
			c := fmt.Sprintf("data: %s\n\n", token)
			w.Write([]byte(c))
			w.(http.Flusher).Flush()
			time.Sleep(time.Millisecond * 500)
		}
	})

	http.ListenAndServe(":3000", mux)
}

func StartTCPServer() {
	tcpOpts := server.ServerOpts{
		ListenAddr: ":3001",
	}

	s := server.NewTCPServer(tcpOpts)
	if err := s.Start(); err != nil {
		fmt.Println(err)
	}
}

func StartUDPServer() {
	udpOpts := server.ServerOpts{
		ListenAddr: ":3000",
	}

	s := server.NewUDPServer(udpOpts)
	if err := s.Start(); err != nil {
		fmt.Println(err)
	}
}
