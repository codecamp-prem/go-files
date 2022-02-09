package main

import "fmt"

// Whistle : a defined type
type Whistle string

// MakeSound : method for Whistle type
func (w Whistle) MakeSound() {
	fmt.Println("siii siiiiii")
}

// Horn : a defined type
type Horn string

// MakeSound : method for Horn type
func (h Horn) MakeSound() {
	fmt.Println("puh puh poh")
}

// NoiseMaker : an interface
type NoiseMaker interface {
	MakeSound()
}

func main() {
	var toy NoiseMaker
	toy = Whistle("christiano")
	toy.MakeSound()
	toy = Horn("train")
	toy.MakeSound()
}
