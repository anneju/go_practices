package main

import (
	"fmt"
	"log"
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
}

// 3. reference types (pointers, slices, maps, functions, channels)
func changeValueOfPointer(num *int) {
	// *int means pointer
	*num = 25
}
