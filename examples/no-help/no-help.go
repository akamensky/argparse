package main

import (
	"fmt"

	"github.com/akamensky/argparse"
)

func main() {
	// Create Parser Settings
	settings := argparse.Settings{HelpDisabled: true}
	// Create new parser object
	parser := argparse.NewParserWithSettings("help", "Demonstrates disabing the help arguments", settings)
	// Create string flag
	parser.String("s", "string", &argparse.Options{Required: false, Help: "String argument example"})
	// Create string flag
	parser.Int("i", "int", &argparse.Options{Required: false, Help: "Integer argument example"})

	// parsing for -h fails
	fmt.Println(parser.Parse([]string{"parser", "-h", "--help", "-s", "testing", "-i", "5"}))
}
