package argparse

import "reflect"

type command struct {
	name        string
	description string
	args        map[string]arg
	commands    []command
}

func (o *command) Flag(sname string, lname string, opts *Options) *bool {
	var result bool

	a := arg{
		resultPointerType: reflect.TypeOf(result),
		result:            result,
		sname:             sname,
		lname:             lname,
		size:              1,
		opts:              opts,
	}

	if sname != "" {
		o.args["-"+sname] = a
	}

	if lname != "" {
		o.args["--"+lname] = a
	}

	return &result
}

// Will parse provided list of arguments
// common usage would be to pass directly os.Args
func (o *command) Parse(args []string) {
	// If user did not set progname in advance, set progname from os.Args
	if o.name == "" {
		o.name = args[0]
	}

	args = args[1:]

	// TODO: Implement sub-commands parsing

	// FIXME: Below is wrong, must iterate over reduced argument list o.args, not over args ([]string)
	// Iterate over the rest of args
	for i := 0; i < len(args); {
		v := args[i]
		if arg, ok := o.args[v]; ok {
			arg.parse(args[i : i+1])
			i = i + arg.size
		} else {
			i++
		}
	}
}
