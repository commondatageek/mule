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
	logger := log.New(os.Stderr, "", 0)

	client := http.Client{}

	hash := sha256.New()
	input := io.TeeReader(os.Stdin, hash)

	// construct the PUT request
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:%s/%s", Port, key), input)
	if err != nil {
		logger.Fatalf("Could not create request: %s\n", err)
	}
	req.ContentLength = -1
	req.Header.Set("Content-Type", "application/octet-stream")

	// send the PUT request
	logger.Printf("Receive key: %s\n", key)
	// TODO: copy this key to the clipboard so it's easy to send over slack
	resp, err := client.Do(req)
	if err != nil {
		logger.Fatalf("Error while executing PUT operation: %s\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		logger.Fatalf("%s\n", resp.Status)
	} else {
		logger.Printf("SHA256: %x\n", hash.Sum(nil))
	}
}
