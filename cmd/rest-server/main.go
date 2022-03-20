package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const Path = "test_data"

func handler(w http.ResponseWriter, r *http.Request) {
	src, err := os.Open(Path)
	if err != nil {
		log.Fatalf("Could not open file %s: %s\n", Path, err)
	}
	defer src.Close()

	io.Copy(w, src)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
