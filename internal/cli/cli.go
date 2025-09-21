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
	ByType      bool
	ShowVersion bool
}

func ParseArgs() *Config {
	config := &Config{}

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: godu [options] [directory]\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, "  -h, --human     Human readable sizes\n")
		fmt.Fprintf(os.Stderr, "  -d, --depth N   Limit depth to N levels\n")
		fmt.Fprintf(os.Stderr, "  -t, --type      Group by file type\n")
		fmt.Fprintf(os.Stderr, "  -v, --verbose   Verbose output\n")
		fmt.Fprintf(os.Stderr, "  --version       Show version\n")
	}

	flag.BoolVar(&config.Verbose, "v", false, "verbose output")
	flag.BoolVar(&config.Verbose, "verbose", false, "verbose output")

	flag.BoolVar(&config.Human, "h", false, "human readable sizes")
	flag.BoolVar(&config.Human, "human", false, "human readable sizes")

	flag.IntVar(&config.Depth, "d", -1, "maximum depth to display (-1 for unlimited)")
	flag.IntVar(&config.Depth, "depth", -1, "maximum depth to display (-1 for unlimited)")

	flag.BoolVar(&config.ByType, "t", false, "group by file type")
	flag.BoolVar(&config.ByType, "type", false, "group by file type")

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
	if _, err := os.Stat(c.Directory); os.IsNotExist(err) {
		return err
	}
	return nil
}
