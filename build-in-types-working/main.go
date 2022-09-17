package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/eiannone/keyboard"
)

// 1. basic types (numbers, strings, booleans)
var myInt int

// if sure will develope on old computers. But almost never use below types and all discuraged
// var myInt16 int16
// var myInt32 int32
// var myInt64 int64

var muUint uint // holds posive value or zero

// use float64 if dealing with very large number
var myFloat32 float32 // hold 32-digits number
var myFloat64 float64 // hold 64-digits number

// 2. aggregate types (array, struct)
type Car struct {
	numberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

// 3. reference types (pointers, slices, maps, functions, channels)
type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

// (a *Animal) is called receiver, *Animal means pointer to Animal. This comes right before the function name
func (a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println()
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
	fmt.Println()
}

// channel
// can only pass one type to channel
var keyPressChan chan rune

// 4. interface type

func main() {
	// 1. basic types (numbers, strings, booleans)
	myInt = 10
	muUint = 20
	myFloat32 = 10.1
	myFloat64 = 100.1

	log.Println(myInt, muUint, myFloat32, myFloat64)

	myString := "Anne"
	log.Println(myString)

	// in fact in go, string is immutable, it's create a new string.
	myString = "Gogo"
	log.Println(myString)

	var myBool = true
	myBool = false
	log.Println(myBool)
	fmt.Println("")

	// 2. aggregate types (array, struct)
	var myStrings [3]string
	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("First element in array is", myStrings[0])

	var myCar Car
	myCar.numberOfTires = 4
	myCar.Luxury = false
	myCar.Make = "Benz"

	myNewCar := Car{
		numberOfTires: 4,
		Luxury:        true,
		Make:          "Volvo",
		Model:         "XC90",
		Year:          2019,
	}

	fmt.Printf("My New Car is a %d %s %s", myNewCar.Year, myNewCar.Make, myNewCar.Model)
	fmt.Println(" ")

	// 3. reference types (pointers, slices, maps, functions, channels)
	// 3-1. pointers
	x := 10

	// use `&` to get memory address that store the value of x
	myFirstPointer := &x

	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	// the star * means to go to the pointer myFirstPointer is pointing, and change the value at that memory address
	// use `*` to re-assign the value for that pointer
	*myFirstPointer = 15
	fmt.Println("x is now", x)

	changeValueOfPointer(&x)
	fmt.Println("After function call, x is now", x)

	// 3-2. slices
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "fish")
	animals = append(animals, "horse")
	fmt.Println(animals)

	// another way to loop the slice
	for index, animal := range animals {
		fmt.Println(index, " ", animal)
	}

	fmt.Println("First two elements are", animals[0:2])

	fmt.Println("The slice is", len(animals), "elements long")

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)
	fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals))

	DeleteFromSlice(animals, 1)
	fmt.Println(animals)

	// 3-3. maps
	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	// map is never sorted in golang
	for key, value := range intMap {
		fmt.Println(key, value)
	}

	// delete a element from map
	fmt.Println("delete a element from map")
	delete(intMap, "four")
	for key, value := range intMap {
		fmt.Println(key, value)
	}

	// find if an element is in a map
	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		// if cannot find it, it fallback to default value of el type, here it's int, so it's 0
		fmt.Println(el, "is not in map")
	}

	// change an element in map
	intMap["two"] = 4

	// 3-4. functions
	z := addTwoNumbers(2, 4)
	fmt.Println(z)

	z_2 := add2Numbers(3, 4)
	fmt.Println(z_2)

	myTotal := sumMany(2, 3, 4, 5, 88, -7)
	fmt.Println(myTotal)

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.NumberOfLegs = 4
	dog.Says()

	cat := Animal{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}
	cat.Says()
	cat.HowManyLegs()

	// 3-5. channels
	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit.")

	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		// send char to channel
		keyPressChan <- char
	}
}

// 3. reference types (pointers, slices, maps, functions, channels)
func changeValueOfPointer(num *int) {
	// *int means pointer
	*num = 25
}

// delete element from slice for a given index
func DeleteFromSlice(a []string, i int) []string {
	// copy the last element from slice to this index
	a[i] = a[len(a)-1]
	// replace the last element of slice with nothing
	a[len(a)-1] = ""
	// truncate the slice, here we return every things in that slice except the last element
	a = a[:len(a)-1]
	return a
}

// function
func addTwoNumbers(x, y int) int {
	return x + y
}

// naked return function, but rarely used.
func add2Numbers(x, y int) (sum int) {
	sum = x + y
	return
}

// variadic function
func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total = total + x
	}
	return total
}

// channels

func listenForKeyPress() {
	for {
		key := <-keyPressChan // key is whatever passed into keyPressChan
		fmt.Println("You pressed", string(key))
	}
}
