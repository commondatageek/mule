package client

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
)

func SendStream(host string, port int, key string, src io.Reader) {
	client := http.Client{}

	hash := sha256.New()
	teedSrc := io.TeeReader(src, hash)

	// construct the PUT request
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%s:%d/%s", host, port, key), teedSrc)
	if err != nil {
		log.Fatalf("Could not create request: %s\n", err)
	}
	req.ContentLength = -1
	req.Header.Set("Content-Type", "application/octet-stream")

	// send the PUT request
	log.Printf("Receive key: %s\n", key)
	// TODO: copy this key to the clipboard so it's easy to send over slack
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

func GenerateRandomKey(keySize int) string {
	keyBytes := make([]byte, keySize)
	n, err := rand.Read(keyBytes)
	if err != nil {
		log.Fatalf("Could not get a random key: %s\n", err)
	}
	if n != keySize {
		log.Fatalf("Expected %d random bytes, received %d\n", keySize, n)
	}
	return fmt.Sprintf("%x", keyBytes)
}
