// if type declares methods with the pointers reciever,
// then you'll only able to use the pointers to that type when
// assigning to interface variables
package main

import (
	"fmt"
)

// Switch : defined type string
type Switch string

// toggle : method with pointers reciever
func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

// Toggleable : interface with method toggle
type Toggleable interface {
	toggle()
}

func main() {
	s := Switch("on")
	var t Toggleable = &s //Assign a pointer
	t.toggle()
	t.toggle()
}
