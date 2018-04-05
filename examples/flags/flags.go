package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func main() {
	// Create new parser object
	parser := argparse.NewParser("flags", "Simple example of argparse flags")

	// Create verbose flag
	verb := parser.Flag("v", "verbose", &argparse.Options{Help: "Enable verbose mode"})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	// if verbose mode enabled, print debugging information
	if *verb {
		fmt.Println("Args: ", os.Args)
		fmt.Println("PID: ", os.Getpid())
	}

	// Get and print the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The current working directory is", cwd)
	}
}
