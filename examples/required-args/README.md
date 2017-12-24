# Required vs Optional Arguments

The purpose of this example is to demonstrate the difference between required
and optional arguments using string arguments.

## Setup

The first step for any project is to build a new parser:

```go
parser := argparse.NewParser("Required-args", "Example of required vs optional arguments")
```

For this demonstration, two shorthand arguments are used:

```go
// Required shorthand argument
foo := parser.String("f", "foo", &argparse.Options{Required: true, Help: "foo is a required string option"})

// Optional shorthand argument
bar := parser.String("b", "bar", &argparse.Options{Required:false, Help: "bar is not a required option"})
```

`foo` is a required string argument, and `bar` is an optional string argument.

Using `required-arguments.go` without foo demonstrates this:

```bash
$ go run required-args.go
[-f|--foo] is required
usage: Required-args [-h|--help] -f|--foo "<value>" [-b|--bar "<value>"]

                     Example of required vs optional arguments

Arguments:

   -h   --help    Print help information
   -f   --foo     foo is a required string option
   -b   --bar     bar is not a required option


```

Output when foo is provided:

```bash
$ go run required-args.go --foo foo
Provided arguments:
foo (required): foo
```
