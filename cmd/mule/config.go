package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/commondatageek/mule/cmd/mule/client"
)

const DefaultKeySize = 3 // bytes, or *2 hex digits in string form
const DefaultHost = "localhost"
const DefaultPort = 8080
const DotfileName = ".mule"

type Configuration struct {
	Command  *string
	Host     *string
	Port     *int
	Key      *string
	FilePath *string
}

func Configure(dotfilePath string, cmd string) *Configuration {
	cfg := DefaultConfiguration(cmd)

	err := ConfigureFromDotfile(dotfilePath, cmd, cfg)
	if err != nil {
		log.Fatalf("Error loading configurattion from dotfile: %s\n", err)
	}

	err = ConfigureFromEnvVars(cmd, cfg)
	if err != nil {
		log.Fatalf("Error loading configuration from env vars: %s\n", err)
	}

	ConfigureFromFlags(cmd, cfg)

	return cfg
}

func DefaultConfiguration(cmd string) *Configuration {
	cfg := Configuration{
		Command: &cmd,
	}

	switch cmd {
	case "serve":
		var port int = DefaultPort
		cfg.Port = &port

	case "send":
		var host string = DefaultHost
		cfg.Host = &host

		var port int = DefaultPort
		cfg.Port = &port

		var key string = client.GenerateRandomKey(DefaultKeySize)
		cfg.Key = &key

	case "receive":
		var host string = DefaultHost
		cfg.Host = &host

		var port int = DefaultPort
		cfg.Port = &port
	}

	return &cfg
}

type FileConfiguration struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Key  string `json:"key"`
}

func ConfigureFromDotfile(path string, cmd string, cfg *Configuration) error {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("could not open config file at %s: %s", path, err)
	}

	fc := &FileConfiguration{}
	dec := json.NewDecoder(f)
	err = dec.Decode(fc)
	if err != nil {
		return fmt.Errorf("could not decode JSON: %s", err)
	}

	switch cmd {
	case "serve":
		if fc.Port != 0 {
			cfg.Port = &fc.Port
		}

	case "send":
		if fc.Host != "" {
			cfg.Host = &fc.Host
		}
		if fc.Port != 0 {
			cfg.Port = &fc.Port
		}
		if fc.Key != "" {
			cfg.Key = &fc.Key
		}

	case "receive":
		if fc.Host != "" {
			cfg.Host = &fc.Host
		}
		if fc.Port != 0 {
			cfg.Port = &fc.Port
		}
	}
	return nil
}

func ConfigureFromEnvVars(cmd string, cfg *Configuration) error {
	switch cmd {
	case "serve":
		if val, exists := os.LookupEnv("MULE_PORT"); exists {
			v, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("could not convert to int: MULE_PORT=%s", val)
			}
			cfg.Port = &v
		}

	case "send":
		if val, exists := os.LookupEnv("MULE_HOST"); exists {
			cfg.Host = &val
		}
		if val, exists := os.LookupEnv("MULE_PORT"); exists {
			v, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("could not convert to int: MULE_PORT=%s", val)
			}
			cfg.Port = &v
		}
		if val, exists := os.LookupEnv("MULE_KEY"); exists {
			cfg.Key = &val
		}

	case "receive":
		if val, exists := os.LookupEnv("MULE_HOST"); exists {
			cfg.Host = &val
		}
		if val, exists := os.LookupEnv("MULE_PORT"); exists {
			v, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("could not convert to int: MULE_PORT=%s", val)
			}
			cfg.Port = &v
		}
	}

	return nil
}

func ConfigureFromFlags(cmd string, cfg *Configuration) {
	switch cmd {
	case "serve":
		serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)
		serveCmd.IntVar(cfg.Port, "port", *cfg.Port, "the port on which to listen")
		serveCmd.Parse(os.Args[2:])

	case "send":
		sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
		sendCmd.StringVar(cfg.Host, "host", *cfg.Host, "the IP or DNS of a host with a running mule server")
		sendCmd.IntVar(cfg.Port, "port", *cfg.Port, "the port to send data to")
		sendCmd.StringVar(cfg.Key, "key", *cfg.Key, "the key to designate the send")
		sendCmd.Parse(os.Args[2:])

		if filePath := sendCmd.Arg(0); filePath != "" {
			cfg.FilePath = &filePath
		}

	case "receive":
		recvCmd := flag.NewFlagSet("receive", flag.ExitOnError)
		recvCmd.StringVar(cfg.Host, "host", *cfg.Host, "the IP or DNS of a host with a running mule server")
		recvCmd.IntVar(cfg.Port, "port", *cfg.Port, "the port to receive data from")
		recvCmd.Parse(os.Args[2:])

		if key := recvCmd.Arg(0); key != "" {
			cfg.Key = &key
		}
	}
}
