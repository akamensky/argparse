package main

import (
	"fmt"

	"github.com/akamensky/argparse"
)

func main() {
	// Create new parser object
	parser := argparse.NewParser("help", "Demonstrates overriding the default help output")
	// Replace parser.Usage as the help message
	parser.HelpFunc = func(c *argparse.Command, msg interface{}) string {
		var help string
		help += fmt.Sprintf("Name: %s, Description: %s\n", c.GetName(), c.GetDescription())
		for _, arg := range c.GetArgs() {
			if arg.GetOpts() != nil {
				help += fmt.Sprintf("Sname: %s, Lname: %s, Help: %s\n", arg.GetSname(), arg.GetLname(), arg.GetOpts().Help)
			} else {
				help += fmt.Sprintf("Sname: %s, Lname: %s\n", arg.GetSname(), arg.GetLname())
			}
		}
		return help
	}

	// Create command
	command := parser.NewCommand("sub", "Demonstrates sub command help")
	// Create string flag
	command.String("s", "string", &argparse.Options{Required: true, Help: "String argument example"})
	// Create string flag
	command.Int("i", "int", &argparse.Options{Required: true, Help: "Integer argument example"})
	// Without declaring command help function, defaults to first non nil parent help function
	fmt.Print(command.Help(nil))
}
