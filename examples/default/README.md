# Default Argument Value
To demonstrate usage of default values of arguments.

## Setup

``` go
parser := argparse.NewParser("default", "Program to test default values")
```
	
## Required is false
Create an argument where the required option is set to false:

``` go
s := parser.String("s", "string", &argparse.Options{Required: false, Help: "String to print", Default: "Hello"})
```

On execution of program, which prints the argument as

``` go
fmt.Println(*s)
```

``` bash
$ go run default.go
Hello
```

Also, the default value is printed in the help.

## Required is true

``` go
s := parser.String("s", "string", &argparse.Options{Required: true, Help: "String to print", Default: "Hello"})
```

On execution,

``` bash
$ go run default.go
[-s|--string] is required
usage: default [-h|--help] -s|--string "<value>"

               Program to test default values

Arguments:

  -h  --help    Print help information
  -s  --string  String to print

exit status 1
```

