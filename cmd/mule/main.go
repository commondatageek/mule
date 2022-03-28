package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/commondatageek/mule/cmd/mule/client"
	"github.com/commondatageek/mule/cmd/mule/server"
)

const DefaultKeySize = 3 // bytes, or *2 hex digits in string form
const DefaultHost = "localhost"
const DefaultPort = 8080

func main() {
	// "mule serve"
	serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)
	servePort := serveCmd.Int("port", LookupEnvOrInt("MULE_PORT", DefaultPort), "the port on which to listen")

	// "mule send"
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	sendHost := sendCmd.String("host", LookupEnvOrString("MULE_HOST", DefaultHost), "the IP or DNS of a host with a running mule server")
	sendPort := sendCmd.Int("port", LookupEnvOrInt("MULE_PORT", DefaultPort), "the port to send data to")
	sendKey := sendCmd.String("key", LookupEnvOrString("MULE_KEY", client.GenerateRandomKey(DefaultKeySize)), "the key to designate the send")

	// "mule recv"
	recvCmd := flag.NewFlagSet("receive", flag.ExitOnError)
	recvHost := recvCmd.String("host", LookupEnvOrString("MULE_HOST", DefaultHost), "the IP or DNS of a host with a running mule server")
	recvPort := recvCmd.Int("port", LookupEnvOrInt("MULE_PORT", DefaultPort), "the port to receive data from")

	if len(os.Args) < 2 {
		log.Fatal("mule expects a subcommand: {serve | send | receive}")
	}

	cmd := os.Args[1]

	switch cmd {
	case "serve":
		serveCmd.Parse(os.Args[2:])
		server.Serve(*servePort)
	case "send":
		sendCmd.Parse(os.Args[2:])
		client.SendStream(*sendHost, *sendPort, *sendKey, os.Stdin)
	case "receive":
		recvCmd.Parse(os.Args[2:])
		recvKey := recvCmd.Arg(0)
		if recvKey != "" {
			client.ReceiveStream(*recvHost, *recvPort, recvKey)
		} else {
			log.Fatal("Must specify key: mule receive {key}")
		}
	default:
		log.Fatal("subcommands are: {serve | send | receive}")
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
