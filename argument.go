package argparse

import "reflect"

type Options struct {
	Required bool
	Validate func(arg string) error
	Help     string
}

type arg struct {
	resultPointerType reflect.Type
	result            interface{}
	opts              *Options
	sname             string
	lname             string
	size              int
}

func (o *arg) parse([]string) {
	switch o.result.(type) {
	case bool:

	}
}
