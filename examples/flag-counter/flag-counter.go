package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func main() {
	// initialize parser
	parser := argparse.NewParser("FlagCounter", "Example of FlagCounter usage")

	// create FlagCounter argument
	opts := &argparse.Options{
		Required: true,
		Help:     "Will print out how many instances of the flag are found. For example, both -nn and --number --number will be 2",
	}
	count := parser.FlagCounter("n", "number", opts)

	// parse arguments
	err := parser.Parse(os.Args)

	// check for errors in parsing
	if err != nil {
		fmt.Printf("Error parsing: [%+v]\n", err)
		return
	}

	// print out the number of occurrences of the flag
	fmt.Printf("Number of flags detected: [%d]\n", *count)
}
