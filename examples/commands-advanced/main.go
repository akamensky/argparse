package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/akamensky/argparse/examples/commands-advanced/zoo"
	"log"
	"os"
)

func main() {
	// Create new parser object
	parser := argparse.NewParser("zooprog", "Program that walks us through the zoo")

	// Note that we create argument off the parser, so it will work on all scopes
	// also this is required argument, meaning we cannot avoid but to always provide it
	name := parser.String("", "name", &argparse.Options{Help: "Provide an optional name for the animal", Required: true})

	// dog command
	dogCmd := parser.NewCommand("dog", "We are going to see dog")
	// dog sub-commands
	dogSpeak := dogCmd.NewCommand("speak", "Make the dog speak")
	dogFeed := dogCmd.NewCommand("feed", "Make the dog eat")
	dogSummon := dogCmd.NewCommand("summon", "Make the dog come over")
	dogPlay := dogCmd.NewCommand("play", "Make the dog play")

	// cat command
	catCmd := parser.NewCommand("cat", "We are going to see cat")
	// cat sub-commands
	catSpeak := catCmd.NewCommand("speak", "Make the cat speak")
	catFeed := catCmd.NewCommand("feed", "Make the cat eat")
	catSummon := catCmd.NewCommand("summon", "Make the cat come over")
	catPlay := catCmd.NewCommand("play", "Make the cat play")

	// Optional argument for dog.
	// Note that we create this argument for dogCmd, which means catCmd will not have this argument
	wiggleFlag := dogCmd.Flag("", "wiggle", &argparse.Options{Help: "Makes the dog to wiggle its tail"})

	// Now parse the arguments
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		return
	}

	// When there is no name provided
	// we need to give default value
	if *name == "" {
		*name = "unnamed"
	}

	// If user gave `dog` command
	if dogCmd.Happened() {
		// Get a dog and give it a name
		animal := zoo.NewDog(*name)

		if dogSpeak.Happened() {
			animal.Do("speak")
		} else if dogFeed.Happened() {
			animal.Do("feed")
		} else if dogSummon.Happened() {
			animal.Do("summon")
		} else if dogPlay.Happened() {
			animal.Do("play")
		} else {
			// This should be unreachable
			log.Fatal("Uh-oh, unknown command! Impossible! Ex-ter-mi-nate!")
		}

		// If we got wiggle flag, then well... wiggle, ok?
		if *wiggleFlag {
			animal.WiggleTail()
		}
	} else if catCmd.Happened() {
		// Get a cat and give it a name
		animal := zoo.NewCat(*name)

		if catSpeak.Happened() {
			animal.Do("speak")
		} else if catFeed.Happened() {
			animal.Do("feed")
		} else if catSummon.Happened() {
			animal.Do("summon")
		} else if catPlay.Happened() {
			animal.Do("play")
		} else {
			// This should be unreachable
			log.Fatal("Uh-oh, unknown command! Impossible! Ex-ter-mi-nate!")
		}
	} else {
		// This should be unreachable
		log.Fatal("Uh-oh, something weird happened")
	}
}
