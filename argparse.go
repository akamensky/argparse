package argparse

import (
	"reflect"
)

type Options struct {
	Required bool
	Validate func(arg string) error
	Help string
}

type arg struct {
	resultPointerType reflect.Type
	result interface{}
	opts *Options
	sname string
	lname string
	size int
}

func NewParser(name string, description string) parser {
	p := parser{}
	
	p.name = name
	p.description = description
	
	p.args = make(map[string]arg)
	p.commands = make([]command, 0)
	
	return p
}


