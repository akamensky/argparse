package argparse

import (
	"fmt"
	"strings"
)

func (o *Command) help(sname, lname string) {
	result := &help{}

	if lname == "" {
		sname, lname = "h", "help"
	}

	a := &arg{
		result: result,
		sname:  sname,
		lname:  lname,
		size:   1,
		opts:   &Options{Help: "Print help information"},
		unique: true,
	}

	o.addArg(a)
}

func (o *Command) addArg(a *arg) error {
	// long name should be provided
	if a.lname == "" {
		return fmt.Errorf("long name should be provided")
	}
	// short name could be provided and must not exceed 1 character
	if len(a.sname) > 1 {
		return fmt.Errorf("short name must not exceed 1 character")
	}
	// Search parents for overlapping commands and fail if any
	current := o
	for current != nil {
		if current.args != nil {
			for _, v := range current.args {
				if a.lname != "help" || a.sname != "h" {
					if a.sname != "" && a.sname == v.sname {
						return fmt.Errorf("short name %s occurs more than once", a.sname)
					}
					if a.lname == v.lname {
						return fmt.Errorf("long name %s occurs more than once", a.lname)
					}
				}
			}
		}
		current = current.parent
	}
	a.parent = o

	if a.GetPositional() {
		switch a.argType { // Secondary guard
		case Flag, FlagCounter, StringList, IntList, FloatList, FileList:
			return fmt.Errorf("argument type cannot be positional")
		}
		a.sname = ""
		a.opts.Required = false
		a.size = 1 // We could allow other sizes in the future
	}
	o.args = append(o.args, a)

	return nil
}

//parseSubCommands - Parses subcommands if any
func (o *Command) parseSubCommands(args *[]string) error {
	if o.commands != nil && len(o.commands) > 0 {
		// If we have subcommands and 0 args left
		// that is an error of SubCommandError type
		if len(*args) < 1 {
			return newSubCommandError(o)
		}
		for _, v := range o.commands {
			err := v.parse(args)
			if err != nil {
				return err
			}
			if v.happened {
				return nil
			}
		}
		// If we got here, there were subcommands to parse,
		// but none were found, so return an error
		return newSubCommandError(o)
	}
	return nil
}

// Breadth-first parse style for positionals
// Each command proceeds left to right consuming as many
//     positionals as it needs before beginning sub-command parsing
// All flags must have been parsed and reduced prior to calling this
// Positionals will consume any remaining values,
//     disregarding if they have dashes or equals signs or other "delims".
func (o *Command) parsePositionals(inputArgs *[]string) error {
	for _, oarg := range o.args {
		// Two-stage parsing, this is the second stage
		if !oarg.GetPositional() {
			continue
		}
		for j := 0; j < len(*inputArgs); j++ {
			arg := (*inputArgs)[j]
			if arg == "" {
				continue
			}
			if err := oarg.parsePositional(arg); err != nil {
				return err
			}
			oarg.reduce(j, inputArgs)
			break // Positionals can only occur once
		}
		// positional was unsatisfiable, use the default
		if !oarg.parsed {
			err := oarg.setDefault()
			if err != nil {
				return err
			}
		}
	}
	for _, c := range o.commands {
		if c.happened { // presumption of only one sub-command happening
			return c.parsePositionals(inputArgs)
		}
	}
	return nil
}

//parseArguments - Parses arguments
func (o *Command) parseArguments(inputArgs *[]string) error {
	// Iterate over the args
	for _, oarg := range o.args {
		if oarg.GetPositional() { // Two-stage parsing, this is the first stage
			continue
		}
		for j := 0; j < len(*inputArgs); j++ {
			arg := (*inputArgs)[j]
			if arg == "" {
				continue
			}
			if strings.Contains(arg, "=") {
				splitInd := strings.LastIndex(arg, "=")
				equalArg := []string{arg[:splitInd], arg[splitInd+1:]}
				if cnt, err := oarg.check(equalArg[0]); err != nil {
					return err
				} else if cnt > 0 { // No args implies we supply default
					if equalArg[1] == "" {
						return fmt.Errorf("not enough arguments for %s", oarg.name())
					}
					oarg.eqChar = true
					oarg.size = 1
					currArg := []string{equalArg[1]}
					err := oarg.parse(currArg, cnt)
					if err != nil {
						return err
					}
					oarg.reduce(j, inputArgs)
					continue
				}
			}
			if cnt, err := oarg.check(arg); err != nil {
				return err
			} else if cnt > 0 {
				if len(*inputArgs) < j+oarg.size {
					return fmt.Errorf("not enough arguments for %s", oarg.name())
				}
				err := oarg.parse((*inputArgs)[j+1:j+oarg.size], cnt)
				if err != nil {
					return err
				}
				oarg.reduce(j, inputArgs)
				continue
			}
		}

		// Check if arg is required and not provided
		if oarg.opts != nil && oarg.opts.Required && !oarg.parsed {
			return fmt.Errorf("[%s] is required", oarg.name())
		} else if oarg.opts != nil && oarg.opts.Default != nil && !oarg.parsed {
			// Check for argument default value and if provided try to type cast and assign
			err := oarg.setDefault()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Will parse provided list of arguments
// common usage would be to pass directly os.Args
// Depth-first parsing: We will reach the deepest
//    node of the command tree and then parse arguments,
//    stepping back up only after each node is satisfied.
func (o *Command) parse(args *[]string) error {
	// If already been parsed do nothing
	if o.parsed {
		return nil
	}

	// If no arguments left to parse do nothing
	if len(*args) < 1 {
		return nil
	}

	// Parse only matching commands
	// But we always have to parse top level
	if o.name == "" {
		o.name = (*args)[0]
	} else {
		if o.name != (*args)[0] && o.parent != nil {
			return nil
		}
	}

	// Set happened status to true when command happened
	o.happened = true

	// Reduce arguments by removing Command name
	*args = (*args)[1:]

	// Parse subcommands if any
	if err := o.parseSubCommands(args); err != nil {
		return err
	}

	// Parse arguments if any
	if err := o.parseArguments(args); err != nil {
		return err
	}

	// Set parsed status to true and return quietly
	o.parsed = true
	return nil
}
