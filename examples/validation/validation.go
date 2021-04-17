// Validation functions always get a string slice as input and an error as the return type
// the string slice represents the 1+ instances of data passed to the argument
// the error represents what is returned in case the validation finds an error
// it must be nil if everything is ok
package main

import (
	"errors"
	"fmt"
	"github.com/akamensky/argparse"
	"os"
	"strconv"
	"strings"
)

// An example that splits one or more strings based on the passed separator and returns the first n entries
func main() {
	// we create a new parser
	parser := argparse.NewParser("split", "split one or more strings using a separator")

	// we expect a list of strings passed as an input
	inputStrings := parser.StringList("i", "input", &argparse.Options{
		Required: true,
		Help:     "input strings to split",
	})

	// each string is split based on a separator argument, that has to be one character long
	// in this case only one input is passed, and it can be accessed via args[0]
	// we then check that the separator string is only made of one character
	separator := parser.String("s", "separator", &argparse.Options{
		Required: true,
		Validate: func(args []string) error {
			if len(args[0]) != 1 {
				return errors.New("a separator can only have one character in it")
			}
			return nil
		},
		Help: "a one character separator, used to parse the string",
	})

	// we create a limit argument, that must be positive
	// data passed to a list argument can be accessed by iterating over the args slice
	// be careful: if you specify a validation function and you catch a bad input that
	// generates an error inside of it, the default error check will not be executed
	limits := parser.IntList("l", "limit", &argparse.Options{
		Required: false,
		Validate: func(args []string) error {
			for _, limit := range args {
				if n, err := strconv.ParseInt(limit, 10, 64); err != nil || n < 0 {
					return errors.New("limit must be a positive integer")
				}
			}
			return nil
		},
		Help:    "a list of limits of tokens to return per input",
		Default: nil,
	})

	if err := parser.Parse(os.Args); err != nil {
		fmt.Println(parser.Usage(err))
		os.Exit(1)
	}

	// We split every input string according to its limit, if passed and print the result
	var actualLimit int
	for idx, inputString := range *inputStrings {
		if idx >= len(*limits) {
			// if no limit is specified, we return every substring
			actualLimit = -1
		} else {
			actualLimit = (*limits)[idx]
		}
		fmt.Printf("string(%s) => %v\n", inputString, strings.SplitN(inputString, *separator, actualLimit))
	}
}
