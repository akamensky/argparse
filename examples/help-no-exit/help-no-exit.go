package main

import (
	"fmt"

	"github.com/akamensky/argparse"
)

func main() {
	// Create Parser Settings
	settings := argparse.Settings{NoExitOnHelp: true}
	// Create new parser object
	parser := argparse.NewParserWithSettings("help", "Demonstrates changing the help argument names", settings)
	// Create string flag
	parser.String("s", "string", &argparse.Options{Required: false, Help: "String argument example"})
	// Create string flag
	parser.Int("i", "int", &argparse.Options{Required: false, Help: "Integer argument example"})
	// Use the help function
	fmt.Println(parser.Parse([]string{"parser", "-h"}))
	fmt.Println("Didn't exit, still printing")
}
