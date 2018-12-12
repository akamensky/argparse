package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	// Create new parser object
	parser := argparse.NewParser("default", "Program to test default values")

	// Creating string argument with default value and required set to false
	s := parser.String("s", "string", &argparse.Options{Required: false, Help: "String to print", Default: "Hello"})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	// Finally print the collected string
	fmt.Println(*s)
}
