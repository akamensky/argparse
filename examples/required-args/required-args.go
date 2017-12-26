package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func main() {
	parser := argparse.NewParser("Required-args", "Example of required vs optional arguments")

	// Required shorthand argument
	foo := parser.String("f", "foo", &argparse.Options{Required: true, Help: "foo is a required string option"})

	// Optional shorthand argument
	bar := parser.String("b", "bar", &argparse.Options{Required: false, Help: "bar is not a required option"})

	// Parse args
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	} else {
		fmt.Println("Provided arguments:")
		fmt.Printf("foo (required): %s\n", *foo)

		// As an optional string argument, bar will be empty if unused
		if *bar != "" {
			fmt.Printf("bar (optional): %s\n", *bar)
		}
	}
}
