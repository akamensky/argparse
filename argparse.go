package argparse

import (
	"errors"
	"os"
	"strings"
)

func NewParser(name string, description string) *parser {
	p := &parser{}

	p.name = name
	p.description = description

	p.args = make([]*arg, 0)
	p.commands = make([]*command, 0)

	p.help()

	return p
}

func (o *command) NewCommand(name string, description string) *command {
	c := new(command)
	c.name = name
	c.description = description
	c.parsed = false
	c.parent = o

	c.help()

	if o.commands == nil {
		o.commands = make([]*command, 0)
	}

	o.commands = append(o.commands, c)

	return c
}

func (o *command) Flag(sname string, lname string, opts *Options) *bool {
	var result bool

	a := &arg{
		result: &result,
		sname:  sname,
		lname:  lname,
		size:   1,
		opts:   opts,
		unique: true,
	}

	o.addArg(a)

	return &result
}

func (o *command) String(sname string, lname string, opts *Options) *string {
	var result string

	a := &arg{
		result: &result,
		sname:  sname,
		lname:  lname,
		size:   2,
		opts:   opts,
		unique: true,
	}

	o.addArg(a)

	return &result
}

func (o *command) File(sname string, lname string, flag int, perm os.FileMode, opts *Options) *os.File {
	var result os.File

	a := &arg{
		result:   &result,
		sname:    sname,
		lname:    lname,
		size:     2,
		opts:     opts,
		unique:   true,
		fileFlag: flag,
		filePerm: perm,
	}

	o.addArg(a)

	return &result
}

func (o *command) List(sname string, lname string, opts *Options) *[]string {
	result := make([]string, 0)

	a := &arg{
		result: &result,
		sname:  sname,
		lname:  lname,
		size:   2,
		opts:   opts,
		unique: false,
	}

	o.addArg(a)

	return &result
}

func (o *command) Selector(sname string, lname string, options []string, opts *Options) *string {
	var result string

	a := &arg{
		result:   &result,
		sname:    sname,
		lname:    lname,
		size:     2,
		opts:     opts,
		unique:   true,
		selector: &options,
	}

	o.addArg(a)

	return &result
}

func (o *command) Happened() bool {
	return o.parsed
}

func (o *command) Usage() string {
	// Stay classy
	maxWidth := 100
	// List of arguments from all preceding commands
	arguments := make([]*arg, 0)
	// First get line of commands until root
	var chain []string
	current := o
	for current != nil {
		chain = append(chain, current.name)
		// Also add arguments
		if current.args != nil {
			arguments = append(arguments, current.args...)
		}
		current = current.parent
	}
	// Reverse the slice
	last := len(chain) - 1
	for i := 0; i < len(chain)/2; i++ {
		chain[i], chain[last-i] = chain[last-i], chain[i]
	}
	// If this command has sub-commands we need their list
	commands := make([]command, 0)
	if o.commands != nil && len(o.commands) > 0 {
		chain = append(chain, "<command>")
		for _, v := range o.commands {
			commands = append(commands, *v)
		}
	}

	// Build result description
	var result string = "usage:"
	leftPadding := len("usage: " + chain[0] + "")
	// Add preceding commands
	for _, v := range chain {
		result = addToLastLine(result, v, maxWidth, leftPadding, true)
	}
	// Add arguments from this and all preceding commands
	for _, v := range arguments {
		result = addToLastLine(result, v.usage(), maxWidth, leftPadding, true)
	}

	// Add program/command description to the result
	result = result + "\n\n" + strings.Repeat(" ", leftPadding)
	result = addToLastLine(result, o.description, maxWidth, leftPadding, true)
	result = result + "\n\n"

	// Add list of sub-commands to the result
	if len(commands) > 0 {
		cmdContent := "Commands:\n\n"
		// Get biggest padding
		var cmdPadding int
		for _, com := range commands {
			if len("   "+com.name+"   ") > cmdPadding {
				cmdPadding = len("   " + com.name + "   ")
			}
		}
		// Now add commands with known padding
		for _, com := range commands {
			cmd := "   " + com.name
			cmd = cmd + strings.Repeat(" ", cmdPadding-len(cmd))
			cmd = addToLastLine(cmd, com.description, maxWidth, cmdPadding, true)
			cmdContent = cmdContent + cmd + "\n"
		}
		result = result + cmdContent + "\n"
	}

	// Add list of arguments to the result
	if len(arguments) > 0 {
		argContent := "Arguments:\n\n"
		// Get biggest padding
		var argPadding int
		// Find biggest padding
		for _, argument := range arguments {
			if len(argument.lname)+13 > argPadding {
				argPadding = len(argument.lname) + 13
			}
		}
		// Now add args with padding
		for _, argument := range arguments {
			arg := "   "
			if argument.sname != "" {
				arg = arg + "-" + argument.sname + "   "
			} else {
				arg = arg + "     "
			}
			arg = arg + "--" + argument.lname
			arg = arg + strings.Repeat(" ", argPadding-len(arg))
			if argument.opts != nil && argument.opts.Help != "" {
				arg = addToLastLine(arg, argument.opts.Help, maxWidth, argPadding, true)
			}
			argContent = argContent + arg + "\n"
		}
		result = result + argContent + "\n"
	}

	return result
}

func getLastLine(input string) string {
	slice := strings.Split(input, "\n")
	return slice[len(slice)-1]
}

func addToLastLine(base string, add string, width int, padding int, canSplit bool) string {
	// If last line has less than 10% space left, do not try to fill in by splitting else just try to split
	hasTen := (width - len(getLastLine(base))) > width/10
	if len(getLastLine(base)+" "+add) >= width {
		if hasTen && canSplit {
			adds := strings.Split(add, " ")
			for _, v := range adds {
				base = addToLastLine(base, v, width, padding, false)
			}
			return base
		}
		base = base + "\n" + strings.Repeat(" ", padding)
	}
	base = base + " " + add
	return base
}

func (o *parser) Parse(args []string) error {
	subargs := make([]string, len(args))
	copy(subargs, args)

	result := o.parse(&subargs)
	unparsed := make([]string, 0)
	for _, v := range subargs {
		if v != "" {
			unparsed = append(unparsed, v)
		}
	}
	if result == nil && len(unparsed) > 0 {
		return errors.New("too many arguments")
	}

	return result
}
