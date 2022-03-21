package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const Port = "8080"

func main() {
	key := os.Args[1]

	resp, err := http.Get(fmt.Sprintf("http://localhost:%s/%s", Port, key))
	if err != nil {
		log.Fatalf("Could not get the file: %s\n", err)
	}
	defer resp.Body.Close()

	hash := sha256.New()
	output := io.TeeReader(resp.Body, hash)
	_, err = io.Copy(os.Stdout, output)
	if err != nil {
		log.Fatalf("Could not copy data: %s\n", err)
	} else {
		log.Printf("SHA256: %x\n", hash.Sum(nil))
	}
}
