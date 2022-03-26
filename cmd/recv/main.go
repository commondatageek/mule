package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const DefaultHost = "localhost"
const DefaultPort = "8080"

func main() {
	logger := log.New(os.Stderr, "", 0)

	host := getHost()
	port := getPort()
	key := getKey()

	resp, err := http.Get(fmt.Sprintf("http://%s:%s/%s", host, port, key))
	if err != nil {
		logger.Fatalf("Could not get the file: %s\n", err)
	}
	defer resp.Body.Close()

	hash := sha256.New()
	output := io.TeeReader(resp.Body, hash)
	_, err = io.Copy(os.Stdout, output)
	if err != nil {
		logger.Fatalf("Could not copy data: %s\n", err)
	} else {
		logger.Printf("SHA256: %x\n", hash.Sum(nil))
	}
}

func getHost() string {
	if envHost, exists := os.LookupEnv("MULE_HOST"); exists {
		return envHost
	} else {
		return DefaultHost
	}
}

func getPort() string {
	if envPort, exists := os.LookupEnv("MULE_PORT"); exists {
		return envPort
	} else {
		return DefaultPort
	}
}

func getKey() string {
	return os.Args[1]
}
