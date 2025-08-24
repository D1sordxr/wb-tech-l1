package main

import (
	"fmt"
	"strings"
)

type Human struct {
	name    string
	age     int
	phone   string
	isAdult bool
}

func (h Human) Say(words ...string) {
	speech := strings.Join(words, " ")

	fmt.Printf("%s says: %s!\n", h.name, speech)
}

func (h Human) Name() string {
	return h.name
}

func (h Human) Age() int {
	return h.age
}

func (h Human) Phone() string {
	return h.phone
}

func (h Human) IsAdult() bool {
	return h.age >= 18
}

type Action struct {
	Human
}

func main() {
	action := Action{
		Human{
			name:    "Andrea",
			age:     25,
			phone:   "+8(800)555-35-35",
			isAdult: true,
		},
	}

	action.Say("Hello", "this", "is", "completed", "l1.1", "task")

	fmt.Printf("Name: %s.\n", action.Name())
	fmt.Printf("Age: %d.\n", action.Age())
	fmt.Printf("Is adult: %v.\n", action.IsAdult())
}
