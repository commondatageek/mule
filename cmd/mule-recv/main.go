package main

import (
	"io"
	"log"
	"net"
	"os"
)

const Network = "tcp"
const Address = "localhost:8882"

func main() {
	// get connection
	conn, err := net.Dial(Network, Address)
	if err != nil {
		log.Fatalf("Could not connect to server: %s\n", err)
	}
	defer conn.Close()

	recvFile := os.Args[1]

	f, err := os.Create(recvFile)
	if err != nil {
		log.Fatalf("Could not open %s for writing: %s\n", recvFile, err)
	}
	defer f.Close()

	n, err := io.Copy(f, conn)
	if err != nil {
		log.Fatalf("Could not read from socket: %s\n", err)
	}

	log.Printf("Consumer read %d bytes from server\n", n)
}
