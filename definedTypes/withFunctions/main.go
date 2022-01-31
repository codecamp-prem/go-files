package main

import "fmt"

type part struct {
	description string
	count       int
}

// showInfo function that's prints a part's fields
func showInfo(p part) { //Declare one parameter, with "part" as it's type
	fmt.Println("Description: ", p.description)
	fmt.Println("Count: ", p.count)
}

// minimumOrder function that create's a part with
// a specified description and
// predefined value for count field.
func minimumOrder(description string) part { // Declare one return value, with a type of "part"
	var p part
	p.description = description
	p.count = 100
	return p
}

func main() {
	var bolts part //Create a "part" value
	bolts.description = "Hex bolts"
	bolts.count = 24
	showInfo(bolts)

	p := minimumOrder("Hex bolts")
	fmt.Println(p.description, p.count)
}
