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

	client := http.Client{}

	hash := sha256.New()
	input := io.TeeReader(os.Stdin, hash)

	// construct the PUT request
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:%s/%s", Port, key), input)
	if err != nil {
		log.Fatalf("Could not create request: %s\n", err)
	}
	req.ContentLength = -1
	req.Header.Set("Content-Type", "application/octet-stream")

	// send the PUT request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error while executing PUT operation: %s\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("%s\n", resp.Status)
	} else {
		log.Printf("SHA256: %x\n", hash.Sum(nil))
	}
}
