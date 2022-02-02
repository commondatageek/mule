package main

import (
	"io"
	"log"
	"net"
	"os"
)

const Network = "tcp"
const Address = "localhost:8881"

func main() {
	// get connection
	conn, err := net.Dial(Network, Address)
	if err != nil {
		log.Fatalf("Could not connect to server: %s\n", err)
	}
	defer conn.Close()

	sendFile := os.Args[1]

	f, err := os.Open(sendFile)
	if err != nil {
		log.Fatalf("Could not open %s: %s\n", sendFile, err)
	}
	defer f.Close()

	n, err := io.Copy(conn, f)
	if err != nil {
		log.Fatalf("Could not write to socket: %s\n", err)
	}

	log.Printf("Producer sent %d bytes to server\n", n)
}
