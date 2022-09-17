package main

import (
	"fmt"
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 50000, FullTime: true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 60000, FullTime: true},
	{FirstName: "Allen", LastName: "Anderson", Salary: 15000, FullTime: false},
	{FirstName: "Margaret", LastName: "Carter", Salary: 100000, FullTime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	log.Println(myStaff.All())
	fmt.Println()

	staff.OverPaidLimit = 60000
	log.Println("Over paid staffs", myStaff.Overpaid())

	log.Println("Under paid staffs", myStaff.Underpaid())

	// myStaff.notVisible()
}
