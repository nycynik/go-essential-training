package main

import (
	"fmt"
	"log"
	"os"
	// "github.com/pkg/errors"
)

// some config type, like an applicationConfig
type Config struct {
	// configurations fields go here
	version float32
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err // errors.Wrap(err, "can't open configuration file")
	}

	defer file.Close()

	cfg := &Config{}
	// parse the config here...

	return cfg, nil
}

func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}
func main() {

	setupLogging()

	cfg, err := readConfig("../config.toml")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		log.Printf("error: %+v\n", err)
		os.Exit(1)
	}
	fmt.Println(cfg)
}