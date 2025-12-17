package main

import (
	"net"
	"reflect"
	"sync"
	"testing"
	"time"
)

// TestTCPServer_FullLifecycle provides an end-to-end integration test.
// It verifies that the TCPServer can be started, accept connections,
// echo messages back to a client, and be stopped cleanly.
func TestTCPServer_FullLifecycle(t *testing.T) {
	// 1. Setup the server with a specific address for testing.
	opts := ServerOpts{
		ListenAddr: "127.0.0.1:54321",
	}
	server := NewTCPServer(opts)

	// 2. Start the server in a goroutine because its Start() method is blocking.
	// We use a WaitGroup to wait for the server to shut down completely.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// We expect Start() to eventually return nil when the server is stopped.
		// If it returns an error, it will be caught by other parts of the test.
		server.Start()
	}()

	// Allow a brief moment for the OS to start the TCP listener.
	time.Sleep(50 * time.Millisecond)

	// 3. Connect to the server as a client.
	conn, err := net.Dial("tcp", opts.ListenAddr)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// 4. Send a message and verify the server echoes it back correctly.
	msg := []byte("hello")
	_, err = conn.Write(msg)
	if err != nil {
		t.Fatalf("Failed to write to server: %v", err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatalf("Failed to read from server: %v", err)
	}

	if string(buf[:n]) != string(msg) {
		t.Errorf("Expected echo '%s', got '%s'", string(msg), string(buf[:n]))
	}

	// 5. Stop the server.
	if err := server.Stop(); err != nil {
		t.Errorf("Failed to stop server: %v", err)
	}

	// 6. Wait for the server's main loop to exit.
	wg.Wait()

	// 7. Verify the server is shut down by attempting to connect again.
	_, err = net.Dial("tcp", opts.ListenAddr)
	if err == nil {
		t.Fatalf("Server did not shut down; a new connection was accepted.")
	}
}

// TestTCPTransport_Lifecycle specifically tests the Transport's ability
// to start listening and then close cleanly.
func TestTCPTransport_Lifecycle(t *testing.T) {
	// Use a unique address for this test.
	addr := "127.0.0.1:54322"
	transport := NewTCPTransport(addr)

	// The Listen method is blocking, so we run it in a goroutine
	// and use a channel to capture its return error.
	errCh := make(chan error, 1)
	go func() {
		errCh <- transport.Listen()
	}()

	// Wait for the transport to start listening.
	time.Sleep(50 * time.Millisecond)

	// Ensure the transport is accepting connections.
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("Could not connect to transport: %v", err)
	}
	conn.Close()

	// Close the transport, which should cause the Listen method to return.
	if err := transport.Close(); err != nil {
		t.Fatalf("Failed to close transport: %v", err)
	}

	// Verify that the Listen method returns without an error.
	select {
	case err := <-errCh:
		if err != nil {
			t.Errorf("Listen returned an unexpected error on close: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("Listen did not return after Close was called.")
	}
}

// TestTCPTransport_MessagePassing verifies that a message sent to the transport
// is correctly received and pushed into the public Messages channel.
func TestTCPTransport_MessagePassing(t *testing.T) {
	addr := "127.0.0.1:54323"
	transport := NewTCPTransport(addr)

	go func() {
		// We can ignore the error here as we control the lifecycle.
		_ = transport.Listen()
	}()
	defer transport.Close()

	time.Sleep(50 * time.Millisecond)

	// Connect and send a payload.
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	payload := []byte("test payload")
	_, err = conn.Write(payload)
	if err != nil {
		t.Fatalf("Could not write to connection: %v", err)
	}

	// Verify the transport forwards the message to its channel.
	select {
	case msg := <-transport.Messages():
		if !reflect.DeepEqual(msg, payload) {
			t.Errorf("Expected message %v, got %v", payload, msg)
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatal("Did not receive message from transport channel in time.")
	}
}