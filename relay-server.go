package main

import (
	"io"
	"log"
	"net"
	"sync"
)

func handleProxy(conn net.Conn, cConn net.Conn) {
	defer conn.Close()

	// Proxy between conn and cConn
	go io.Copy(conn, cConn)
	io.Copy(cConn, conn)
}

func main() {
	var wg sync.WaitGroup
	// start listen for relay client
	relayServerListener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen on port 8000: %v", err)
	}
	defer relayServerListener.Close()
	log.Println("Start listening on port 8000 for realy client connection")

	relayConn, err := relayServerListener.Accept()
	if err != nil {
		log.Fatalf("Failed to accept C connection: %v", err)
	}
	wg.Add(1)

	serverListener, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("Failed to listen on port 8001: %v", err)
	}
	defer serverListener.Close()
	log.Println("Start listening on port 8001 for client connections")
	go func() {
		for {
			conn, err := serverListener.Accept()
			if err != nil {
				log.Printf("Failed to accept client connection: %v", err)
				continue
			}
			go handleProxy(conn, relayConn)
		}
	}()

	wg.Wait()
}
