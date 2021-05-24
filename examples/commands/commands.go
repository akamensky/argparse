package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

// Run this as `go run commands.go [start|stop]`
func main() {
	// Create new parser object
	parser := argparse.NewParser("commands", "Simple example of argparse commands")

	// Add top level command `start`
	startCmd := parser.NewCommand("start", "Will start a process")

	// Add top level commands `stop`
	stopCmd := parser.NewCommand("stop", "Will stop a process")
	stopCmd.Int("-t", "--time", nil)

	// Parse command line arguments and in case of any error print error and help information
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	// Check if `start` command was given
	if startCmd.Happened() {
		// Starting a process
		fmt.Println("Started process")
	} else if stopCmd.Happened() { // Check if `stop` command was given
		// Stopping a process
		for _, arg := range stopCmd.GetArgs() {
			switch val := arg.GetResult().(type) {
			case int:
				fmt.Printf("Arg: %s = %d\n", arg.GetLname(), val)
			}
		}
		fmt.Println("Stopped process")
	} else {
		// In fact we will never hit this one
		// because commands and sub-commands are considered as required
		err := fmt.Errorf("bad arguments, please check usage")
		fmt.Print(parser.Usage(err))
	}
}
