package argparse

import "fmt"

func ExampleCommand_Help() {
	parser := NewParser("parser", "")
	parser.HelpFunc = func(c *Command, msg interface{}) string {
		return fmt.Sprintf("Name: %s\n", c.GetName())
	}
	fmt.Println(parser.Help(nil))

	// Output:
	// Name: parser
}

func ExampleCommand_Help_subcommandDefaulting() {
	parser := NewParser("parser", "")
	parser.HelpFunc = func(c *Command, msg interface{}) string {
		helpString := fmt.Sprintf("Name: %s\n", c.GetName())
		for _, com := range c.GetCommands() {
			// Calls parser.HelpFunc, because command.HelpFuncs are nil
			helpString += com.Help(nil)
		}
		return helpString
	}
	parser.NewCommand("subcommand1", "")
	parser.NewCommand("subcommand2", "")
	fmt.Println(parser.Help(nil))

	// Output:
	// Name: parser
	// Name: subcommand1
	// Name: subcommand2
}
func ExampleCommand_Help_subcommandHelpFuncs() {
	parser := NewParser("parser", "")
	parser.HelpFunc = func(c *Command, msg interface{}) string {
		helpString := fmt.Sprintf("Name: %s\n", c.GetName())
		for _, com := range c.GetCommands() {
			// Calls command.HelpFunc, because command.HelpFuncs are not nil
			helpString += com.Help(nil)
		}
		return helpString
	}
	com1 := parser.NewCommand("subcommand1", "Test description")
	com1.HelpFunc = func(c *Command, msg interface{}) string {
		helpString := fmt.Sprintf("Name: %s, Description: %s\n", c.GetName(), c.GetDescription())
		return helpString
	}
	com2 := parser.NewCommand("subcommand2", "")
	com2.String("s", "string", &Options{Required: false})
	com2.String("i", "integer", &Options{Required: true})
	com2.HelpFunc = func(c *Command, msg interface{}) string {
		helpString := fmt.Sprintf("Name: %s\n", c.GetName())
		for _, arg := range c.GetArgs() {
			helpString += fmt.Sprintf("\tLname: %s, Required: %t\n", arg.GetLname(), arg.GetOpts().Required)
		}
		return helpString
	}
	fmt.Print(parser.Help(nil))
	fmt.Print(com1.Help(nil))
	fmt.Print(com2.Help(nil))

	// Output:
	// Name: parser
	// Name: subcommand1, Description: Test description
	// Name: subcommand2
	//	Lname: string, Required: false
	//	Lname: integer, Required: true
	// Name: subcommand1, Description: Test description
	// Name: subcommand2
	//	Lname: string, Required: false
	//	Lname: integer, Required: true
}
