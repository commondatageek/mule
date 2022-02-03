package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const Network = "tcp"
const Address = "localhost:8881"

func main() {
	// command line options
	port := flag.Int("port", 8881, "mule-server consumer port")
	host := flag.String("host", "localhost", "host name or IP for mule-server")
	inFile := flag.String("infile", "", "path of input file")
	flag.Parse()

	// get connection
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalf("Could not connect to %s on port %d: %s\n", *host, *port, err)
	}
	defer conn.Close()

	// open up the file to send
	f, err := os.Open(*inFile)
	if err != nil {
		log.Fatalf("Could not open %s: %s\n", *inFile, err)
	}
	defer f.Close()

	// write file to socket
	n, err := io.Copy(conn, f)
	if err != nil {
		log.Fatalf("Could not write to socket: %s\n", err)
	}

	log.Printf("Producer sent %d bytes to server\n", n)
}
