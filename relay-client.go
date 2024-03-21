package main

import (
	"io"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("Failed to connect to port 8000: %v", err)
	}
	defer conn.Close()
	log.Println("Connected to port 8000")

	go func() {
		buf := make([]byte, 1024)
		for {
			// set a short read dealline to check for available data
			conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))

			// Attempt to read data from cConn
			n, err := conn.Read(buf)
			if err != nil {
				// If the error is a timeout, it means no data is available
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					continue
				}
				log.Printf("Failed to read from C: %v", err)
				return
			}
			log.Printf("Received from C: %s", buf[:n])

			// connect to d
			cConn, err := net.Dial("tcp", "localhost:3001")
			if err != nil {
				log.Printf("Failed to connect to D: %v", err)
				return
			}

			cConn.Write(buf[:n])
			go io.Copy(conn, cConn)
			io.Copy(cConn, conn)
			cConn.Close()
		}
	}()
	wg.Add(1)
	wg.Wait()
}
