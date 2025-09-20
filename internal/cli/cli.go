package cli

import (
	"flag"
	"fmt"
	"os"
)

const Version = "0.1.0"

type Config struct {
	Directory   string
	Verbose     bool
	Human       bool
	Depth       int
	ShowVersion bool
}

func ParseArgs() *Config {
	config := &Config{}

	flag.BoolVar(&config.Verbose, "v", false, "verbose output")
	flag.BoolVar(&config.Verbose, "verbose", false, "verbose output")

	flag.BoolVar(&config.Human, "h", false, "human readable sizes")
	flag.BoolVar(&config.Human, "human", false, "human readable sizes")

	flag.IntVar(&config.Depth, "d", -1, "maximum depth to display (-1 for unlimited)")
	flag.IntVar(&config.Depth, "depth", -1, "maximum depth to display (-1 for unlimited)")

	flag.BoolVar(&config.ShowVersion, "version", false, "show version information")

	flag.Parse()

	if config.ShowVersion {
		fmt.Printf("godu %s\n", Version)
		os.Exit(0)
	}

	config.Directory = "."
	if flag.NArg() > 0 {
		config.Directory = flag.Arg(0)
	}

	return config
}

func (c *Config) Validate() error {
	// Check if directory exists
	if _, err := os.Stat(c.Directory); os.IsNotExist(err) {
		return err
	}
	return nil
}
