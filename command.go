package argparse

import (
	"fmt"
)

func (o *Command) help() {
	result := &help{}

	a := &arg{
		result: result,
		sname:  "h",
		lname:  "help",
		size:   1,
		opts:   &Options{Help: "Print help information"},
		unique: true,
	}

	o.addArg(a)
}

func (o *Command) addArg(a *arg) {
	if a.lname != "" {
		if a.sname == "" || len(a.sname) == 1 {
			// Search parents for overlapping commands and fail silently if any
			sswitch, lswitch := "-"+a.sname, "--"+a.lname
			current := o
			for current != nil {
				_, snameconflict := current.mapargs[sswitch]
				_, lnameconflict := current.mapargs[lswitch]
				if snameconflict || lnameconflict {
					return
				}
				current = current.parent
			}
			a.parent = o
			o.args = append(o.args, a)
			if len(a.sname) != 0 {
				o.mapargs[sswitch] = a
			}
			o.mapargs[lswitch] = a
		}
	}
}

// Will parse provided list of arguments
// common usage would be to pass directly os.Args
func (o *Command) parse(args *[]string) error {
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

	// Reduce arguments by removing Command name
	*args = (*args)[1:]

	// Parse subcommands if any
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
		}
	}

	// Iterate over the input args
	for j := 0; j < len(*args); {
		var oarg *arg
		var match bool

		arg := (*args)[j]
		if arg == "" {
			j++
			continue
		}

		oarg, match = o.mapargs[arg]
		if !match { // couldn't match argument directly
			// match short names without following space
			// long names should always appear as separate argument
			if arg[0] == '-' && arg[1] != '-' {
				if oarg, match = o.mapargs[arg[:2]]; match {
					// is a Flag and there are following characters
					if oarg.size == 1 && len(arg) > 2 {
						// leave '-' behind for next iteration
						(*args)[j] = "-" + arg[2:]
					} else {
						// rest of the characters will be parameters to this arg
						(*args)[j] = arg[2:]
					}
					// dont increment j
					arg = arg[:2]
				}
			}
		} else { // matched argument directly
			// consume the arg name
			(*args)[j] = ""
			j++
		}

		if !match {
			j++
			continue
		}

		if len(*args) < j+oarg.size-1 {
			return fmt.Errorf("not enough arguments for %s", oarg.name())
		}
		// parse that many arguments and skipover j
		err := oarg.parse((*args)[j : j+oarg.size-1])
		if err != nil {
			return err
		}

		// consume whatever is parsed
		removeTill := j + oarg.size - 1
		for ; j < removeTill; j++ {
			(*args)[j] = ""
		}
	}

	// Iterate over known args to check required and assign defaults
	for i := 0; i < len(o.args); i++ {
		oarg := o.args[i]
		// Check if arg is required and not provided
		if oarg.opts != nil && oarg.opts.Required && !oarg.parsed {
			return fmt.Errorf("[%s] is required", oarg.name())
		}

		// Check for argument default value and if provided try to type cast and assign
		if oarg.opts != nil && oarg.opts.Default != nil && !oarg.parsed {
			err := oarg.setDefault()
			if err != nil {
				return err
			}
		}
	}

	// Set parsed status to true and return quietly
	o.parsed = true
	return nil
}
