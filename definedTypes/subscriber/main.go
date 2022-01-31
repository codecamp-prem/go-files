package main

import (
	"fmt"

	"github.com/codecamp-prem/definedTypes/newspaper"
)

func printInfo(s *newspaper.Subscriber) {
	fmt.Println("Name: ", s.Name)
	fmt.Println("Rate: ", s.Rate)
	fmt.Println("Active? :", s.Active)
}

func applyDiscount(s *newspaper.Subscriber) {
	s.Rate = 3.99
}

func defaultSubscriber(name string) *newspaper.Subscriber {
	var s newspaper.Subscriber
	s.Name = name
	s.Rate = 4.99
	s.Active = true
	return &s //return the type Subscriber struct
}

func main() {
	subs1 := defaultSubscriber("Rocket Singh") // This is no longer a struct, it's a struct pointer
	//subs1.Rate = 2.99
	applyDiscount(subs1)
	printInfo(subs1)

	subs2 := defaultSubscriber("Honey Singh")
	printInfo(subs2)
}
