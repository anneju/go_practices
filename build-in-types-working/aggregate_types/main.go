package main

import (
	"fmt"
)

// 2. aggregate types (array, struct)
type Car struct {
	numberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {
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
}
