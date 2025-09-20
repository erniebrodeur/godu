package main

import (
	"flag"
	"log"

	"github.com/erniebrodeur/godu/internal/scanner"
)

func main() {
	verbose := flag.Bool("v", false, "verbose output")
	flag.BoolVar(verbose, "verbose", false, "verbose output")

	human := flag.Bool("h", false, "human readable sizes")
	flag.BoolVar(human, "human", false, "human readable sizes")

	depth := flag.Int("d", -1, "maximum depth to display (-1 for unlimited)")
	flag.IntVar(depth, "depth", -1, "maximum depth to display (-1 for unlimited)")

	flag.Parse()

	dir := "."
	if flag.NArg() > 0 {
		dir = flag.Arg(0)
	}

	err := scanner.Scan(dir, *verbose, *human, *depth)
	if err != nil {
		log.Fatal(err)
	}
}
