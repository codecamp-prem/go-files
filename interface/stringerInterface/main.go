// Package main : fmt.Stringer interface that is in fmt package. It allow any type to decide how it will displayed when printed. It's easy to setup any type to Satisfy Stringer; just define the String() method that returns the string.
package main

import "fmt"

// Stinger : interface that allows how any type will be displayed when printed.
type Stinger interface {
	String() string
}

type gallons float64

func (g gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type liters float64

func (l liters) String() string {
	return fmt.Sprintf("%0.2f ltr", l)
}

type millimeters float64

func (m millimeters) String() string {
	return fmt.Sprintf("%0.2f ml", m)
}

func main() {
	fmt.Println(gallons(12.908345))
	fmt.Println(liters(12.908345))
	fmt.Println(millimeters(12.908345))
}
