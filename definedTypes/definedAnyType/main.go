// Package main: Go defined Types MOST often use structs as their underlying types, but they can be also based on ints, strings , booleans and any othe types
package main

import "fmt"

//Liters :new defined type float64
type Liters float64

//Gallons :new defined type float64
type Gallons float64

// Millimeters :new defined type float64
type Millimeters float64

// Number :new defined type int
type Number int

// MyType new defined type string
type MyType string

// ToGallons converts 'liters to Gallon'
func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

// ToGallons converts 'Millimeters to Gallon'
func (m Millimeters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

// ToLitters converts gallons to Liters
func (g Gallons) ToLitters() Liters {
	return Liters(g * 3.785)
}

// Double demos the pointer reciever parameter
func (n *Number) Double() {
	*n *= 2
}

// method : DEmo for method with value reciever
func (m MyType) method() {
	fmt.Println("Method with value reciever")
}

// pointerMethod: Demo for method with pointer reciever
func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer reciever")
}

func main() {
	fmt.Println(Liters(6.7) + Liters(2.3))
	fmt.Println(Gallons(99.1) + Gallons(0.9))
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	additionalCarFuelInltrs := Liters(40.0)
	carFuel += additionalCarFuelInltrs.ToGallons() // converts Liter to Gallon before adding
	additionalBusFuelInGallons := Gallons(30.0)
	busFuel += additionalBusFuelInGallons.ToLitters() // converts Gallons to Liter before adding
	fmt.Printf("Car Fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus Fuel: %0.1f liters\n", busFuel)

	number := Number(4) // you’ll need to store the value in a variable, which will then allow Go to get a pointer to it: Number(4).Double() or &Number(4) results in error
	number.Double()
	fmt.Println("number after calling Double: ", number)

	value := MyType("a value")
	pointer := &value
	value.method()
	value.pointerMethod() //Value automatically converted to pointer
	pointer.method()      // Value at pointer automatically retrieved
	pointer.pointerMethod()

	/*
	  Method with value reciever
	  Method with pointer reciever
	  Method with value reciever
	  Method with pointer reciever
	*/
}
