package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
	"time"
)

func main() {
	// Create new parser object
	parser := argparse.NewParser("flags", "Simple example of argparse flags")

	// Create count integer argument
	count := parser.Int("c", "count", &argparse.Options{Required: true, Help: "Number of iterations to count"})

	// Create delay float argument
	delay := parser.Float("d", "delay", &argparse.Options{Default: 1.0, Help: "Delay between iterations"})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	// Print (*count) iterations with (*delay) second delays
	delay_duration := time.Duration((*delay)*1000) * time.Millisecond
	for i := 0; i < *count; i++ {
		time.Sleep(delay_duration)
		fmt.Println("Iteration:", i+1)
	}
}
