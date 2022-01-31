package main

import (
	"fmt"

	"github.com/codecamp-prem/definedTypes/newspaper"
)

func main() {
	var employee newspaper.Employee
	employee.Name = "John Doe"
	employee.Salary = 29999
	employee.HomeAddress.Street = "Empire Street, block 007"
	employee.HomeAddress.City = "New York City"
	employee.HomeAddress.State = "New York"
	employee.HomeAddress.PostalCode = "10004"

	fmt.Println("--Employee Details--")
	fmt.Println("Name: ", employee.Name)
	fmt.Println("Salary: ", employee.Salary)
	fmt.Println("Street: ", employee.HomeAddress.Street)

	employee2 := newspaper.Employee{
		Name:   "Jane Doe",
		Salary: 30000,
		HomeAddress: newspaper.Address{
			Street:     "CalTech Uni Rd, block 006",
			City:       "Los Angles",
			State:      "California",
			PostalCode: "90006",
		},
	}

	fmt.Println("-----------")
	fmt.Println("Name: ", employee2.Name)
	fmt.Println("Salary: ", employee2.Salary)
	fmt.Println("Street: ", employee2.HomeAddress.Street)
}
