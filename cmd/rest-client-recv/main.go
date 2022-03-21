package main

import (
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

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatalf("Could not copy data: %s\n", err)
	}
}
