package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func main() {
	// Create new parser object
	parser := argparse.NewParser("print", "Prints provided string to stdout")
	// Create string flag
	s := parser.String("s", "string", &argparse.Options{Required: true, Help: "String to print"})
	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	// Finally print the collected string
	fmt.Println(*s)
}
