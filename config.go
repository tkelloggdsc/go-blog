package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config - represents env vars
type Config struct {
	Server     string
	Database   string
	ServerPort string
}

// Read and parse the config file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
