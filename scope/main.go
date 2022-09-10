package main

import (
	"fmt"
	"myapp/packageone"
)

var one = "One"

func main() {
	var somethingElse = "Block level variable"

	fmt.Println(somethingElse)

	myFunc()

	fmt.Print(packageone.PackageVar)
}

func myFunc() {
	fmt.Println(one)
}
