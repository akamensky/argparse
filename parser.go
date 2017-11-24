package argparse


type parser struct {
	command
}

// Will parse provided list of arguments
// common usage would be to pass directly os.Args
func (o *parser) Parse(args []string) {
	// If user did not set progname in advance, set progname from os.Args
	if o.name == "" {
		o.name = args[0]
	}
	
	args = args[1:]
	
	// TODO: Implement sub-commands parsing
	
	// Iterate over the rest of args
	for i, v := range args {
		i++
		k := len(v)
		k++
	}
}