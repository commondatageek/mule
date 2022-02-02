package main

import (
	"io"
	"log"
	"net"
)

const Protocol = "tcp"
const ProducerAddress = "localhost:8881"
const ConsumerAddress = "localhost:8882"

func ListenProducer(listenAddr string, out chan<- net.Conn) {
	lsn, err := net.Listen(Protocol, listenAddr)
	if err != nil {
		log.Fatalf("Could not listen for producer on %s: %s\n", listenAddr, err)
	}
	defer lsn.Close()

	conn, err := lsn.Accept()
	log.Println("Received producer")
	if err != nil {
		log.Fatalf("Could not accept connections for producer on %s: %s\n", listenAddr, err)
	}

	out <- conn
	close(out)
}

func ListenConsumer(listenAddr string, out chan<- net.Conn) {
	lsn, err := net.Listen(Protocol, listenAddr)
	if err != nil {
		log.Fatalf("Could not listen for consumer on %s: %s\n", listenAddr, err)
	}
	defer lsn.Close()

	conn, err := lsn.Accept()
	if err != nil {
		log.Fatalf("Could not accept connections for consumer on %s: %s\n", listenAddr, err)
	}

	out <- conn
	close(out)
}

func Transfer(pChan, cChan <-chan net.Conn) (int64, error) {
	pConn := <-pChan
	defer pConn.Close()

	cConn := <-cChan
	defer cConn.Close()

	written, err := io.Copy(cConn, pConn)
	return written, err
}

func main() {
	log.Printf("Listening for producer on %s ...", ProducerAddress)
	log.Printf("Listening for consumer on %s ...", ConsumerAddress)

	// listen for producer
	pChan := make(chan net.Conn)
	go ListenProducer(ProducerAddress, pChan)

	// listen for consumer
	cChan := make(chan net.Conn)
	go ListenConsumer(ConsumerAddress, cChan)

	// make the transfer
	written, err := Transfer(pChan, cChan)
	if err != nil {
		log.Fatalf("Could not transfer data: %s\n", err)
	}

	log.Printf("Transferred %d bytes", written)
}
