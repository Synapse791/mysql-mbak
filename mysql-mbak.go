package main

import (
	"flag"
)

var verbose bool

func init() {
	flag.BoolVar(&verbose, "v", false, "enable verbose logging")
	flag.BoolVar(&verbose, "verbose", false, "enable verbose logging")
}

func main() {
	flag.Parse()
}
