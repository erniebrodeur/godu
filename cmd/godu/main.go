package main

import (
	"log"

	"github.com/erniebrodeur/godu/internal/analyzer"
	"github.com/erniebrodeur/godu/internal/cli"
	"github.com/erniebrodeur/godu/internal/scanner"
)

const Version = "0.2.0"

func main() {
	config := cli.ParseArgs(Version)

	if err := config.Validate(); err != nil {
		log.Fatal(err)
	}

	if config.ByType {
		err := analyzer.Run(config.Directory, config.Verbose, config.Human)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := scanner.Scan(config.Directory, config.Verbose, config.Human, config.Depth)
		if err != nil {
			log.Fatal(err)
		}
	}
}
