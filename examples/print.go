package main

import (
	"github.com/akamensky/argparse"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	parser := argparse.NewParser("print", "Prints provided string to stdout")

	s := parser.String("s", "string", &argparse.Options{Required:true, Help: "String to pring"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Print(parser.Usage())
	}
	fmt.Println(*s)
}