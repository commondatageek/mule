package main

import (
	"fmt"
	"log"
	"os"

	"github.com/commondatageek/mule/cmd/mule/client"
	"github.com/commondatageek/mule/cmd/mule/server"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("mule expects a subcommand: {serve | send | receive}")
	}

	cmd := os.Args[1]

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Could not find home directory: %s\n", err)
	}

	cfg := Configure(fmt.Sprintf("%s/.mule", homeDir), cmd)

	switch cmd {
	case "serve":
		server.Serve(*cfg.Port)
	case "send":
		if *cfg.FilePath != "" {
			client.SendFile(*cfg.Host, *cfg.Port, *cfg.Key, *cfg.FilePath)
		} else {
			client.SendStream(*cfg.Host, *cfg.Port, *cfg.Key, os.Stdin)
		}
	case "receive":
		if *cfg.Key != "" {
			client.ReceiveStream(*cfg.Host, *cfg.Port, *cfg.Key)
		} else {
			log.Fatal("Must specify key: mule receive {key}")
		}
	default:
		log.Fatal("subcommands are: {serve | send | receive}")
	}
}
