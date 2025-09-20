package main

import (
	"log"

	"github.com/erniebrodeur/godu/internal/cli"
	"github.com/erniebrodeur/godu/internal/scanner"
)

func main() {
	config := cli.ParseArgs()

	if err := config.Validate(); err != nil {
		log.Fatal(err)
	}

	err := scanner.Scan(config.Directory, config.Verbose, config.Human, config.Depth)
	if err != nil {
		log.Fatal(err)
	}
}
