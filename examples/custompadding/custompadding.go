package main

import (
	"fmt"

	"github.com/akamensky/argparse"
)

func main() {
	// Create new parser object
	parser := argparse.NewParser("help",
		`Demonstrates multiline description and help messages.
With some very long lines like this one which is broken automatically by argparse.
And some shorter.`)
	// Create string flag
	parser.String("s", "string", &argparse.Options{Required: true, Help: "String argument example\non several lines"})
	// Create string flag
	parser.Int("i", "int", &argparse.Options{Required: true, Help: "Integer argument example"})
	parser.DescriptionPadding = 3
	// Use the help function
	fmt.Print(parser.Help(nil))
}
