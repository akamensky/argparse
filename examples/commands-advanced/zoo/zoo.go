package zoo

import (
	"fmt"
	"log"
)

var actions = map[string]int{
	"speak":  1,
	"feed":   2,
	"summon": 3,
	"play":   4,
}

type animal struct {
	Name   string
	says   string
	feed   string
	summon string
	play   string
}

func (o *animal) Do(action string) {
	switch actions[action] {
	case 1:
		o.doSay()
	case 2:
		o.doFeed()
	case 3:
		o.doSummon()
	case 4:
		o.doPlay()
	default:
		log.Fatal("Wow, we got unknown action, that should have never happened")
	}
}

func (o *animal) doSay() {
	fmt.Println(o.says)
}

func (o *animal) doFeed() {
	fmt.Println(o.feed)
}

func (o *animal) doSummon() {
	fmt.Println(o.summon)
}

func (o *animal) doPlay() {
	fmt.Println(o.play)
}

type dog struct {
	animal
}

// Dog has an extra action to wiggle its tail
func (o *dog) WiggleTail() {
	fmt.Println("* Dog wiggles its tail and stares at you with love")
}

type cat struct {
	animal
}

// NewDog makes new dog in the zoo
func NewDog(name string) dog {
	result := new(dog)
	result.Name = name
	result.says = "Woof"
	result.feed = "* Dog eats the food and stays now happily wiggles it tail"
	result.summon = "* Dog immediately shows up wiggling its tail"
	result.play = "* Dog runs around in excitement"

	return *result
}

// NewCat makes new cat in the zoo (who has cats in the zoo like ever?)
func NewCat(name string) cat {
	result := new(cat)
	result.Name = name
	result.says = "Meow"
	result.feed = "* Cat eats the food and slowly walks away"
	result.summon = "* Cat is nowhere to be found"
	result.play = "* Cat stares at you in disgust for awhile and goes back to sleep"

	return *result
}
