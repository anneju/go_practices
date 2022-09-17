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
}
