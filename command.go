package argparse

import "reflect"

type command struct {
	name string
	description string
	args map[string]arg
	commands []command
}

func (o *command) Flag(sname string, lname string, opts *Options) *bool {
	var result *bool
	
	a := arg{
		resultPointerType: reflect.TypeOf(result),
		result: result,
		sname: sname,
		lname: lname,
		size: 1,
		opts: opts,
	}
	
	if sname != "" {
		o.args["-" + sname] = a
	}
	
	if lname != "" {
		o.args["--" + lname] = a
	}
	
	return result
}
