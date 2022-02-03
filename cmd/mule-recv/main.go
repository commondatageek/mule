package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// command line options
	port := flag.Int("port", 8881, "mule-server consumer port")
	host := flag.String("host", "localhost", "host name or IP for mule-server")
	outFile := flag.String("outfile", "received", "path of output file")
	flag.Parse()

	// get connection
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalf("Could not connect to %s on port %d: %s\n", *host, *port, err)
	}
	defer conn.Close()

	// open up the file to write received contents to
	f, err := os.Create(*outFile)
	if err != nil {
		log.Fatalf("Could not open %s for writing: %s\n", *outFile, err)
	}
	defer f.Close()

	// read until the connection is closed
	n, err := io.Copy(f, conn)
	if err != nil {
		log.Fatalf("Could not read from socket: %s\n", err)
	}

	log.Printf("Consumer read %d bytes from server\n", n)
}
