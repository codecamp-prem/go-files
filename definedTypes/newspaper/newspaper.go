// Package newspaper manage app functionality
/*
- type Subscriber helps create new subscriber
- type Employee helps create new employee details
*/
package newspaper

// Subscriber helps create new subscriber
type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

// Employee helps create new employee details
type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}

// Address helps to create employee and subscriber address
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
