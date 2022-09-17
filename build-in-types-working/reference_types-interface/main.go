package main

import "fmt"

// interface type
type Animal interface {
	Says() string
	HowManyLegs() int
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (c *Cat) Says() string {
	return c.Sound
}

func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}

func main() {
	// ask a riddle
	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	// Riddle(dog)
	Riddle(&dog) // reference to dog

	var cat Cat
	cat.Name = "cat"
	cat.NumberOfLegs = 4
	cat.Sound = "meow"
	cat.HasTail = true

	// Riddle(cat) // cannot do this because it expect a Dog
	Riddle(&cat)
}

// func Riddle(d Dog) {
// 	riddle := fmt.Sprintf(`This is an animal says %s and has %d legs. What animal is it?`, d.Sound, d.NumberOfLegs)
// 	fmt.Println(riddle)
// }

func Riddle(a Animal) {
	riddle := fmt.Sprintf(`This is an animal says %s and has %d legs. What animal is it?`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
