package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/commondatageek/mule/cmd/mule/server/pipemap"
)

func Serve(port int) {
	http.HandleFunc("/", NewHandler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func NewHandler() func(http.ResponseWriter, *http.Request) {
	pm := pipemap.New()
	return func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		key := strings.ToLower(r.URL.Path)
		log.Printf("Received request: %s %s\n", method, key)

		switch method {

		case http.MethodPut:
			rdr, wtr := io.Pipe()
			defer wtr.Close()

			if err := pm.Create(key, rdr); err != nil {
				defer rdr.Close()
				http.Error(w, err.Error(), http.StatusConflict)
			} else {
				io.Copy(wtr, r.Body)
			}

		case http.MethodGet:
			if rdr, ok := pm.Get(key); ok {
				defer rdr.Close()
				io.Copy(w, rdr)
				pm.Delete(key)
			} else {
				http.NotFound(w, r)
			}
		}
	}
}
