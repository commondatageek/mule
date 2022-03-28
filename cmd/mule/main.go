package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/commondatageek/mule/cmd/mule/server"
)

func main() {
	serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)
	servePort := serveCmd.Int("port", LookupEnvOrInt("MULE_PORT", 8080), "the port on which to listen")

	if len(os.Args) < 2 {
		log.Fatal("mule expects a subcommand: {serve | send | recv}")
	}

	cmd := os.Args[1]

	switch cmd {
	case "serve":
		serveCmd.Parse(os.Args[2:])
		server.Serve(*servePort)
	}
}

func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func LookupEnvOrInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("LookupEnvOrInt[%s]: %v", key, err)
		}
		return v
	}
	return defaultVal
}
