// package main ; Describe usefulness of empty interface in goLang
package main

import "fmt"

// acceptAnything ; To call methods on a value with the empty interface type, you’d need to use a type assertion to get a value of the concrete type back.
func acceptAnything(thing interface{}) {
	fmt.Println(thing)
	car, ok := thing.(car)
	if ok {
		car.makeSound()
	}
}

type car string

func (c car) makeSound() {
	fmt.Println("BRRR R brrr BRR")
}

func main() {
	acceptAnything(3.14)
	acceptAnything(car("Mustang"))
}
