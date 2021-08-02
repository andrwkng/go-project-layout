package main

// Reads configs and starts the app

import (
	"log"
)

type Config struct{}

func main() {
	var cfg Config
	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg Config) error {
	server := Server{}
	return server.ListenAndServe(":"+config.Port, server)
}
