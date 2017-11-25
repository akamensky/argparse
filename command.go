package argparse

import (
	"errors"
	"fmt"
)

type command struct {
	name        string
	description string
	args        []arg
	commands    []*command
	parsed      bool
	parent      *command
}

type parser struct {
	command
}

func (o *command) help() {
	result := &help{}

	a := arg{
		result: &result,
		sname:  "h",
		lname:  "help",
		size:   1,
		opts:   &Options{Help:"Print help information"},
		unique: true,
		parent: o,
	}

	o.addArg(a)
}

func (o *command) addArg(a arg) {
	if a.lname != "" {
		if a.sname == "" || len(a.sname) == 1 {
			// Search parents for overlapping commands and fail silently if any
			current := o
			for current != nil {
				if current.args != nil {
					for _, v := range current.args {
						if (a.sname != "" && a.sname == v.sname) || a.lname == v.lname {
							return
						}
					}
				}
				current = current.parent
			}
			o.args = append(o.args, a)
		}
	}
}

// Will parse provided list of arguments
// common usage would be to pass directly os.Args
func (o *command) parse(args *[]string) error {
	// If we already been parsed do nothing
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

	// Reduce arguments by removing command name
	*args = (*args)[1:]

	// Parse subcommands if any
	if o.commands != nil && len(o.commands) > 0 {
		// If we have subcommands and 0 args left
		// that is an error
		if len(*args) < 1 {
			return errors.New("[sub]command required")
		}
		for _, v := range o.commands {
			err := v.parse(args)
			if err != nil {
				return err
			}
		}
	}

	// Iterate over the args
	for i := 0; i < len(o.args); i++ {
		oarg := o.args[i]
		for j := 0; j < len(*args); j++ {
			arg := (*args)[j]
			if arg == "" {
				continue
			}
			if oarg.check(arg) {
				err := oarg.parse((*args)[j+1 : j+oarg.size])
				if err != nil {
					return err
				}
				//*args = append((*args)[:j], (*args)[j+oarg.size:]...)
				oarg.reduce(j, args)
				//j-- // Bump down j to account for reduced args size
				continue
			}
		}

		// Check if arg is required and not provided
		if oarg.opts != nil && oarg.opts.Required && !oarg.parsed {
			return errors.New(fmt.Sprintf("[%s] is required", oarg.name()))
		}
	}

	// Set parsed status to true and return quietly
	o.parsed = true
	return nil
}
